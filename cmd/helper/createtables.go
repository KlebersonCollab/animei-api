// Arquivo helper/database.go
package helper

import (
	"log"
)

func CreateTables() {
	// Abrir conexão com o banco de dados
	db, err := OpenDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criação da tabela Animes
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS animes (
		id INT PRIMARY KEY AUTO_INCREMENT,
		nome VARCHAR(255),
		urlcapa VARCHAR(255),
		tipo VARCHAR(255),
		intro TEXT,
		urlbanner VARCHAR(255),
		nota DECIMAL(5,2)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Criação da tabela Temporadas
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS temporadas (
		id INT PRIMARY KEY AUTO_INCREMENT,
		id_anime INT,
		numero INT,
		FOREIGN KEY (id_anime) REFERENCES animes(id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Criação da tabela Episodios
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS episodios (
		id INT PRIMARY KEY AUTO_INCREMENT,
		id_temporada INT,
		numero INT,
		urlthumb VARCHAR(255),
		intro TEXT,
		urlvideo VARCHAR(255),
		FOREIGN KEY (id_temporada) REFERENCES temporadas(id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

}
