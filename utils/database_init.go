package utils

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func DbConnection() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	userName := os.Getenv("DBUserName")
	dbPasswrod := os.Getenv("DBPassword")
	dbName := os.Getenv("DBName")
	connectionQuery := fmt.Sprintf("host=localhost port=5437 user=%s password=%s dbname=%s sslmode=disable", userName, dbPasswrod, dbName)
	db, err := sqlx.Connect("pgx", connectionQuery)
	if err != nil {
		return nil, err
	}
	return db, nil
}
