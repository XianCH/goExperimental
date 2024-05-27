package main

import (
	"fmt"
	"log"
	"net/http"
)

func startBackendServer(address string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handled by backend server %s\n", address)
	})

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Printf("Starting backend server on %s\n", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server on %s: %v", address, err)
	}
}

func main() {
	// 启动多个后端服务器
	go startBackendServer(":8081")
	go startBackendServer(":8082")
	go startBackendServer(":8083")

	// 阻止main函数退出
	select {}
}
