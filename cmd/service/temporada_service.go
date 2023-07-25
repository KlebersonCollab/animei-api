package service

import (
	"fmt"
	"klebersonromero/github.com/animeapi/cmd/helper"
	"klebersonromero/github.com/animeapi/cmd/model"
)

// Implementação da interface EntityService para a entidade Temporada
type TemporadaService struct{}

// Método GetAll para obter todas as temporadas do banco de dados
func (ts *TemporadaService) GetAll() ([]interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Executar a consulta SQL
	rows, err := db.Query("SELECT * FROM temporadas")
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

	// Converter a slice de temporadas para uma slice de interface{}
	var temporadasInterface []interface{}
	for _, temporada := range temporadas {
		temporadasInterface = append(temporadasInterface, temporada)
	}

	return temporadasInterface, nil
}

// Método GetByID para obter uma temporada pelo ID do banco de dados
func (ts *TemporadaService) GetByID(id int) (interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Executar a consulta SQL
	row := db.QueryRow("SELECT * FROM temporadas WHERE id = ?", id)

	// Criar uma variável para armazenar a temporada
	var temporada model.Temporada

	// Ler os valores da linha atual e armazenar na variável temporada
	err = row.Scan(&temporada.ID, &temporada.IDAnime, &temporada.Numero)
	if err != nil {
		return nil, err
	}

	// Obter os episódios da temporada
	temporada.Episodios, err = getEpisodiosByTemporadaID(db, temporada.ID)
	if err != nil {
		return nil, err
	}

	return temporada, nil
}

// Método Create para criar uma temporada no banco de dados
func (ts *TemporadaService) Create(data interface{}) (interface{}, error) {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Fazer o cast do parâmetro data para o tipo model.Temporada
	temporada, ok := data.(model.Temporada)
	if !ok {
		return nil, fmt.Errorf("dados inválidos para criar a temporada")
	}

	// Executar a consulta SQL para inserir a temporada no banco de dados
	result, err := db.Exec("INSERT INTO temporadas (id_anime, numero) VALUES (?, ?)",
		temporada.IDAnime, temporada.Numero)
	if err != nil {
		return nil, err
	}

	// Obter o ID gerado para a temporada inserida
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return id, nil
}

// Método Update para atualizar uma temporada no banco de dados
func (ts *TemporadaService) Update(id int, data interface{}) error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Fazer o cast do parâmetro data para o tipo model.Temporada
	temporada, ok := data.(model.Temporada)
	if !ok {
		return fmt.Errorf("dados inválidos para atualizar a temporada")
	}

	// Executar a consulta SQL para atualizar a temporada no banco de dados
	_, err = db.Exec("UPDATE temporadas SET id_anime = ?, numero = ? WHERE id = ?",
		temporada.IDAnime, temporada.Numero, id)
	if err != nil {
		return err
	}

	return nil
}

// Método Delete para deletar uma temporada do banco de dados
func (ts *TemporadaService) Delete(id int) error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Executar a consulta SQL para deletar a temporada no banco de dados
	_, err = db.Exec("DELETE FROM temporadas WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Método DeleteAll para deletar todas as temporadas do banco de dados
func (ts *TemporadaService) DeleteAll() error {
	// Abrir conexão com o banco de dados
	db, err := helper.OpenDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Executar a consulta SQL para deletar todas as temporadas no banco de dados
	_, err = db.Exec("DELETE FROM temporadas")
	if err != nil {
		return err
	}

	return nil
}
