package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Função para gerar um token JWT
func GenerateToken() (string, error) {
	// Definir a chave secreta para assinar o token
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("Chave secreta JWT não configurada")
	}

	// Definir as informações do token
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
		// Adicione aqui outros dados que você queira incluir no token
	}

	// Criar o token com as informações e a assinatura
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implemente a lógica de autenticação aqui
		// Por exemplo, verifique o token de autenticação enviado no cabeçalho da solicitação
		authHeader := r.Header.Get("Authorization")
		//validToken := os.Getenv("AUTH_TOKEN")
		/*
			if authHeader != validToken {
				// Se o token for inválido, retorne um erro de autenticação
				http.Error(w, "Autenticação falhou", http.StatusUnauthorized)
				return
			}
		*/
		// Verificar se o token está no formato correto (Bearer <token>)
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Token de autenticação inválido", http.StatusUnauthorized)
			return
		}
		authTokenJWT := tokenParts[1]
		secretKey := os.Getenv("JWT_SECRET_KEY")
		// Verificar a validade e autenticidade do token
		token, err := jwt.Parse(authTokenJWT, func(token *jwt.Token) (interface{}, error) {
			// Certifique-se de usar o mesmo algoritmo de assinatura que foi usado para gerar o token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Método de assinatura inesperado: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Autenticação falhou", http.StatusUnauthorized)
			return
		}
		// Se a autenticação for bem-sucedida, chame a próxima função de manipulador
		next(w, r)
	}
}
