package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	cfg := config.NewConfig()

	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sql.Open("postgres", connUrl)

	if err != nil {
		log.Fatalln("Unable to open database connection:", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalln("Unable to initialize db driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalln("Error creating migration database instance", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalln("Unable to run migration:", err)
	}

	log.Println("Migration completed successfully.")

}
