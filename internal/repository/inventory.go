// Package repository contains the model
package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Inventory []Inventory
}

type Inventory struct {
	gorm.Model
	Item     string
	Location string
	Status   bool // for now status is not used
	UserID   uint
}
