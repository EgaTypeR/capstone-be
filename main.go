package main

import (
	"context"
	"log"
	"os"

	"github.com/capstone-be/controllers"
	"github.com/capstone-be/routers"
	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Loading .env file")
	}

	router := gin.Default()
	client, err := utils.ConnectDB()
	if err != nil {
		log.Fatal("Error connect to database")
	}
	defer client.Disconnect(context.TODO())

	routers.InitRouters(router)

	go controllers.HandleBroadcast()

	log.Fatal(router.Run(":" + os.Getenv("PORT")))

}
