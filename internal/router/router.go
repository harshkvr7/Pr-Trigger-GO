package router

import (
	"pr-trigger-go/internal/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/test", controller.SendGreeting).Methods("GET")

	return r
}
