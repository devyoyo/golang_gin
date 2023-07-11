package routes

import (
	"golang_bsic_gin/config"
	"golang_bsic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetInventory(c *gin.Context) {

	inventories := []models.Inventory{}

	config.DB.Preload(clause.Associations).Find(&inventories)

	ResponseInventory := []models.ResponseInventory{}

	for _, inventory := range inventories {
		resInventory := models.ResponseInventory{
			InventoryName:        inventory.Name,
			InventoryDescription: inventory.Description,
			Archive: models.ResponseArchive{
				ArchiveName:        inventory.Archive.Name,
				ArchiveDescription: inventory.Archive.Description,
			},
		}

		ResponseInventory = append(ResponseInventory, resInventory)
	}

	// for _, p := range positions {
	// 	department := models.DepartmentResponse{}
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "success get position",
		"data":    ResponseInventory,
	})
}

func PostInventory(c *gin.Context) {
	var reqInventory models.RequestInventory

	c.BindJSON(&reqInventory)

	inventory := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInventory.ArchiveName,
			Description: reqInventory.ArchiveDescription,
		},
	}

	config.DB.Create(&inventory)

	c.JSON(http.StatusCreated, gin.H{
		"data":    reqInventory,
		"message": "success created position",
	})
}

func PutInventory(c *gin.Context) {
	id := c.Param("id")

	var inventory models.Inventory

	data := config.DB.First(&inventory, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	var reqInventory models.RequestInventory

	c.BindJSON(&reqInventory)

	inventory_update := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInventory.ArchiveName,
			Description: reqInventory.ArchiveDescription,
		},
	}

	log.Println(inventory)

	update := config.DB.Model(&inventory).Where("id = ?", id).Updates(&inventory_update)

	if update.Error != nil {
		log.Printf(update.Error.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": update.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    reqInventory,
	})

}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")

	var inventories models.Inventory

	data := config.DB.First(&inventories, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	config.DB.Delete(&inventories, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func GetInventoryById(c *gin.Context) {
	id := c.Param("id")

	var inventory models.Inventory

	data := config.DB.Preload("Archive").First(&inventory, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	ResponseInventory := models.ResponseInventory{
		InventoryName:        inventory.Name,
		InventoryDescription: inventory.Description,
		Archive: models.ResponseArchive{
			ArchiveName:        inventory.Archive.Name,
			ArchiveDescription: inventory.Archive.Description,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    ResponseInventory,
	})
}
