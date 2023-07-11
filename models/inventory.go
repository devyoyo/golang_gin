package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     Archive
	// Employees   []*Employe `gorm:"many2many:employees_inventories"`
	// Employe []EmployeeInventory `json="employees"`
}

type RequestInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type ResponseInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	Archive              ResponseArchive
}

// type EmployeeInventory struct {
// 	gorm.Model
// 	EmployeID   uint   `json:"employee_id"`
// 	InventoryID uint   `json:"inventory_id"`
// 	Description string `json:"description"`
// 	Employe     Employe
// 	Inventory   Inventory
// }
