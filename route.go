package router

import (
	"net/http"

	"github.com/gouniverse/responses"
)

type RouteInterface interface {
	GetHandler() func(w http.ResponseWriter, r *http.Request)
	SetHandler(handler func(w http.ResponseWriter, r *http.Request))
	GetName() string
	SetName(name string)
	GetMethods() []string
	SetMethods(methods []string)
	GetMiddlewares() []Middleware
	SetMiddlewares(middlewares []Middleware)
	GetPath() string
	SetPath(path string)

	AddMiddlewares(middlewares ...Middleware)
	PrependMiddlewares(middlewares ...Middleware)
}

type Route struct {
	// Domain      string
	Path        string
	Methods     []string // optional, default all methods
	Handler     func(w http.ResponseWriter, r *http.Request)
	HTMLHandler func(w http.ResponseWriter, r *http.Request) string
	JSONHandler func(w http.ResponseWriter, r *http.Request) string
	Middlewares []Middleware
	Name        string // optional, default empty string
}

var _ RouteInterface = (*Route)(nil)

func (route Route) String() string {
	return route.Path
}

func (route Route) GetHandler() func(w http.ResponseWriter, r *http.Request) {
	if route.Handler != nil {
		return route.Handler
	}

	if route.HTMLHandler != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			responses.HTMLResponse(w, r, route.HTMLHandler(w, r))
		}
	}

	if route.JSONHandler != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			responses.JSONResponse(w, r, route.JSONHandler(w, r))
		}
	}

	return nil
}

func (route *Route) SetHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	route.Handler = handler
}

func (route *Route) GetName() string {
	return route.Name
}

func (route *Route) SetName(name string) {
	route.Name = name
}

func (route *Route) GetMethods() []string {
	if len(route.Methods) == 0 {
		return []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		}
	}

	if len(route.Methods) == 1 && route.Methods[0] == "all" {
		return []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		}
	}

	return route.Methods
}

func (route *Route) SetMethods(methods []string) {
	route.Methods = methods
}

func (route Route) GetMiddlewares() []Middleware {
	return route.Middlewares
}

func (route *Route) SetMiddlewares(middlewares []Middleware) {
	route.Middlewares = middlewares
}

func (route *Route) GetPath() string {
	return route.Path
}

func (route *Route) SetPath(path string) {
	route.Path = path
}

func (route *Route) AddMiddlewares(middlewares ...Middleware) {
	route.Middlewares = append(route.Middlewares, middlewares...)
}

func (route *Route) PrependMiddlewares(middlewares ...Middleware) {
	route.Middlewares = append(middlewares, route.Middlewares...)
}

func (route Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := route.GetHandler()

	if handler != nil {
		handler(w, r)
	}

	w.WriteHeader(http.StatusInternalServerError)
}
