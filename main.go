package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nacharya01/GOWeb/controller"
	_ "github.com/nacharya01/GOWeb/db"
	"github.com/nacharya01/GOWeb/logger"
	"github.com/nacharya01/GOWeb/router"
)

var LOG *logger.Logger = logger.LOG;

func main(){	
	//Load .env file
	errEnv := godotenv.Load(".env")
	if errEnv != nil{
		LOG.Error("Error loading .env file")
	}
	LOG.Info("Successfully loaded .env file.")

	//GET Router
	router := router.GetRouter()
	router.HandleFunc("/", controller.HandleHomePage).Methods("GET")

	// Serve static files from the "static" directory
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static").Handler(s)

	//Start server 
	server := getServer(router)
	
	LOG.Info("Server successfully started on port : 8000")
	err := server.ListenAndServe()

	if err != nil{
		LOG.Error("An unexpected error ocurred while starting the server" + err.Error() )
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
