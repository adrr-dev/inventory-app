package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/adrr-dev/inventory-app/internal/database"
	"github.com/adrr-dev/inventory-app/internal/handlers"
	"github.com/adrr-dev/inventory-app/internal/routes"
	"github.com/adrr-dev/inventory-app/internal/service"
)

func main() {
	dataFile := "data.db"
	myDB := &database.Database{DataFile: dataFile}
	db := myDB.InitializeDB()

	myInvenService := &service.InvenService{DB: db}
	myUserService := &service.UserService{DB: db}

	tmpls, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fragments, err := template.ParseGlob("templates/fragments/*.html")
	if err != nil {
		log.Fatal(err)
	}
	myHandling := &handlers.Handling{UserService: myUserService, InvenService: myInvenService, Tmpls: tmpls, Fragments: fragments}

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, myHandling)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
