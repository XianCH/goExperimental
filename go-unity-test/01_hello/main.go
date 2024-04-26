package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	// 将消息写入响应
	fmt.Fprintf(w, "Received message from client: %s\n", string(body))
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	fmt.Println("golang server start at port 12345")
	http.ListenAndServe(":12345", nil)
}
