package config

import (
	"golang_bsic_gin/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/golang_gin?charset=utf8mb4&parseTime=True&loc=Local"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("error")
	}

	DB.AutoMigrate(&models.Department{}, &models.Position{}, &models.Employe{}, &models.Inventory{}, &models.Archive{}, &models.User{})

	// DB.Create(&models.Department{
	// 	Name: "Human Resources",
	// 	Code: "HR",
	// 	Positions: []models.Position{
	// 		{Name: "Manager HR", Code: "MHR"},
	// 		{Name: "Staff HR", Code: "SHR"},
	// 	},
	// })

	// DB.Create(&models.Employe{
	// 	Name:       "iksan",
	// 	Address:    "cipinang muara",
	// 	Email:      "iksan@iksan.com",
	// 	PositionID: 1,
	// })

	// DB.Create(&models.Employe{
	// 	Name:       "heru",
	// 	Address:    "purbalingga",
	// 	Email:      "heru@heru.com",
	// 	PositionID: 1,
	// })

	// DB.Create(&models.Employe{
	// 	Name:       "fadlan",
	// 	Address:    "jogja",
	// 	Email:      "fadlan@fadlan.com",
	// 	PositionID: 2,
	// })

	// DB.Create(&models.Employe{
	// 	Name:       "messi",
	// 	Address:    "argentina",
	// 	Email:      "messi@messi.com",
	// 	PositionID: 2,
	// })
}
