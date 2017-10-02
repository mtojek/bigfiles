package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/c2h5oh/datasize"
	"github.com/julienschmidt/httprouter"
)

var (
	indexMessage = []byte(`
<html>
<body>
<title>BigFiles</title>
<h2>Welcome to mtojek/bigfiles</h2>
<h3>Try following sample links to download binary files:</h3>
<div>
<div><a href="/files/random/1KB">/files/random/1KB</a></div>
<div><a href="/files/random/1GB">/files/random/1GB</a></div>
<div><a href="/files/sparse/100MB">/files/sparse/100MB</a></div>
<div><a href="/files/sparse/8191PB">/files/sparse/8191PB</a></div>
</div>
<br/>
<div>
Feel free to visit project Github page - <a target="_new" href="https://github.com/mtojek/bigfiles">github.com/mtojek/bigfiles</a>.
</div>
</body>
</html>`)
	errTooBigFileSize      = errors.New("too big file size")
	errMissingFileSize     = errors.New("missing file size")
	errMissingFileType     = errors.New("missing file type")
)

type httpServer struct {
	config     *configuration
	fileSystem http.FileSystem
}

func (s *httpServer) listenAndServe() {
	router := httprouter.New()
	router.GET("/", s.index)
	router.GET("/files/:type/:size", s.createFilesHandler())
	http.ListenAndServe(s.config.hostPort, router)
}

func (s *httpServer) index(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write(indexMessage)
}

func (s *httpServer) createFilesHandler() httprouter.Handle {
	fileServer := http.FileServer(s.fileSystem)
	return func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		fileSize := ps.ByName("size")
		
		if fileSize == "" {
			http.Error(rw, errMissingFileSize.Error(), http.StatusBadRequest)
			return
		}

		if len(ps.ByName("type")) < 1 {
			http.Error(rw, errMissingFileType.Error(), http.StatusBadRequest)
			return
		}

		var bs datasize.ByteSize
		err := bs.UnmarshalText([]byte(fileSize))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if int64(bs.Bytes()) > s.config.maxFileSize {
			http.Error(rw, errTooBigFileSize.Error(), http.StatusBadRequest)
			return
		}

		rw.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.bin"`, fileSize))
		fileServer.ServeHTTP(rw, req)
	}
}

func createHttpServer(c *configuration, fs http.FileSystem) *httpServer {
	return &httpServer{c, fs}
}
