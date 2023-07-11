package routes

import (
	"golang_bsic_gin/auth"
	"golang_bsic_gin/config"
	"golang_bsic_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": "bad request",
			"error":  err.Error(),
		})

		c.Abort()
		return
	}

	err := user.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed Hash Password",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	insert := config.DB.Create(&user)

	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   insert.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

func GenerateToken(c *gin.Context) {
	reqToken := models.TokenRequest{}
	user := models.User{}

	if err := c.ShouldBindJSON(&reqToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": "bad request",
			"error":  err.Error(),
		})

		c.Abort()
		return
	}

	// check email
	checkEmail := config.DB.Where("email = ?", reqToken.Email).First(&user)
	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "email not found",
			"error":   checkEmail.Error.Error(),
		})

		c.Abort()
		return
	}

	// check password
	credentialError := user.CheckPassword(reqToken.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password not match",
			"error":   credentialError.Error(),
		})

		c.Abort()
		return
	}

	// generate token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	// response token

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
	})
}
