package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	lba "github.com/x14n/goExperimental/LBA"
)

// 实现一个代理服务器
type Proxy struct {
	balancer lba.Balancer
}

func NewProxy(balancer lba.Balancer) *Proxy {
	return &Proxy{balancer: balancer}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backendServer := p.balancer.NextServer()
	backendURL, err := url.Parse(backendServer)
	if err != nil {
		http.Error(w, "Invalid backend server URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		http.Error(writer, "Error contacting backend server", http.StatusInternalServerError)
	}

	// 修改请求的URL和Host信息
	r.URL.Host = backendURL.Host
	r.URL.Scheme = backendURL.Scheme
	r.Host = backendURL.Host

	proxy.ServeHTTP(w, r)
}

func main() {
	// 后端服务器列表
	servers := []string{"http://localhost:8081", "http://localhost:8082", "http://localhost:8083"}

	// 创建一个轮询负载均衡器
	balancer := lba.NewRoundRobinBalancer(servers)

	// 创建一个代理服务器
	proxy := NewProxy(balancer)

	// 启动代理服务器
	log.Println("Starting proxy server on :8080")
	if err := http.ListenAndServe(":8080", proxy); err != nil {
		log.Fatalf("Error starting proxy server: %v", err)
	}
}
