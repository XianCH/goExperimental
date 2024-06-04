package ggweb

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRouter(method string, path string, handlerFunc HandlerFunc) {
	key := method + "-" + path
	e.router[key] = handlerFunc
}

func (e *Engine) Get(path string, handlerFunc HandlerFunc) {
	e.addRouter("GET", path, handlerFunc)
}

func (e *Engine) Post(path string, handlerFunc HandlerFunc) {
	e.addRouter("POST", path, handlerFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %s\n", r.URL.Path)
	}
}

func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}
