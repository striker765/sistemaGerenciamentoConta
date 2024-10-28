package routes

import (
	"sistema/contas/controllers"

	"github.com/gorilla/mux"
)

func ConfigurarRotas() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/contas", controllers.CriarConta).Methods("POST")
	router.HandleFunc("/contas", controllers.ListarContas).Methods("GET")
	router.HandleFunc("/contas", controllers.DeletarConta).Methods("DELETE")
	return router
}
