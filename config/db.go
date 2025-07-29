package config

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() *sql.DB {
	connString := "server=localhost,1433;database=GoCrudDB;trusted_connection=yes;encrypt=disable"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	return db
}
