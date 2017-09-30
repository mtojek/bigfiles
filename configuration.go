package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/c2h5oh/datasize"
)

type configuration struct {
	hostPort    string
	maxFileSize int64
}

func readConfiguration() *configuration {
	var hostPort string
	var maxFileSize string

	flag.StringVar(&hostPort, "hostPort", ":5000", "host:port of the HTTP server")
	flag.StringVar(&maxFileSize, "maxFileSize", "9223372036854775807B", "maximum size of file that can be downloaded")
	flag.Parse()

	var c *configuration
	var err error
	if c, err = validate(hostPort, maxFileSize); err != nil {
		log.Fatal(err)
	}
	return c
}

func validate(hostPort, maxFileSize string) (*configuration, error) {
	if hostPort == "" {
		return nil, errors.New("host:post must not be empty")
	}

	if maxFileSize == "" {
		return nil, errors.New("maxFileSize must no be empty")
	}

	var uSize datasize.ByteSize
	if err := uSize.UnmarshalText([]byte(maxFileSize)); err != nil {
		return nil, err
	}

	if uSize > math.MaxInt64 {
		return nil, fmt.Errorf("file size cannot be greater than %v B", math.MaxInt64)
	}

	return &configuration{hostPort: hostPort, maxFileSize: int64(uSize)}, nil
}
