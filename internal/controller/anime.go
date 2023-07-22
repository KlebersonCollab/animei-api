package controller

import (
	"database/sql"
	"encoding/json"
	entity "klebersonromero/github.com/animeapi/internal/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Função para obter todos os animes
func GetAnimes(w http.ResponseWriter, r *http.Request) {
	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM anime")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Criar uma slice para armazenar os animes
	animes := []entity.Anime{}

	// Iterar sobre as linhas do resultado da consulta
	for rows.Next() {
		// Criar uma variável para armazenar cada anime
		var anime entity.Anime
		// Ler os valores da linha atual e armazenar na variável anime
		err := rows.Scan(&anime.ID, &anime.Nome, &anime.URLCapa, &anime.Tipo, &anime.Intro, &anime.URLBanner, &anime.Nota)
		if err != nil {
			log.Fatal(err)
		}
		// Adicionar o anime à slice de animes
		animes = append(animes, anime)
	}

	// Verificar se houve algum erro durante a iteração das linhas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
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

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Executar a consulta SQL
	row := db.QueryRow("SELECT * FROM anime WHERE id = ?", id)

	// Criar uma variável para armazenar o anime
	var anime entity.Anime

	// Ler os valores da linha atual e armazenar na variável anime
	err = row.Scan(&anime.ID, &anime.Nome, &anime.URLCapa, &anime.Tipo, &anime.Intro, &anime.URLBanner, &anime.Nota)
	if err != nil {
		log.Fatal(err)
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Serializar o anime para o formato JSON
	json.NewEncoder(w).Encode(anime)
}

// Função para criar um anime
func CreateAnime(w http.ResponseWriter, r *http.Request) {
	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Decodificar o corpo da requisição JSON em uma variável anime
	var anime entity.Anime
	err = json.NewDecoder(r.Body).Decode(&anime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Executar a consulta SQL para inserir o anime no banco de dados
	result, err := db.Exec("INSERT INTO anime (nome, urlcapa, tipo, intro, urlbanner, nota) VALUES (?, ?, ?, ?, ?, ?)",
		anime.Nome, anime.URLCapa, anime.Tipo, anime.Intro, anime.URLBanner, anime.Nota)
	if err != nil {
		log.Fatal(err)
	}

	// Obter o ID gerado para o anime inserido
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
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

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Decodificar o corpo da requisição JSON em uma variável anime
	var anime entity.Anime
	err = json.NewDecoder(r.Body).Decode(&anime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Executar a consulta SQL para atualizar o anime no banco de dados
	_, err = db.Exec("UPDATE anime SET nome = ?, urlcapa = ?, tipo = ?, intro = ?, urlbanner = ?, nota = ? WHERE id = ?",
		anime.Nome, anime.URLCapa, anime.Tipo, anime.Intro, anime.URLBanner, anime.Nota, id)
	if err != nil {
		log.Fatal(err)
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

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Executar a consulta SQL para deletar o anime no banco de dados
	_, err = db.Exec("DELETE FROM anime WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar uma resposta de sucesso no formato JSON
	json.NewEncoder(w).Encode("Anime deletado com sucesso")
}
