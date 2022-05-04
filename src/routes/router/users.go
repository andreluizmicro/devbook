package router

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:            "/users",
		Method:         http.MethodGet,
		Controller:     controllers.Index,
		Authentication: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Controller:     controllers.Create,
		Authentication: false,
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodGet,
		Controller:     controllers.SearchById,
		Authentication: false,
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodPut,
		Controller:     controllers.Update,
		Authentication: false,
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodDelete,
		Controller:     controllers.Delete,
		Authentication: false,
	},
}
