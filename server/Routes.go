package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Cameras", "/cameras", "GET", CameraIndex},
	Route{"Report", "/report", "GET", CameraIndex},
	Route{"CameraImage", "/image/{camera_path}", "GET", CameraImage},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
