package router

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router{
	// Get the main router
	r := mux.NewRouter()
	return r
}