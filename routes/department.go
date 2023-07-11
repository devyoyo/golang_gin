package routes

import (
	"golang_bsic_gin/config"
	"golang_bsic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDepartment(c *gin.Context) {

	departments := []models.Department{}

	// siapin struct penampungnya
	GetDepartmentResponses := []models.GetDepartmentResponse{}

	// config.DB.Find(&departments)
	config.DB.Preload("Positions").Find(&departments)

	for _, d := range departments {
		positions := []models.PositionResponse{}

		for _, p := range d.Positions {
			pos := models.PositionResponse{
				ID:   p.ID,
				Name: p.Name,
				Code: p.Code,
			}

			positions = append(positions, pos)
		}

		dept := models.GetDepartmentResponse{
			ID:        d.ID,
			Name:      d.Name,
			Code:      d.Code,
			Positions: positions,
		}

		GetDepartmentResponses = append(GetDepartmentResponses, dept)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    GetDepartmentResponses,
	})
}

func PostDepartment(c *gin.Context) {
	// var department models.Department

	// cara port ver x-www-form
	// department := models.Department{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// }

	// cara post dengan json
	// ini hanya bisa bekerja jika di dalam struct sudah ada formatan penamaan json
	var department models.Department
	c.BindJSON(&department)

	config.DB.Create(&department)

	c.JSON(http.StatusCreated, gin.H{
		"data":    department,
		"message": "success created department",
	})
}

func PutDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	data := config.DB.First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	// ver Post Form
	// config.DB.Model(&department).Updates(models.Department{
	// 	Name: c.PostForm("name"),
	// // 	Code: c.PostForm("code"),
	// })

	// versi json

	c.BindJSON(&department)

	config.DB.Model(&department).Where("id = ?", id).Updates(&department)

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    department,
	})

}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	data := config.DB.First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	config.DB.Delete(&department, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	data := config.DB.Preload("Positions").First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		return
	}

	position := []models.PositionResponse{}
	for _, p := range department.Positions {
		pos := models.PositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		}

		position = append(position, pos)
	}

	GetDepartmentResponse := models.GetDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		Code:      department.Code,
		Positions: position,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    GetDepartmentResponse,
	})
}
