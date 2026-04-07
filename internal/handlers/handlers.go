// Package handlers contains the handlers
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/adrr-dev/inventory-app/internal/repository"
)

type MyUserService interface {
	FetchUser(username, password string) (*repository.User, error)
	CreateUser(username, password string) error
	DeleteUser(id uint) error
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
	Tmpls        *template.Template
	Fragments    *template.Template
}

func (h Handling) renderError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)

	data := struct {
		Message string
		Code    int
	}{
		message,
		code,
	}

	err := h.Tmpls.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		log.Println("somehow an err happened")
	}
}

func (h Handling) LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpls.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) NewAccount(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpls.ExecuteTemplate(w, "create-user.html", nil)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := h.UserService.CreateUser(username, password)
	if err != nil {
		message := fmt.Sprintf("something went wrong when craeting account: %e", err)
		h.renderError(w, message, http.StatusBadRequest)
		return
	}

	log.Printf("new user craeted with username:%s, password:%s\n", username, password)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h Handling) DeleteUser(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")
	i, _ := strconv.Atoi(strID)
	userID := uint(i)

	err := h.UserService.DeleteUser(userID)
	if err != nil {
		message := fmt.Sprintf("user could not be deleted: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}

	log.Printf("user with user id: %s has been deleted", strID)

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (h Handling) InventoryHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	data, err := h.UserService.FetchUser(username, password)
	if err != nil {
		message := fmt.Sprintf("user could not be found. Create a new account instead: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}

	err = h.Tmpls.ExecuteTemplate(w, "inventory.html", data)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) NewItem(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("userid")
	i, _ := strconv.Atoi(strID)
	userID := uint(i)
	item := r.FormValue("item")
	location := r.FormValue("location")

	err := h.InvenService.CreateInventory(item, location, userID)
	if err != nil {
		message := fmt.Sprintf("trouble creating inventory: %e", err)
		h.renderError(w, message, http.StatusBadRequest)
		return
	}

	items, err := h.InvenService.ListInventory(userID)
	if err != nil {
		message := fmt.Sprintf("trouble creating inventory: %e", err)
		h.renderError(w, message, http.StatusBadRequest)
		return
	}
	data := struct {
		UserID uint
		Items  []repository.Inventory
	}{
		userID,
		items,
	}

	err = h.Fragments.ExecuteTemplate(w, "items.html", data)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) DeleteItem(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("userid")
	i, _ := strconv.Atoi(strID)
	userID := uint(i)
	strItemID := r.PathValue("id")
	i, _ = strconv.Atoi(strItemID)
	itemID := uint(i)

	err := h.InvenService.RemoveInventory(itemID, userID)
	if err != nil {
		message := fmt.Sprintf("could not delete item: %e", err)
		h.renderError(w, message, http.StatusBadRequest)
		return
	}
}

func (h Handling) EditInventory(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("userid")
	i, _ := strconv.Atoi(strID)
	userID := uint(i)
	strItemID := r.PathValue("id")
	i, _ = strconv.Atoi(strItemID)
	itemID := uint(i)

	data, err := h.InvenService.FetchInventory(itemID, userID)
	if err != nil {
		message := fmt.Sprintf("culd not fetch item data: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}

	err = h.Fragments.ExecuteTemplate(w, "edit.html", data)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) CancelEdit(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("userid")
	i, _ := strconv.Atoi(strID)
	userID := uint(i)
	strItemID := r.PathValue("id")
	i, _ = strconv.Atoi(strItemID)
	itemID := uint(i)

	data, err := h.InvenService.FetchInventory(itemID, userID)
	if err != nil {
		message := fmt.Sprintf("culd not fetch item data: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}

	err = h.Fragments.ExecuteTemplate(w, "item.html", data)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) EditItem(w http.ResponseWriter, r *http.Request) {
	log.Println("editing entire item")
	strID := r.PathValue("userid")
	i, _ := strconv.Atoi(strID)
	userID := uint(i)
	strItemID := r.PathValue("id")
	i, _ = strconv.Atoi(strItemID)
	itemID := uint(i)

	item := r.FormValue("item")
	location := r.FormValue("location")

	err := h.InvenService.EditItem(itemID, userID, item)
	if err != nil {
		message := fmt.Sprintf("could not edit item: %e", err)
		h.renderError(w, message, http.StatusBadRequest)
		return
	}

	err = h.InvenService.EditLocation(itemID, userID, location)
	if err != nil {
		message := fmt.Sprintf("could not edit location: %e", err)
		h.renderError(w, message, http.StatusBadRequest)
		return
	}

	data, err := h.InvenService.FetchInventory(itemID, userID)
	if err != nil {
		message := fmt.Sprintf("could not fetch data: %e", err)
		h.renderError(w, message, http.StatusNoContent)
		return

	}

	err = h.Fragments.ExecuteTemplate(w, "item.html", data)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		h.renderError(w, message, http.StatusNotFound)
		return
	}
}
