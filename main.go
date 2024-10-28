package main

import (
	"net/http"
	"sistema/contas/routes"
)

func main() {
	router := routes.ConfigurarRotas()
	http.ListenAndServe(":8080", router)
}
