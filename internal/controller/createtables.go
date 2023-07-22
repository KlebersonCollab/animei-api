package controller

import (
	"database/sql"
	"log"
)

func CreateTables() {
	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", "root:Thor0528@tcp(192.168.15.7:3306)/animes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criação da tabela Anime
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS anime (
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

	// Criação da tabela Episodios
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS episodios (
		id INT PRIMARY KEY AUTO_INCREMENT,
		id_anime INT,
		temporada INT,
		episodio INT,
		urlthumb VARCHAR(255),
		intro TEXT,
		urlvideo VARCHAR(255),
		FOREIGN KEY (id_anime) REFERENCES anime(id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabelas criadas com sucesso!")
}
