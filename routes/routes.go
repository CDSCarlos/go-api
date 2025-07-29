package routes

import (
	"go-api/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pessoas", controllers.GetPessoas).Methods("GET")
	r.HandleFunc("/pessoas/{id}", controllers.GetPessoa).Methods("GET")
	r.HandleFunc("/pessoas", controllers.CreatePessoa).Methods("POST")
	r.HandleFunc("/pessoas/{id}", controllers.UpdatePessoa).Methods("PUT")
	r.HandleFunc("/pessoas/{id}", controllers.DeletePessoa).Methods("DELETE")
	return r
}
