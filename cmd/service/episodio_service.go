package service

import (
	"fmt"
	"klebersonromero/github.com/animeapi/cmd/helper"
	"klebersonromero/github.com/animeapi/cmd/model"
)

// Implementação da interface EntityService para a entidade Episodio
type EpisodioService struct{}

// Método GetAll para obter todos os episódios do banco de dados
func (es *EpisodioService) GetAll() ([]interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM episodios")
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
		// Ler os valores da linha atual e armazenar na variável episódio
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
	// Converter a slice de episódios para uma slice de interface{}
	var episodiosInterface []interface{}
	for _, episodio := range episodios {
		episodiosInterface = append(episodiosInterface, episodio)
	}

	return episodiosInterface, nil
}

// Método GetByID para obter um episódio pelo ID do banco de dados
func (es *EpisodioService) GetByID(id int) (interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Executar a consulta SQL
	row := db.QueryRow("SELECT * FROM episodios WHERE id = ?", id)

	// Criar uma variável para armazenar o episódio
	var episodio model.Episodio

	// Ler os valores da linha atual e armazenar na variável episódio
	err = row.Scan(&episodio.ID, &episodio.IDTemporada, &episodio.Numero, &episodio.URLThumb, &episodio.Intro, &episodio.URLVideo)
	if err != nil {
		return nil, err
	}
	return episodio, nil
}

// Método Create para criar um episódio no banco de dados
func (es *EpisodioService) Create(data interface{}) (interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Fazer o cast do parâmetro data para o tipo model.Episodio
	episodio, ok := data.(model.Episodio)
	if !ok {
		return nil, fmt.Errorf("dados inválidos para criar o episódio")
	}

	// Executar a consulta SQL para inserir o episódio no banco de dados
	result, err := db.Exec("INSERT INTO episodios (id_temporada, numero, urlthumb, intro, urlvideo) VALUES (?, ?, ?, ?, ?)",
		episodio.IDTemporada, episodio.Numero, episodio.URLThumb, episodio.Intro, episodio.URLVideo)
	if err != nil {
		return nil, err
	}

	// Obter o ID gerado para o episódio inserido
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return id, nil
}

// Método Update para atualizar um episódio no banco de dados
func (es *EpisodioService) Update(id int, data interface{}) error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Fazer o cast do parâmetro data para o tipo model.Episodio
	episodio, ok := data.(model.Episodio)
	if !ok {
		return fmt.Errorf("dados inválidos para atualizar o episódio")
	}

	// Executar a consulta SQL para atualizar o episódio no banco de dados
	_, err = db.Exec("UPDATE episodios SET id_temporada = ?, numero = ?, urlthumb = ?, intro = ?, urlvideo = ? WHERE id = ?",
		episodio.IDTemporada, episodio.Numero, episodio.URLThumb, episodio.Intro, episodio.URLVideo, id)
	if err != nil {
		return err
	}

	return nil
}

// Método Delete para deletar um episódio do banco de dados
func (es *EpisodioService) Delete(id int) error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Executar a consulta SQL para deletar o episódio no banco de dados
	_, err = db.Exec("DELETE FROM episodios WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Método DeleteAll para deletar todos os episódios do banco de dados
func (es *EpisodioService) DeleteAll() error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Executar a consulta SQL para deletar todos os episódios no banco de dados
	_, err = db.Exec("DELETE FROM episodios")
	if err != nil {
		return err
	}

	return nil
}
