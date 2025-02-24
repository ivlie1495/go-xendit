package utils

import (
	"github.com/xendit/xendit-go/v6"
)

// InitXenditClient initializes and returns a new Xendit client
func InitXenditClient() *xendit.APIClient {
	xenditApiKey := Env["XENDIT_API_KEY"]

	return xendit.NewClient(xenditApiKey)
}

// XenditClient is a globally accessible Xendit client
var XenditClient = InitXenditClient()
