// Package service contains the service logic for inventory and user
package service

import (
	"github.com/adrr-dev/inventory-app/internal/repository"
	"gorm.io/gorm"
)

type InvenService struct {
	DB *gorm.DB
}

func (s InvenService) ListInventory(userID uint) ([]repository.Inventory, error) {
	var inventories []repository.Inventory
	result := s.DB.Where("user_id = ?", userID).Find(&inventories)
	if result.Error != nil {
		return nil, result.Error
	}

	return inventories, nil
}

func (s InvenService) FetchInventory(itemID, userID uint) (*repository.Inventory, error) {
	var inventory repository.Inventory

	result := s.DB.Where("id = ? AND user_id = ?", itemID, userID).First(&inventory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &inventory, nil
}

func (s InvenService) RemoveInventory(itemID, userID uint) error {
	result := s.DB.Where("id = ? AND user_id = ?", itemID, userID).Delete(&repository.Inventory{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s InvenService) EditItem(itemID, userID uint, item string) error {
	result := s.DB.Where("id = ? AND user_id = ?", itemID, userID).Model(&repository.Inventory{}).Update("item", item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s InvenService) EditLocation(itemID, userID uint, location string) error {
	result := s.DB.Where("id = ? AND user_id = ?", itemID, userID).Model(&repository.Inventory{}).Update("location", location)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s InvenService) ToggleStatus(itemID, userID uint) error {
	inventory, err := s.FetchInventory(itemID, userID)
	if err != nil {
		return err
	}

	status := inventory.Status
	result := s.DB.Where("id = ? AND user_id = ?", itemID, userID).Model(&repository.Inventory{}).Update("status", !status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s InvenService) CreateInventory(item, location string, userID uint) error {
	inventory := &repository.Inventory{Item: item, Location: location, Status: false, UserID: userID}
	result := s.DB.Create(inventory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
