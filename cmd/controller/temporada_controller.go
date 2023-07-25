package controller

import (
	"encoding/json"
	"klebersonromero/github.com/animeapi/cmd/model"
	"klebersonromero/github.com/animeapi/cmd/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Função para obter todas as temporadas
func GetAllTemporadas(w http.ResponseWriter, r *http.Request) {
	temporadaService := &service.TemporadaService{}
	temporadas, err := temporadaService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar a slice de temporadas para o formato JSON
	json.NewEncoder(w).Encode(temporadas)
}

// Função para obter uma temporada pelo ID
func GetTemporada(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	temporadaID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	temporadaService := &service.TemporadaService{}
	temporada, err := temporadaService.GetByID(temporadaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar a temporada para o formato JSON
	json.NewEncoder(w).Encode(temporada)
}

// Função para criar uma temporada
func CreateTemporada(w http.ResponseWriter, r *http.Request) {
	// Decodificar o corpo da requisição JSON em uma variável temporada
	var temporada model.Temporada
	err := json.NewDecoder(r.Body).Decode(&temporada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	temporadaService := &service.TemporadaService{}
	id, err := temporadaService.Create(temporada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar o ID da temporada criada no formato JSON
	json.NewEncoder(w).Encode(id)
}

// Função para atualizar uma temporada
func UpdateTemporada(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	temporadaID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decodificar o corpo da requisição JSON em uma variável temporada
	var temporada model.Temporada
	err = json.NewDecoder(r.Body).Decode(&temporada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	temporadaService := &service.TemporadaService{}
	err = temporadaService.Update(temporadaID, temporada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Temporada atualizada com sucesso")
}

// Função para deletar uma temporada
func DeleteTemporada(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Converter o ID para inteiro
	temporadaID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	temporadaService := &service.TemporadaService{}
	err = temporadaService.Delete(temporadaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Temporada deletada com sucesso")
}

// Função para deletar todas as temporadas
func DeleteAllTemporadas(w http.ResponseWriter, r *http.Request) {
	temporadaService := &service.TemporadaService{}
	err := temporadaService.DeleteAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Todas as temporadas foram deletadas com sucesso")
}
