// Package handlers contains the handlers
package handlers

import (
	"net/http"

	"github.com/adrr-dev/inventory-app/internal/repository"
)

type MyUserService interface {
	FetchUser(username, password string) (*repository.User, error)
	CreateUser(username, password string) error
	DeleteUser(username, password string) error
}
type MyInvenService interface {
	ListInventory(userID uint) ([]repository.Inventory, error)
	FetchInventory(itemID, userID uint) (*repository.Inventory, error)
	RemoveInventory(itemID, userID uint) error
	EditItem(itemID, userID uint, item string) error
	EditLocation(itemID, userID uint, location string) error
	ToggleStatus(itemID, userID uint) error
	CreateInventory(item, location string, userID uint) error
}

type Handling struct {
	UserService  MyUserService
	InvenService MyInvenService
}

func (h Handling) RootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello there"))
}
