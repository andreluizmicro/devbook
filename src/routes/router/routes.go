package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri            string
	Method         string
	Controller     func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Controller).Methods(route.Method)
	}

	return r
}
