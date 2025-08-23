package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() error {
	godotenv.Load()
	connStr := os.Getenv("SUPABASE_CONNECTION_STRING")
	fmt.Println("ENV: ", connStr)

	var err error

	DB, err = sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal(err)
		return err
	}

	// pinging
	pingErr := DB.Ping()
	if err := pingErr; err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Connected with the DB...")
	return nil
}
