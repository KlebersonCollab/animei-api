package service

import (
	"database/sql"
	"fmt"
	"klebersonromero/github.com/animeapi/cmd/helper"
	"klebersonromero/github.com/animeapi/cmd/model"
)

// Implementação da interface EntityService para a entidade Anime
type AnimeService struct{}

// Método GetAll para obter todos os animes do banco de dados
func (as *AnimeService) GetAll() ([]interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM animes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar uma slice para armazenar os animes
	animes := []model.Anime{}

	// Iterar sobre as linhas do resultado da consulta
	for rows.Next() {
		// Criar uma variável para armazenar cada anime
		var anime model.Anime
		// Ler os valores da linha atual e armazenar na variável anime
		err := rows.Scan(&anime.ID, &anime.Nome, &anime.URLCapa, &anime.Tipo, &anime.Intro, &anime.URLBanner, &anime.Nota)
		if err != nil {
			return nil, err
		}

		// Obter as temporadas do anime
		anime.Temporadas, err = getTemporadasByAnimeID(db, anime.ID)
		if err != nil {
			return nil, err
		}

		// Adicionar o anime à slice de animes
		animes = append(animes, anime)
	}

	// Verificar se houve algum erro durante a iteração das linhas
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// Converter a slice de animes para uma slice de interface{}
	animesInterface := make([]interface{}, len(animes))
	for i, anime := range animes {
		animesInterface[i] = anime
	}

	return animesInterface, nil
}

// Método GetByID para obter um anime pelo ID do banco de dados
func (as *AnimeService) GetByID(id int) (interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Executar a consulta SQL
	row := db.QueryRow("SELECT * FROM animes WHERE id = ?", id)

	// Criar uma variável para armazenar o anime
	var anime model.Anime

	// Ler os valores da linha atual e armazenar na variável anime
	err = row.Scan(&anime.ID, &anime.Nome, &anime.URLCapa, &anime.Tipo, &anime.Intro, &anime.URLBanner, &anime.Nota)
	if err != nil {
		return nil, err
	}

	// Obter as temporadas do anime
	anime.Temporadas, err = getTemporadasByAnimeID(db, anime.ID)
	if err != nil {
		return nil, err
	}

	return anime, nil
}

// Método Create para criar um anime no banco de dados
func (as *AnimeService) Create(data interface{}) (interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Fazer o cast do parâmetro data para o tipo model.Anime
	anime, ok := data.(model.Anime)
	if !ok {
		return nil, fmt.Errorf("dados inválidos para criar o anime")
	}

	// Executar a consulta SQL para inserir o anime no banco de dados
	result, err := db.Exec("INSERT INTO animes (nome, urlcapa, tipo, intro, urlbanner, nota) VALUES (?, ?, ?, ?, ?, ?)",
		anime.Nome, anime.URLCapa, anime.Tipo, anime.Intro, anime.URLBanner, anime.Nota)
	if err != nil {
		return nil, err
	}

	// Obter o ID gerado para o anime inserido
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return id, nil
}

// Método Update para atualizar um anime no banco de dados
func (as *AnimeService) Update(id int, data interface{}) error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Fazer o cast do parâmetro data para o tipo model.Anime
	anime, ok := data.(model.Anime)
	if !ok {
		return fmt.Errorf("dados inválidos para atualizar o anime")
	}

	// Executar a consulta SQL para atualizar o anime no banco de dados
	_, err = db.Exec("UPDATE animes SET nome = ?, urlcapa = ?, tipo = ?, intro = ?, urlbanner = ?, nota = ? WHERE id = ?",
		anime.Nome, anime.URLCapa, anime.Tipo, anime.Intro, anime.URLBanner, anime.Nota, id)
	if err != nil {
		return err
	}

	return nil
}

// Método Delete para deletar um anime do banco de dados
func (as *AnimeService) Delete(id int) error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Executar a consulta SQL para deletar o anime no banco de dados
	_, err = db.Exec("DELETE FROM animes WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Método DeleteAll para deletar todos os animes do banco de dados
func (as *AnimeService) DeleteAll() error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Executar a consulta SQL para deletar todos os animes no banco de dados
	_, err = db.Exec("DELETE FROM animes")
	if err != nil {
		return err
	}

	return nil
}

// Função auxiliar para obter as temporadas de um anime pelo ID do anime
func getTemporadasByAnimeID(db *sql.DB, animeID int) ([]model.Temporada, error) {
	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM temporadas WHERE id_anime = ?", animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar uma slice para armazenar as temporadas
	temporadas := []model.Temporada{}

	// Iterar sobre as linhas do resultado da consulta
	for rows.Next() {
		// Criar uma variável para armazenar cada temporada
		var temporada model.Temporada
		// Ler os valores da linha atual e armazenar na variável temporada
		err := rows.Scan(&temporada.ID, &temporada.IDAnime, &temporada.Numero)
		if err != nil {
			return nil, err
		}

		// Obter os episódios da temporada
		temporada.Episodios, err = getEpisodiosByTemporadaID(db, temporada.ID)
		if err != nil {
			return nil, err
		}

		// Adicionar a temporada à slice de temporadas
		temporadas = append(temporadas, temporada)
	}

	// Verificar se houve algum erro durante a iteração das linhas
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return temporadas, nil
}

// Função auxiliar para obter os episódios de uma temporada pelo ID da temporada
func getEpisodiosByTemporadaID(db *sql.DB, temporadaID int) ([]model.Episodio, error) {
	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM episodios WHERE id_temporada = ?", temporadaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar uma slice para armazenar os episódios
	episodios := []model.Episodio{}

	// Iterar sobre as linhas do resultado da consulta
	for rows.Next() {
		// Criar uma variável para armazenar cada episódio
		var episodio model.Episodio
		// Ler os valores da linha atual e armazenar na variável episodio
		err := rows.Scan(&episodio.ID, &episodio.IDTemporada, &episodio.Numero, &episodio.URLThumb, &episodio.Intro, &episodio.URLVideo)
		if err != nil {
			return nil, err
		}

		// Adicionar o episódio à slice de episódios
		episodios = append(episodios, episodio)
	}

	// Verificar se houve algum erro durante a iteração das linhas
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return episodios, nil
}
