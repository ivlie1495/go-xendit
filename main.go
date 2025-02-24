package main

import (
	"github.com/gin-gonic/gin"

	controllers "xenditapps/controllers"
)

func main() {
	server := gin.Default()

	server.GET("/xendit-get-invoice-list", controllers.GetInvoices)
	server.POST("/xendit-create-invoice", controllers.CreateInvoice)

	server.Run(":8080")
}
