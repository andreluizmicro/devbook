package routes

import (
	"api/src/routes/router"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return router.Config(r)
}
