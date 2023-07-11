package routes

import (
	"golang_bsic_gin/config"
	"golang_bsic_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployes(c *gin.Context) {
	employe := []models.Employe{}

	config.DB.Preload("Position").Find(&employe)

	c.JSON(http.StatusOK, gin.H{
		"data": employe,
	})
}

func GetEmployesById(c *gin.Context) {

}

func PostEmployes(c *gin.Context) {

}

func PutEmployes(c *gin.Context) {

}

func DeleteEmployes(c *gin.Context) {

}
