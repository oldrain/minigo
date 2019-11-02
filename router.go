// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import "fmt"

type IGroup interface {
	IRoute

	Group()
}

type IRoute interface {
	Use()

	Post()
}

type Router struct {
	basePath string
	handlers HandlerFuncChain
	routeNames []string
	isAppend bool

	api *Api
}

type RouteInfo struct {
	method string
	path string
	handlers []HandlerFunc
}

func (router *Router) BasePath() string {
	return router.basePath
}

func (router *Router) Group(relativePath string, handlers ...HandlerFunc) *Router {
	return &Router{
		basePath: joinPath(router.basePath, relativePath),
		handlers: router.mergeHandler(handlers),
		api: router.api,
	}
}

func (router *Router) Use(middleware ...HandlerFunc) *Router {
	if router.isAppend == true {
		panic(fmt.Sprintf("Router '%s' used 'LastUse' before.", router.BasePath()))
	}
	router.handlers = router.mergeHandler(middleware)
	return router
}

func (router *Router) LastUse(middleware ...HandlerFunc) *Router {
	if router.isAppend == true {
		panic(fmt.Sprintf("Router '%s' used 'LastUse' before.", router.BasePath()))
	}
	router.isAppend = true
	if len(router.routeNames) == 0 {
		return router
	}
	for _, routeName := range router.routeNames {
		routeInfo := router.api.routes[routeName]
		routeInfo.handlers = append(routeInfo.handlers, middleware...)
	}
	router.handlers = router.mergeHandler(middleware)
	return router
}

func (router *Router) Post(relativePath string, handlers ...HandlerFunc) *Router {
	return router.handle(reqMethodPost, relativePath, handlers)
}

func (router *Router) handle(httpMethod string, relativePath string, handlers HandlerFuncChain) *Router {
	absPath := joinPath(router.basePath, relativePath)
	if nil == handlers {
		panic(fmt.Sprintf("%s need handler", absPath))
	}
	routeInfo := &RouteInfo{
		method: httpMethod,
		path: absPath,
		handlers: router.mergeHandler(handlers),
	}
	router.routeNames = append(router.routeNames, absPath)
	router.api.routes[routeInfo.path] = routeInfo
	logPrintRoute(httpMethod, absPath, handlers)
	return router
}

func (router *Router) mergeHandler(handlers HandlerFuncChain) HandlerFuncChain {
	finalSize := len(router.handlers) + len(handlers)
	if finalSize >= int(maxHandleSize) {
		panic("too many handlers")
	}
	finalHandlers := make(HandlerFuncChain, finalSize)
	copy(finalHandlers, router.handlers)
	copy(finalHandlers[len(router.handlers):], handlers)
	return finalHandlers
}
