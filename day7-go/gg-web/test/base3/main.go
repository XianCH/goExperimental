package main

import (
	"fmt"
	"net/http"
)

func main() {
	engin := New()
	engin.Get("/", indexHandler)
	engin.Post("/hello", helloHandler)
	engin.Run(":9999")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
