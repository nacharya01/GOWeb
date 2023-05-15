package controller

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/nacharya01/GOWeb/logger"
	"github.com/nacharya01/GOWeb/model"
)

var LOG = logger.New()
func HandleHomePage(w http.ResponseWriter, r *http.Request){

	// Making sure that the type of request is GET
	if r.Method == "GET"{
		LOG.Info().Println("One request came from: ", r.RemoteAddr)
		homePage, err := template.ParseFiles("./templates/news.html")

		if err != nil{
			LOG.Fatal().Println("An error ocurred while loading the template for home page.")
			panic(err)
		}else{

			// Now initializing the model
			model := &model.News{
				AppName: os.Getenv("go.application.name"),
				Headline: "Welcome the website made by Golang",
				Body: "This is for practice.",
			}

			err := homePage.Execute(w, model);

			if err != nil{
				log.Fatal("An error ocurred while executing the template for home page.")
				return
			}
		}

	}
}