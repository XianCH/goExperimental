package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v2] Hello,Kubenates!")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	duration := time.Since(start)
	if duration.Seconds() >= 15 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error :%v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":5000", nil)
}
