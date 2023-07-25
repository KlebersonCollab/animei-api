package main

import (
	"klebersonromero/github.com/animeapi/cmd/helper"
	"klebersonromero/github.com/animeapi/router"
	"log"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Print("Sem arquivo .env encontrado")
	}
	// Cria as tabelas
	helper.CreateTables()
	// Inicializar o roteador
	router.Initialize()
}
