package controller

import (
	"database/sql"
	"encoding/json"
	entity "klebersonromero/github.com/animeapi/internal/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Função para obter todos os episódios
func GetEpisodios(w http.ResponseWriter, r *http.Request) {
	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM episodios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Criar uma slice para armazenar os episódios
	episodios := []entity.Episodio{}

	// Iterar sobre as linhas do resultado da consulta
	for rows.Next() {
		// Criar uma variável para armazenar cada episódio
		var episodio entity.Episodio
		// Ler os valores da linha atual e armazenar na variável episódio
		err := rows.Scan(&episodio.ID, &episodio.IDAnime, &episodio.Temporada, &episodio.Episodio, &episodio.URLThumb, &episodio.Intro, &episodio.URLVideo)
		if err != nil {
			log.Fatal(err)
		}
		// Adicionar o episódio à slice de episódios
		episodios = append(episodios, episodio)
	}

	// Verificar se houve algum erro durante a iteração das linhas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar a slice de episódios para o formato JSON
	json.NewEncoder(w).Encode(episodios)
}

// Função para criar um episódio
func CreateEpisodio(w http.ResponseWriter, r *http.Request) {
	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Decodificar o corpo da requisição JSON em uma variável episódio
	var episodio entity.Episodio
	err = json.NewDecoder(r.Body).Decode(&episodio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Executar a consulta SQL para inserir o episódio no banco de dados
	_, err = db.Exec("INSERT INTO episodios (id_anime, temporada, episodio, urlthumb, intro, urlvideo) VALUES (?, ?, ?, ?, ?, ?)",
		episodio.IDAnime, episodio.Temporada, episodio.Episodio, episodio.URLThumb, episodio.Intro, episodio.URLVideo)
	if err != nil {
		log.Fatal(err)
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Episódio criado com sucesso")
}

// Função para atualizar um episódio
func UpdateEpisodio(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro ID da URL
	params := mux.Vars(r)
	id := params["id"]

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Decodificar o corpo da requisição JSON em uma variável episódio
	var episodio entity.Episodio
	err = json.NewDecoder(r.Body).Decode(&episodio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Executar a consulta SQL para atualizar o episódio no banco de dados
	_, err = db.Exec("UPDATE episodios SET id_anime = ?, temporada = ?, episodio = ?, urlthumb = ?, intro = ?, urlvideo = ? WHERE id = ?",
		episodio.IDAnime, episodio.Temporada, episodio.Episodio, episodio.URLThumb, episodio.Intro, episodio.URLVideo, id)
	if err != nil {
		log.Fatal(err)
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

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Executar a consulta SQL para deletar o episódio no banco de dados
	_, err = db.Exec("DELETE FROM episodios WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Episódio deletado com sucesso")
}
