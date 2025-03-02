package controller 


import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"

	"github.com/nethsaraPrabash/chat-app-go/src/models"
	"github.com/nethsaraPrabash/chat-app-go/src/service"
	"github.com/nethsaraPrabash/chat-app-go/src/helpers"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create user"})
		return
	}

	err := service.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := helpers.GenerateJWT(user.ID, secretKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not genereate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": token, "user": user})
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
    token, err := helpers.GenerateJWT(user.ID, secretKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have accessed a protected endpoint"})
}