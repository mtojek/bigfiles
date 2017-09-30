package main

func main() {
	createHttpServer(readConfiguration(), createBinaryFileSystem()).listenAndServe()
}
