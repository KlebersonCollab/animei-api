package controller

import (
	"encoding/json"
	"klebersonromero/github.com/animeapi/cmd/model"
	"klebersonromero/github.com/animeapi/cmd/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Função para obter todos os episódios
func GetAllEpisodios(w http.ResponseWriter, r *http.Request) {
	episodioService := &service.EpisodioService{}
	episodios, err := episodioService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar a slice de episódios para o formato JSON
	json.NewEncoder(w).Encode(episodios)
}

// Função para obter um episódio pelo ID
func GetEpisodio(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	episodioID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	episodioService := &service.EpisodioService{}
	episodio, err := episodioService.GetByID(episodioID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar o episódio para o formato JSON
	json.NewEncoder(w).Encode(episodio)
}

// Função para criar um episódio
func CreateEpisodio(w http.ResponseWriter, r *http.Request) {
	// Decodificar o corpo da requisição JSON em uma variável episódio
	var episodio model.Episodio
	err := json.NewDecoder(r.Body).Decode(&episodio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	episodioService := &service.EpisodioService{}
	id, err := episodioService.Create(episodio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar o ID do episódio criado no formato JSON
	json.NewEncoder(w).Encode(id)
}

// Função para atualizar um episódio
func UpdateEpisodio(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	episodioID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decodificar o corpo da requisição JSON em uma variável episódio
	var episodio model.Episodio
	err = json.NewDecoder(r.Body).Decode(&episodio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	episodioService := &service.EpisodioService{}
	err = episodioService.Update(episodioID, episodio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Episódio atualizado com sucesso")
}

// Função para deletar um episódio
func DeleteEpisodio(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	episodioID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	episodioService := &service.EpisodioService{}
	err = episodioService.Delete(episodioID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Episódio deletado com sucesso")
}

// Função para deletar todos os episódios
func DeleteAllEpisodios(w http.ResponseWriter, r *http.Request) {
	episodioService := &service.EpisodioService{}
	err := episodioService.DeleteAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Todos os episódios foram deletados com sucesso")
}
