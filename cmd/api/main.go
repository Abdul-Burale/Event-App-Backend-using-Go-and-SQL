package main

import (
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
	_ "github.com/joho/godotenv/autoload"
	"github.com/Abdul-Burale/Event-App-Backend-using-Go-and-SQL/internal/database"
	"github.com/Abdul-Burale/Event-App-Backend-using-Go-and-SQL/internal/env"
	_"github.com/Abdul-Burale/Event-App-Backend-using-Go-and-SQL/docs"
	
)

// @title Go Gin Rest API
// @version 1.0
// @description A rest API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

type application struct {
	port int
	jwtSecret string
	models database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port: env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "123456"),
		models: models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
