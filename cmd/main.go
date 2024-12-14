package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/ztolley/goapi/cmd/api"
	"github.com/ztolley/goapi/internal/configs"
)

// @title User API
// @version 1.0
// @description This is a sample server for a user API.

// @contact.name Zac Tolley
// @contact.url https://exsite.dev
// @contact.email ztolley@exsiteconsultants.com

// @license.name MIT
// @BasePath /
func main() {

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		configs.Envs.DBUser,
		configs.Envs.DBPassword,
		configs.Envs.DBHost,
		configs.Envs.DBPort,
		configs.Envs.DBName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	server.Run()
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
