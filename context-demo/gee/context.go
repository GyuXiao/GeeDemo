package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context 的结构体
type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

// 设置 Status 状态方法
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 方法
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 提供访问 Query 和 PostForm 参数的方法
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// Query 轮询方法
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// 提供快速构建 String/Data/JSON/HTML响应的方法

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.Writer.Write([]byte(html))
}

// 对外提供 new 方法
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
	}
}
