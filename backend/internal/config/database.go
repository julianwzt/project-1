package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    host := getenv("DB_HOST", "postgres-service")
    user := getenv("DB_USER", "postgres")
    pass := getenv("DB_PASSWORD", "password")
    name := getenv("DB_NAME", "mahasiswa_db")
    port := getenv("DB_PORT", "5432")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
        host, user, pass, name, port)
    
    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Error opening database: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    
    fmt.Println("Berhasil terhubung ke database!")
}

func getenv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}