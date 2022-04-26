package gee

import (
	"net/http"
)

/**
实现功能：
实现了路由映射表，提供了用户注册静态路由的方法，包装了启动服务的函数
*/

// 定义 gee 需要的 handler 函数
// 定义路由映射的处理方法
// 在 engine 中添加一张路由映射表 router
// key ： method + "-" + pattern
// value：handler
type HandlerFunc func(*Context)

// Engine 定义一个结构体
type Engine struct {
	router *router
}

// Get 请求
func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.AddRoute("GET", pattern, handler)
}

// Post 请求
func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.AddRoute("POST", pattern, handler)
}

func (engine *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// Run 请求
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现 handler 接口里的 ServeHTTP
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.handle(c)
}

// 暴露给使用者一个新建 Engine 对象实例的方法
func NewInstance() *Engine {
	return &Engine{router: newRouter()}
}
