package main

import (
	"github.com/gorilla/mux"
)

func BaseRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
