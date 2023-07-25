package router

import (
	"fmt"
	"klebersonromero/github.com/animeapi/cmd/controller"
	"klebersonromero/github.com/animeapi/internal/auth"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Initialize() {
	// Inicializar o roteador
	router := mux.NewRouter()

	//Auth
	router.HandleFunc("/api/v1/token", controller.GenerateToken).Methods("GET")
	// Definir as rotas da API para a tabela anime
	// para permitir apenas com autenticação, coloca a rota sobre a função authenticate()
	router.HandleFunc("/api/v1/animes", auth.Auth(controller.GetAllAnimes)).Methods("GET")
	router.HandleFunc("/api/v1/animes", auth.Auth(controller.GetAllAnimes)).Methods("GET").Queries("page", "{page:[0-9]+}", "perPage", "{perPage:[0-9]+}")

	router.HandleFunc("/api/v1/animes/{id}", auth.Auth(controller.GetAnime)).Methods("GET")
	router.HandleFunc("/api/v1/animes", auth.Auth(controller.CreateAnime)).Methods("POST")
	router.HandleFunc("/api/v1/animes/{id}", auth.Auth(controller.UpdateAnime)).Methods("PUT")
	router.HandleFunc("/api/v1/animes/{id}", auth.Auth(controller.DeleteAnime)).Methods("DELETE")
	router.HandleFunc("/api/v1/animes", auth.Auth(controller.DeleteAllAnimes)).Methods("DELETE")

	// Definir as rotas da API para a tabela episodios
	router.HandleFunc("/api/v1/episodios", auth.Auth(controller.CreateEpisodio)).Methods("POST")
	router.HandleFunc("/api/v1/episodios", auth.Auth(controller.GetAllEpisodios)).Methods("GET")
	router.HandleFunc("/api/v1/episodios/{id}", auth.Auth(controller.UpdateEpisodio)).Methods("PUT")
	router.HandleFunc("/api/v1/episodios/{id}", auth.Auth(controller.DeleteEpisodio)).Methods("DELETE")
	router.HandleFunc("/api/v1/episodios", auth.Auth(controller.DeleteAllEpisodios)).Methods("DELETE")

	// Definir as rotas da API para a tabela temporadas
	router.HandleFunc("/api/v1/temporadas", auth.Auth(controller.CreateTemporada)).Methods("POST")
	router.HandleFunc("/api/v1/temporadas", auth.Auth(controller.GetAllTemporadas)).Methods("GET")
	router.HandleFunc("/api/v1/temporadas/{id}", auth.Auth(controller.GetTemporada)).Methods("GET")
	router.HandleFunc("/api/v1/temporadas/{id}", auth.Auth(controller.UpdateTemporada)).Methods("PUT")
	router.HandleFunc("/api/v1/temporadas/{id}", auth.Auth(controller.DeleteTemporada)).Methods("DELETE")
	router.HandleFunc("/api/v1/temporadas", auth.Auth(controller.DeleteAllTemporadas)).Methods("DELETE")

	// Configurar o middleware CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Iniciar o servidor HTTP com as configurações de CORS
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, corsMiddleware.Handler(router)))
}
