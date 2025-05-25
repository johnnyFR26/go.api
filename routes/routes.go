package routes

import (
	"github.com/gorilla/mux"
	"github.com/johnnyFR26/go.api/app/controllers"
	"github.com/johnnyFR26/go.api/config"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	config.ConnectDatabase()

	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	return r
}
