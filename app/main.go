package main

import (
	"log"
	"net/http"

	"klebersonromero/github.com/animeapi/internal/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implemente a lógica de autenticação aqui
		// Por exemplo, verifique o token de autenticação enviado no cabeçalho da solicitação
		authToken := r.Header.Get("Authorization")
		validToken := "myAuthToken" // Substitua pelo seu token de autenticação válido

		if authToken != validToken {
			// Se o token for inválido, retorne um erro de autenticação
			http.Error(w, "Autenticação falhou", http.StatusUnauthorized)
			return
		}

		// Se a autenticação for bem-sucedida, chame a próxima função de manipulador
		next(w, r)
	}
}

func main() {
	// Cria as tabelas
	controller.CreateTables()
	// Inicializar o roteador
	router := mux.NewRouter()

	// Definir as rotas da API para a tabela anime
	// para permitir apenas com autenticação coloca a roda sobre a função authenticate()
	router.HandleFunc("/api/v1/animes", authenticate(controller.GetAnimes)).Methods("GET")
	router.HandleFunc("/api/v1/animes/{id}", authenticate(controller.GetAnime)).Methods("GET")
	router.HandleFunc("/api/v1/animes", authenticate(controller.CreateAnime)).Methods("POST")
	router.HandleFunc("/api/v1/animes/{id}", authenticate(controller.UpdateAnime)).Methods("PUT")
	router.HandleFunc("/api/v1/animes/{id}", authenticate(controller.DeleteAnime)).Methods("DELETE")

	// Definir as rotas da API para a tabela episodios
	router.HandleFunc("/api/v1/episodios", authenticate(controller.CreateEpisodio)).Methods("POST")
	router.HandleFunc("/api/v1/episodios", authenticate(controller.GetEpisodios)).Methods("GET")
	router.HandleFunc("/api/v1/episodios/{id}", authenticate(controller.UpdateEpisodio)).Methods("PUT")
	router.HandleFunc("/api/v1/episodios/{id}", authenticate(controller.DeleteEpisodio)).Methods("DELETE")

	// Configurar o middleware CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Configurar o servidor HTTP na porta 8000
	// Iniciar o servidor HTTP com as configurações de CORS
	log.Fatal(http.ListenAndServe(":8000", corsMiddleware.Handler(router)))
}
