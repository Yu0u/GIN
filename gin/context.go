package gin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Request *http.Request

	Path string
	Method string

	StatusCode int
}

func newContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer: writer,
		Request: request,
		Path: request.URL.Path,
		Method: request.Method,
	}
}

func (c *Context) PostFrom(key string) string{
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string,value string) {
	c.Writer.Header().Set(key,value)
}

func (c *Context) String(code int,format string,value...interface{}){
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}

func (c *Context) JSON(code int,obj interface{}){
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil{
		http.Error(c.Writer,err.Error(),500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
