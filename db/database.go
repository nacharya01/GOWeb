package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/nacharya01/GOWeb/logger"
)

type config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DB       *sql.DB
}

var DBConfig = &config{}
var LOG *logger.Logger = logger.LOG;

func init() {
    
	*DBConfig = config{
		Host:     os.Getenv("go.database.host"),
		Port:     os.Getenv("go.database.port"),
		User:     os.Getenv("go.database.user"),
		Password: os.Getenv("go.database.password"),
		DBName:   os.Getenv("go.database.database"),
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBConfig.Host, DBConfig.Port, DBConfig.User, DBConfig.Password, DBConfig.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
        LOG.Info(err.Error())
		panic(err)
	}

	DBConfig.DB = db
    LOG.Info("Database initialization has been successfull")
}

func GetDBConnection()(*sql.DB){
    return DBConfig.DB
}
