package models

import "gorm.io/gorm"

type Employe struct {
	gorm.Model
	Name       string   `json:"name"`
	Address    string   `json:"address"`
	Email      string   `json:"email"`
	PositionID uint     `json:"position_id"`
	Position   Position `json:"department"`
	// Inventories []*Inventory `gorm:"many2many:employees_inventories"`
	// Inventory []EmployeeInventory `json:"inventories"`
}
