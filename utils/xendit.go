package utils

import (
	"os"

	"github.com/xendit/xendit-go/v6"
)

// InitXenditClient initializes and returns a new Xendit client
func InitXenditClient() *xendit.APIClient {
	xenditApiKey := os.Getenv("XENDIT_API_KEY")

	return xendit.NewClient(xenditApiKey)
}

// XenditClient is a globally accessible Xendit client
var XenditClient = InitXenditClient()
