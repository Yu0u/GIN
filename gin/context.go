package gin

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sync"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Request *http.Request

	Path string
	Method string
	Params map[string]string
	Keys map[string]interface{}

	StatusCode int


	handlers []HandlerFunc
	index int8

	mu  sync.RWMutex
}

const abortIndex = math.MaxInt8 / 2

func newContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer: writer,
		Request: request,
		Path: request.URL.Path,
		Method: request.Method,
		index: -1,
	}
}

func (c *Context) Next()  {
	c.index++
	for c.index < int8(len(c.handlers)){
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = int8(len(c.handlers))
	c.JSON(code, H{"message": err})
}

func (c *Context) PostForm(key string) string{
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

func (c *Context) Param(key string) string {
	value,_ := c.Params[key]
	return value
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *Context) AbortWithStatusJSON(code int,jsonObj interface{}){
	c.Abort()
	c.JSON(code,jsonObj)
}

func (c *Context) Set(key string,value interface{})  {
	c.mu.Lock()
	if c.Keys == nil{
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
	c.mu.Unlock()
}

func (c *Context) Get(key string)(value interface{},exists bool){
	c.mu.RLock()
	value,exists = c.Keys[key]
	c.mu.RUnlock()
	return
}
