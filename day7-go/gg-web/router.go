package ggweb

import (
	"log"

	"github.com/x14n/goExperimental/day7_go/gg-web/context"
)

type router struct {
	handles map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handles: make(map[string]HandlerFunc),
	}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4 - %s", method, pattern)
	key := method + "-" + pattern
	r.handles[key] = handler
}

func (r *router) handle(c *context.Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handles[key]; ok {
		handler(c)
	}
}
