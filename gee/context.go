package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}//interface类似Object

type Context struct {
	Writer     http.ResponseWriter //
	Req        *http.Request //request请求，包含所有
	Path       string //当前路径
	Method     string //请求方法
	StatusCode int //请求状态
}

func newContetx(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

//
func (c *Context) JSON(code int, obj interface{}) {
	c.Setheader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer) //Encode获取
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//设置请求头
func (c *Context) Setheader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//解析带的参数后面的参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//获取form指定参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//设置转发头状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//指定状态码，指定要填充的语句，填充语句的参数
//传递string数据
func (c *Context) String(code int, format string, value ...interface{}) {
	c.Setheader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

 //传输字节码数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

 //传输html数据
func (c *Context) HTML(code int, html string) {
	c.Setheader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
