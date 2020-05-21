package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

type shape interfac

func main() {
	resp, error := http.Get("http://example.com/")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
