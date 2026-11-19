package core_http_server

import "net/http"

type Route struct {
	Method  string //hhtp method
	Path    string // e.g /users/{id}
	Handler http.HandlerFunc
}

func NewRoute(method string, path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
