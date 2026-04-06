// Package routes contains the routes
package routes

import (
	"net/http"

	"github.com/adrr-dev/inventory-app/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handlers *handlers.Handling) {
	mux.Handle("GET /{$}", http.RedirectHandler("/login", http.StatusFound))
	mux.HandleFunc("GET /login", handlers.LoginHandler)
	mux.HandleFunc("GET /inventory", handlers.InventoryHandler)
	mux.HandleFunc("GET /new-account", handlers.NewAccount)
	mux.HandleFunc("POST /create-account", handlers.CreateUser)

	mux.HandleFunc("POST /new-item/{userid}", handlers.NewItem)
	mux.HandleFunc("DELETE /item/{userid}/{id}", handlers.DeleteItem)
	mux.HandleFunc("GET /item/edit/{userid}/{id}", handlers.EditInventory)
	mux.HandleFunc("GET /item/{userid}/{id}", handlers.CancelEdit)
	mux.HandleFunc("PUT /item/{userid}/{id}", handlers.EditItem)
}
