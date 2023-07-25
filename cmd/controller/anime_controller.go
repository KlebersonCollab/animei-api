package controller

import (
	"encoding/json"
	"klebersonromero/github.com/animeapi/cmd/model"
	"klebersonromero/github.com/animeapi/cmd/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Função para obter todos os animes
func GetAllAnimes(w http.ResponseWriter, r *http.Request) {
	animeService := &service.AnimeService{}
	animes, err := animeService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar a slice de animes para o formato JSON
	json.NewEncoder(w).Encode(animes)
}

// Função para obter um anime pelo ID
func GetAnime(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]
	// Converter o ID para inteiro
	animeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Instancia o Service e chama o metodo GetByID
	animeService := &service.AnimeService{}
	anime, err := animeService.GetByID(animeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Serializar o anime para o formato JSON
	json.NewEncoder(w).Encode(anime)
}

// Função para criar um anime
func CreateAnime(w http.ResponseWriter, r *http.Request) {
	// Decodificar o corpo da requisição JSON em uma variável anime
	var anime model.Anime
	err := json.NewDecoder(r.Body).Decode(&anime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	animeService := &service.AnimeService{}
	id, err := animeService.Create(anime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar o ID do anime criado no formato JSON
	json.NewEncoder(w).Encode(id)
}

// Função para atualizar um anime
func UpdateAnime(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Decodificar o corpo da requisição JSON em uma variável anime
	var anime model.Anime
	err := json.NewDecoder(r.Body).Decode(&anime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Converter o ID para inteiro
	animeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//
	animeService := &service.AnimeService{}
	err = animeService.Update(animeID, anime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Anime atualizado com sucesso")
}

// Função para deletar um anime
func DeleteAnime(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	animeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	animeService := &service.AnimeService{}
	err = animeService.Delete(animeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Anime deletado com sucesso")
}

// Função para deletar todos os animes
func DeleteAllAnimes(w http.ResponseWriter, r *http.Request) {
	animeService := &service.AnimeService{}
	err := animeService.DeleteAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Todos os animes foram deletados com sucesso")
}
