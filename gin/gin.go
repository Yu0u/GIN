package gin

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

//FIXME:多个结构体别这么定义吧，不太美观。
type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		engine      *Engine
	}
	Engine struct {
		router *router
		*RouterGroup
		// FIXME:这写法蛮奇怪的，会有更好的方法嘛？
		//groups []*RouterGroup

	}
)

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	//	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		engine: engine,
	}
	//	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := r.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	r.engine.router.addRoute(method, pattern, handler)
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

func (r *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	r.addRoute("PUT", pattern, handler)
}

func (r *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	r.addRoute("DELETE", pattern, handler)
}

func (r *RouterGroup) Use(middlewares ...HandlerFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//	var middlewares []HandlerFunc
	//	for _, group := range e.groups {
	//		if strings.HasPrefix(request.URL.Path, group.prefix) {
	//			middlewares = append(middlewares, group.middlewares...)
	//		}
	//	}
	c := newContext(writer, request)
	//	c.handlers = middlewares
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
