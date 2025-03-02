package config

import(
	"github.com/nethsaraPrabash/chat-app-go/src/models"

	"fmt"
    "log"
    "os"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"

)

var DB *gorm.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    databaseName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB = db
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failedto migrate database:", err)
	}

	fmt.Println("✅ Database connected & migrated successfully!")

}
