package gee

import "net/http"

type router struct {
	handlerMap map[string]HandlerHttp
}

// 对外提供 new 方法创建实例
func newRouter() *router {
	return &router{handlerMap: make(map[string]HandlerHttp)}
}

// 添加 路由 的处理逻辑
func (r *router) addRouter(method string, pattern string, handler HandlerHttp) {
	key := method + "-" + pattern
	r.handlerMap[key] = handler
}

// 提供动态路由的支持
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlerMap[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}
