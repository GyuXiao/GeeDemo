package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 在 HandlerFunc 希望能够访问到解析的参数，
// 因此，在 Context 里增加了
// 属性 Params
// 方法 Param(key string)
// 来提供对路由参数的访问
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 提供访问 Query 和 PostForm 参数的方法
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 轮询方法
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 方法
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
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
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
