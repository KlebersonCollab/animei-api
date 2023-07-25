package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Função para abrir a conexão com o banco de dados
func OpenDatabaseConnection() (*sql.DB, error) {
	databasetype := "mysql"
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	ipaddress := os.Getenv("IPSERVER")
	portserver := os.Getenv("PORTSERVER")
	database := os.Getenv("DATABASE")
	linkserver := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, ipaddress, portserver, database)
	db, err := sql.Open(databasetype, linkserver)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
