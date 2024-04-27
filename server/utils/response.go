package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(w http.ResponseWriter, payload interface{}) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}
	_, err2 := w.Write(jsonData)
	if err2 != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err2), http.StatusInternalServerError)
		return
	}
}

type ErrorResponse struct {
	Error      string   `json:"error"`
	StatusCode int      `json:"statusCode"`
	Messages   []string `json:"messages"`
}

func SendErrorJson(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)

	errResponse := ErrorResponse{
		Error:      http.StatusText(statusCode),
		StatusCode: statusCode,
		Messages:   []string{},
	}

	errResponse.Messages = append(errResponse.Messages, err.Error())

	SendJson(w, errResponse)
}
