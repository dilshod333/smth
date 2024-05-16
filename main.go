package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

func main() {
 
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

   
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    
    connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)


    db, err := sql.Open("postgres", connString)
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer db.Close()

   
    if err := db.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    fmt.Println("Connected to PostgreSQL database successfully")
}
