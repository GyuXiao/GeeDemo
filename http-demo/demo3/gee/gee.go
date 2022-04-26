package gee

import (
	"fmt"
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
type HandlerHttp func(http.ResponseWriter, *http.Request)

// Engine 定义一个结构体
type Engine struct {
	router map[string]HandlerHttp
}

// 以下 4 个是结构体方法
// AddRoute 模拟 Router 的注册功能
func (engine *Engine) AddRoute(method string, pattern string, handler HandlerHttp) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// Get 请求
func (engine *Engine) Get(pattern string, handler HandlerHttp) {
	engine.AddRoute("GET", pattern, handler)
}

// Post 请求
func (engine *Engine) Post(pattern string, handler HandlerHttp) {
	engine.AddRoute("POST", pattern, handler)
}

// Run 请求
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现 handler 接口里的函数 ServeHTTP
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		_, err := fmt.Fprintf(writer, "404 Not Found: %s\n", request.URL)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// 暴露给使用者一个新建 Engine 对象实例的方法
func NewInstance() *Engine {
	return &Engine{router: make(map[string]HandlerHttp)}
}
