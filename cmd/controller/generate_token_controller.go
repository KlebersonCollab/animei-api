// Arquivo controller/auth_controller.go
package controller

import (
	"encoding/json"
	"klebersonromero/github.com/animeapi/internal/auth"
	"net/http"
)

// Função para gerar e retornar o token JWT
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GenerateToken()
	if err != nil {
		http.Error(w, "Falha ao gerar o token JWT", http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Retornar o token gerado no formato JSON
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
