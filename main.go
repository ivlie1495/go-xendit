package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	controllers "xenditapps/controllers"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := gin.Default()

	server.GET("/xendit-get-invoice-list", controllers.GetInvoices)
	server.POST("/xendit-create-invoice", controllers.CreateInvoice)

	server.Run(":8080")
}
