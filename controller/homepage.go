package controller

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/nacharya01/GOWeb/logger"
	"github.com/nacharya01/GOWeb/model"
)

var LOG *logger.Logger = logger.LOG
func HandleHomePage(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
		LOG.Info("One request came from: " + r.RemoteAddr)
		homePage, err := template.ParseFiles("./templates/news.html")

		if err != nil{
			LOG.Fatal("An error ocurred while loading the template for home page.")
			panic(err)
		}else{

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