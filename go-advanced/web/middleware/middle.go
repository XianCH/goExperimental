package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }
//
// type HandlerFunc func(ResponseWriter, *Request)
//
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//     f(w, r)
// }

type middleware func(http.Handler) http.Handler

type router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *router {
	return &router{mux: make(map[string]http.Handler)}
}

func (r *router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *router) add(route string, h http.Handler) {
	var mergeHandler = h
	for i := len(r.middlewareChain); i >= 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}
	r.mux[route] = mergeHandler
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging request", r.URL, r.Method)
		next.ServeHTTP(w, r)
	})
}

func timeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Since(start)
			fmt.Println("Request took:", end)
		}()
		next.ServeHTTP(w, r)
	})
}

func ratelimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Rate limiting logic")
		next.ServeHTTP(w, r)
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func (r *router) run(address string) error {
	mux := http.NewServeMux()
	for route, handler := range r.mux {
		mux.Handle(route, handler)
	}

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return server.ListenAndServe()
}

func Test_middle() {
	r := NewRouter()
	r.Use(logger)
	r.Use(timeout)
	r.Use(ratelimit)
	r.add("/", http.HandlerFunc(Hello))
	r.run("127.0.0.1:8080")
}
