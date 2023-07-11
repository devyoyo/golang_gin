package routes

import (
	"golang_bsic_gin/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponsePeminjaman struct {
	EmployeeID    uint      `json:"employee_id"`
	InventoryID   uint      `json:"inventory_id"`
	EmployeeName  string    `json:"employee_name"`
	InventoryName string    `json:"inventory_name"`
	Description   string    `json:"description"`
	RentalDate    time.Time `json:"rental_date"`
}

type RequestPeminjaman struct {
	EmployeeID  uint   `json:"employee_id"`
	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}

func GetPeminjaman(c *gin.Context) {
	// EmployeeInventoies := []models.EmployeeInventory{}

	// config.DB.Preload(clause.Associations).Find(&EmployeeInventoies)

	// resPeminjaman := []ResponsePeminjaman{}

	// for _, eInv := range EmployeeInventoies {
	// 	resPem := ResponsePeminjaman{
	// 		EmployeeID:    eInv.EmployeID,
	// 		InventoryID:   eInv.InventoryID,
	// 		EmployeeName:  eInv.Employe.Name,
	// 		InventoryName: eInv.Inventory.Name,
	// 		RentalDate:    eInv.CreatedAt,
	// 	}

	// 	resPeminjaman = append(resPeminjaman, resPem)
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"data":    resPeminjaman,
	// 	"message": "welcome to data peminjaman",
	// })
}

func GetPeminjamanByEmploye(c *gin.Context) {
	// id := c.Param("id")

	// var employeesInventories []models.EmployeeInventory

	// 	data := config.DB.Preload(clause.Associations).First(&employeesInventories, "employe_id = ?", id)

	// 	if data.Error != nil {
	// 		log.Printf(data.Error.Error())
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"status":  http.StatusNotFound,
	// 			"message": "Data not found",
	// 		})

	// 		return
	// 	}

	// ResponsePeminjaman:

	// 	for _, eInv := range employeesInventories {
	// 		respPeminjaman := ResponsePeminjaman{}

	// 		respPeminjaman = append()
	// 	}
	// 	ResponseInventory := models.ResponsePe{
	// 		InventoryName:        inventory.Name,
	// 		InventoryDescription: inventory.Description,
	// 		Archive: models.ResponseArchive{
	// 			ArchiveName:        inventory.Archive.Name,
	// 			ArchiveDescription: inventory.Archive.Description,
	// 		},
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "update success",
	// 		"data":    ResponseInventory,
	// 	})
}

func GetPeminjamanByInventory(c *gin.Context) {

}

func PostPeminjamanByEmploye(c *gin.Context) {
	var reqPeminjaman RequestPeminjaman

	if err := c.ShouldBindJSON(&reqPeminjaman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"status":  err.Error(),
		})

		c.Abort()
		return
	}

	insert := config.DB.Create(&reqPeminjaman)

	if insert.Error != nil {
		c.JSON(http.StatusCreated, gin.H{
			"data":    reqPeminjaman,
			"message": "insert success",
		})
	}
}
