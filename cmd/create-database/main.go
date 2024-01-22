package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/config"
)

func main() {
	cfg := config.NewConfig()

	connUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
	)

	db, err := sql.Open("postgres", connUrl)

	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	var exists bool

	q, err := db.Query("SELECT 1 FROM pg_database WHERE datname = '?'", cfg.Database.Name)

	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	q.Scan(&exists)

	if exists {
		log.Println("Database already exists.")
		return
	}

	_, err = db.Exec("CREATE DATABASE ?", cfg.Database.Name)

	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	log.Println("Database created.")
}
