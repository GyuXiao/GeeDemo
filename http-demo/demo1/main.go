package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置两个路由，`/` 和 `/hello`，分别绑定 indexHandler 和 helloHandler，根据不同的 HTTP 请求会调用不同的处理函数
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	// 是用来启动 Web 服务的，
	// 第一个参数是地址，:9999表示在 9999 端口监听。
	// 第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 响应的是 URL.Path = /
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
}

// 响应的是请求头的键值对信息
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
