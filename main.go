package main

import (
	"golang_bsic_gin/config"
	"golang_bsic_gin/middleware"
	"golang_bsic_gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("GOLANG BASIC GIN")

	config.InitDB()
	router := gin.Default()
	// router.GET("/ping", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	router.GET("/", getHome)

	v1 := router.Group("api/v1")
	{
		department := v1.Group("/departments").Use(middleware.Auth())
		{
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentById)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)
		}

		position := v1.Group("/positions").Use(middleware.Auth())
		{
			position.GET("/", routes.GetPosition)
			position.GET("/:id", routes.GetPositionById)
			position.POST("/", routes.PostPosition)
			position.PUT("/:id", routes.PutPosition)
			position.DELETE("/:id", routes.DeletePosition)
		}

		employes := v1.Group(("employes")).Use(middleware.Auth())
		{
			employes.GET("/", routes.GetEmployes)
			employes.GET("/:id", routes.GetEmployesById)
			employes.POST("/", routes.PostEmployes)
			employes.PUT("/:id", routes.PutEmployes)
			employes.DELETE("/:id", routes.DeleteEmployes)
		}

		inventory := v1.Group("/inventories").Use(middleware.Auth())
		{
			inventory.GET("/", routes.GetInventory)
			inventory.POST("/", routes.PostInventory)
			inventory.PUT("/:id", routes.PutInventory)
			inventory.GET("/:id", routes.GetInventoryById)
			inventory.DELETE("/:id", routes.DeleteInventory)
		}

		peminjaman := v1.Group("/peminjaman")
		{
			peminjaman.GET("/", routes.GetPeminjaman)
			peminjaman.GET("/employe/:id", routes.GetPeminjamanByEmploye)
			peminjaman.GET("/inventory/:id", routes.GetPeminjamanByInventory)
			peminjaman.POST("/employe", routes.PostPeminjamanByEmploye)
		}

		User := v1.Group("/user")
		{
			User.POST("/register", routes.RegisterUser)
			User.POST("/generate_token", routes.GenerateToken)
		}
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "home",
	})
}
