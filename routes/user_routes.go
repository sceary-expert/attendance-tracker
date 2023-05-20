package routes

import (
	"attendance-tracker/controller"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user", controller.CreateUser()).Methods("POST")
}
