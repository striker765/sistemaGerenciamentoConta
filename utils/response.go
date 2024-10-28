// utils/response.go
package utils

import (
	"encoding/json"
	"net/http"
)

// Resposta é uma estrutura para encapsular respostas da API
type Resposta struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// RespondJSON envia uma resposta JSON
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
