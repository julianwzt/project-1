package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func ConnectDB() {
	host := "db"
	port := 5432
	user := "postgres"
	password := "20060721"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}
	fmt.Println("Berhasil terhubung ke database PostgreSQL!")
}
