package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controllers "xenditapps/controllers"
)

func main() {
	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/xendit-get-invoice-list", controllers.GetInvoices)
	server.POST("/xendit-create-invoice", controllers.CreateInvoice)
	server.POST("/xendit-webhook-invoice-created", controllers.InvoiceCreated)

	server.Run(":8080")
}
