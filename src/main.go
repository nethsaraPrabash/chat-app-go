package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nethsaraPrabash/chat-app-go/src/config"
	"github.com/nethsaraPrabash/chat-app-go/src/routes"
)

func main() {
	loadEnv()
	router := gin.Default()

	config.ConnectDB()

	routes.Routes(router)


	router.Run(":8080")
	fmt.Println("The server is running on port 8080")
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Println("env not found")

		err = godotenv.Load()
		if err != nil {
			log.Fatal("error loading env")
		}
	}
}