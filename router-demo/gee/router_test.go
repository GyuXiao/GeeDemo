package gee

import (
	"fmt"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

// 新建一个 Route 的单测
func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geekGyu")
	if n == nil {
		t.Fatal("nil should not be return")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "geekGyu" {
		t.Fatal("name should be equal to 'geekGyu'")
	}
	fmt.Printf("match path : %s, params['name'] : %s\n", n.pattern, ps["name"])
}
