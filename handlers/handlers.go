package handlers

import "net/http"

type Middleware = func(http.Handler) http.Handler

func WrapHandlerFunc(handlerFunc http.HandlerFunc, middlewares ...Middleware) http.Handler {
	var handler http.Handler = handlerFunc
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
