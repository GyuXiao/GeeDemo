package main

import (
	"GeeDemo/context-demo/gee"
	"net/http"
)

// 这个 demo 要完成的事情：
// 1，将 router 独立出来，方便之后增强,并支持静态路由查询
// 2，设计 Context，封装 Request 和 Response，提供对 JSON、HTML 等返回类型的支持
func main() {
	r := gee.NewInstance()
	r.Get(
		"/",
		func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

	r.Get(
		"/hello",
		func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
		})

	r.Post(
		"/login",
		func(c *gee.Context) {
			c.JSON(
				http.StatusOK,
				gee.H{
					"username": c.PostForm("username"),
					"password": c.PostForm("password"),
				})
		})

	r.Run(":9999")
}
