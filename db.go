package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func iniciaDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)

	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("Banco nao inicializado")
	}

	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	sqlTable := `
		Create table if not exists tarefas(
			id Integer not null primary key autoincrement,
			titulo text not null
		)
	`

	_, err := db.Exec(sqlTable)

	if err != nil {
		log.Fatalf("Falha ao criar tabelas: 	%v", err)
	}
}
