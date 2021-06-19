package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

//

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//只允许*一次
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' { //找到第一个带*的参数的组
				break
			}
		}
	}
	return parts
} //

//添加route
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	parts := parsePattern(pattern) //解析路径
	key := method + "-" + pattern
	_, ok := r.roots[method] //查看请求方法是否存在
	if !ok {                 //不存在就存一个新的
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts /* 被用'/'分割过的路径 */, 0) 
	r.handlers[key] = handler
}

//获取route
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handler(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND : %s\n", c.Path)
	}
}
