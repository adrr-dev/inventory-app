package main

import (
	"log"
	"net/http"

	"github.com/adrr-dev/inventory-app/internal/database"
	"github.com/adrr-dev/inventory-app/internal/handlers"
	"github.com/adrr-dev/inventory-app/internal/service"
)

func main() {
	dataFile := "data.db"
	myDB := database.Database{DataFile: dataFile}
	db := myDB.InitializeDB()
	myInvenService := service.InvenService{DB: db}
	myUserService := service.UserService{DB: db}
	myHandling := handlers.Handling{UserService: myUserService, InvenService: myInvenService}

	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", myHandling.RootHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
