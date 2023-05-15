package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nacharya01/GOWeb/controller"
	"github.com/nacharya01/GOWeb/db"
	"github.com/nacharya01/GOWeb/logger"
	"github.com/nacharya01/GOWeb/router"
)

var LOG *logger.LogDir = logger.New();

func main(){	
	//Load .env file
	errEnv := godotenv.Load(".env")
	if errEnv != nil{
		LOG.Error().Println("Error loading .env file", errEnv)
	}
	LOG.Info().Println("Successfully loaded .env file.")

	//Init Database
	db.Init()
	
	//GET Router
	router := router.GetRouter()
	router.HandleFunc("/", controller.HandleHomePage).Methods("GET")

	// Serve static files from the "static" directory
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static").Handler(s)

	//Start server 
	server := getServer(router)
	
	LOG.Info().Printf("Server successfully started on port : %v", 8000)
	err := server.ListenAndServe()

	if err != nil{
		LOG.Error().Println("An unexpected error ocurred while starting the server", err)
	}
}

func getServer(r *mux.Router) *http.Server{
	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	return srv
}
