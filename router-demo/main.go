package main

import (
	"GeeDemo/router-demo/gee"
	"net/http"
)

// 简要用 Trie 树实现动态路由
func main() {
	r := gee.NewInstance()
	r.Get(
		"/",
		func(context *gee.Context) {
			context.HTML(http.StatusOK, "<h1>Hello GeekGyu</h1>")
		})

	r.Get(
		"/hello",
		func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
		})

	r.Get(
		"/hello/:name",
		func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
		},
	)

	r.Get(
		"/assets/*filepath",
		func(context *gee.Context) {
			context.JSON(http.StatusOK, gee.H{"filepath": context.Param("filepath")})
		})

	r.Run(":9999")
}
