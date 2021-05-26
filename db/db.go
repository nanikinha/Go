package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConectaComBancoDeDados() *sql.DB {

	err := godotenv.Load{}
	if err != nil {
		panic("Erro recuperando informações de ambiente")
	}

	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open(driver, user+":"+pass+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
