package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct{}

func main() {
 response, err := http.Get("https://google.com")
 if err != nil {
	fmt.Println("Error: ", err)
	os.Exit(1)
 } else {
	lw := logWriter{}
	io.Copy(lw, response.Body)
 }
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}