package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/capstone-be/controllers"
	"github.com/capstone-be/routers"
	"github.com/capstone-be/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Loading .env file")
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,                                              // Allow your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	client, err := utils.ConnectDB()
	if err != nil {
		log.Fatal("Error connect to database")
	}
	defer client.Disconnect(context.TODO())

	routers.InitRouters(router)

	go controllers.HandleBroadcast()

	log.Fatal(router.Run(":" + os.Getenv("PORT")))

}
