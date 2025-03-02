package routes

import (
	"github.com/nethsaraPrabash/chat-app-go/src/controller"
	"github.com/nethsaraPrabash/chat-app-go/src/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controller.RegisterUser)
		userRoutes.POST("/login", controller.LoginUser)
	}

	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/protected", controller.ProtectedEndpoint)
	}
}