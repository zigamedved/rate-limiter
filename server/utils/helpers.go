package utils

import (
	"log"
	"net/http"
)

func GetClientId(r *http.Request) string {
	clientID := r.URL.Query().Get("clientId")
	if clientID == "" {
		log.Printf("No clientId query parameter found!")
		return ""
	}

	return clientID
}
