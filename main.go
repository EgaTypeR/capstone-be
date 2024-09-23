package main

import (
	"context"
	"log"
	"os"

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
	client := utils.ConnectDB()
	defer client.Disconnect(context.TODO())
	

	db := client.Database("CrimeAlertCapstone")
	routers.InitRouters(db, router)

	log.Fatal(router.Run(":" + os.Getenv("PORT")))

}
