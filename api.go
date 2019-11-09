// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import (
	"net/http"
	"sync"
)

type HandlerFunc func(ctx *Context)
type HandlerFuncChain []HandlerFunc

type Api struct {
	Router
	routes map[string]*RouteInfo

	ctxPool sync.Pool
}

func Default() *Api {
	api := New()
	api.Use(Logger(), Recovery())
	return api
}

func New() *Api {
	logStart()
	api := &Api{
		Router: Router{
			basePath: "/",
			handlers: nil,
		},
		routes: make(map[string]*RouteInfo),
	}

	api.Router.api = api

	api.ctxPool.New = func() interface{} {
		return api.newContext()
	}

	return api
}

func (api *Api) Run(addr string) (err error) {
	logRun(addr)
	err = http.ListenAndServe(addr, api)
	return
}

func (api *Api) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := api.ctxPool.Get().(*Context)
	ctx.reset()
	ctx.Input.Request = req
	ctx.Output.Response = res
	api.handleHTTPRequest(ctx)
	api.ctxPool.Put(ctx)
}

func (api *Api) Use(middleware ...HandlerFunc) *Router {
	return api.Router.Use(middleware...)
}

func (api *Api) handleHTTPRequest(ctx *Context) {
	httpMethod := ctx.Method()

	if httpMethod != reqMethodPost {
		ctx.AbortWithError(code405, "Unsupported http method")
		return
	}
	if ctx.GetInHeader(contentTypeHeader) != inputContentTypeJson {
		ctx.AbortWithError(code400, "Unsupported content type")
		return
	}

	path := ctx.Path()

	if _, ok := api.routes[path]; !ok {
		ctx.AbortWithError(code404, msg404)
		return
	}

	routeInfo := api.routes[path]
	ctx.handlers = routeInfo.handlers
	ctx.Continue()
}

func (api *Api) newContext() *Context {
	return &Context{}
}
