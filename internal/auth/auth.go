package auth

import (
	"net/http"
	"os"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implemente a lógica de autenticação aqui
		// Por exemplo, verifique o token de autenticação enviado no cabeçalho da solicitação
		authToken := r.Header.Get("Authorization")
		validToken := os.Getenv("AUTH_TOKEN")

		if authToken != validToken {
			// Se o token for inválido, retorne um erro de autenticação
			http.Error(w, "Autenticação falhou", http.StatusUnauthorized)
			return
		}

		// Se a autenticação for bem-sucedida, chame a próxima função de manipulador
		next(w, r)
	}
}
