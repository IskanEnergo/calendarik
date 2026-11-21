package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/IskanEnergo/calendarik/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string //hhtp method
	Path       string // e.g /users/{id}
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(r.Handler, r.Middleware...)
}
