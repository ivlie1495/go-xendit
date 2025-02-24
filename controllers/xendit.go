package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xendit/xendit-go/v6/invoice"

	models "xenditapps/models"
	utils "xenditapps/utils"
)

// RandomInt generates a random integer between min and max (inclusive)
func RandomInt(min, max int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func GetInvoices(c *gin.Context) {
	invoices := models.GetInvoices()
	c.JSON(200, gin.H{
		"data": invoices,
	})
}

func StrPtr(s string) *string {
	return &s
}

func CreateInvoice(c *gin.Context) {
	invoiceBody := models.InvoiceBody{}
	errBody := c.ShouldBindJSON(&invoiceBody)

	if errBody != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body",
		})
		return
	}

	externalId := RandomInt(1, 10000)
	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(externalId, float64(invoiceBody.Amount)) // [REQUIRED] | CreateInvoiceRequest
	createInvoiceRequest.PaymentMethods = []string{invoiceBody.PaymentMethod}
	createInvoiceRequest.SuccessRedirectUrl = StrPtr(utils.Env["XENDIT_SUCCESS_REDIRECT_URL"])
	createInvoiceRequest.FailureRedirectUrl = StrPtr(utils.Env["XENDIT_FAILURE_REDIRECT_URL"])

	resp, r, err := utils.XenditClient.InvoiceApi.CreateInvoice(c).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice``: %v\n", err.Error())

		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create invoice",
		})
		return
	}

	// response from `CreateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoice`: %v\n", resp)

	c.JSON(http.StatusCreated, resp)
}
