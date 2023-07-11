package routes

import (
	"golang_bsic_gin/config"
	"golang_bsic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosition(c *gin.Context) {

	positions := []models.Position{}

	config.DB.Preload("Department").Find(&positions)

	GetPositionResponse := []models.GetPositionResponse{}

	for _, p := range positions {
		pos := models.GetPositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
			Department: models.DepartmentResponse{
				Name: p.Department.Name,
				Code: p.Department.Code,
			},
		}

		GetPositionResponse = append(GetPositionResponse, pos)
	}

	// for _, p := range positions {
	// 	department := models.DepartmentResponse{}
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "success get position",
		"data":    GetPositionResponse,
	})
}

func PostPosition(c *gin.Context) {
	// var position models.Position

	// cara port ver x-www-form
	// position := models.Position{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// }

	// cara post dengan json
	// ini hanya bisa bekerja jika di dalam struct sudah ada formatan penamaan json
	var position models.Position
	c.BindJSON(&position)

	data := config.DB.Create(&position)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": data.Error.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    position,
		"message": "success created position",
	})
}

func PutPosition(c *gin.Context) {
	id := c.Param("id")

	var position models.Position

	data := config.DB.First(&position, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	// // ver Post Form
	// config.DB.Model(&position).Updates(models.Position{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// })

	// versi json

	c.BindJSON(&position)

	update := config.DB.Model(&position).Where("id = ?", id).Updates(&position)

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
		"data":    position,
	})

}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")

	var position models.Position

	data := config.DB.First(&position, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	config.DB.Delete(&position, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func GetPositionById(c *gin.Context) {
	id := c.Param("id")

	var position models.Position

	data := config.DB.Preload("Department").First(&position, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	GetPositionResponse := models.GetPositionResponse{
		ID:   position.ID,
		Name: position.Name,
		Code: position.Code,
		Department: models.DepartmentResponse{
			Name: position.Department.Code,
			Code: position.Department.Name,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get position",
		"data":    GetPositionResponse,
	})
}
