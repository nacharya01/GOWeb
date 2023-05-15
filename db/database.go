package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

var DB *sql.DB

func Init() {
    cfg := Config{
        Host:     os.Getenv("go.database.host"),
        Port:     os.Getenv("go.database.port"),
        User:     os.Getenv("go.database.user"),
        Password: os.Getenv("go.database.password"),
        DBName:   os.Getenv("go.database.database"),
    }
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    DB = db
}
