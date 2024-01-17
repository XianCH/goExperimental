package router

import "net/http"

type nodeType uint8

const (
	static nodeType = iota
	root
	param
)

type node struct {
	path      string
	WildChild string
	Children  *node
	nType     nodeType
	handle    Handle
}

type router struct {
	root *node
}

type Handle func(*http.Request, http.ResponseWriter)

// 创建新的路由
func NewRouter() *router {
	return &router{
		root: new(node),
	}
}

func (r *router) addRouter(path string, handler Handle) error {}
