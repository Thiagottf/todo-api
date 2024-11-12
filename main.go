// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seu-usuario/todo-api/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
