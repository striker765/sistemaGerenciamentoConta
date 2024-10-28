package controllers

import (
	"encoding/json"
	"net/http"
	"sistema/contas/models"
	"strconv"
)

var contas = make(map[int]*models.ContaCorrente)
var nextID = 1

func CriarConta(w http.ResponseWriter, r *http.Request) {
	var novaConta models.ContaCorrente
	if err := json.NewDecoder(r.Body).Decode(&novaConta); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	novaConta.ID = nextID
	nextID++
	contas[novaConta.ID] = &novaConta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novaConta)
}

func ListarContas(w http.ResponseWriter, r *http.Request) {
	var listaContas []*models.ContaCorrente
	for _, conta := range contas {
		listaContas = append(listaContas, conta)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listaContas)
}

func DeletarConta(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || contas[id] == nil {
		http.Error(w, "Conta não encontrada", http.StatusNotFound)
		return
	}
	delete(contas, id)
	w.WriteHeader(http.StatusNoContent)
}
