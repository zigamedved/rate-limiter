package handler

import (
	"net/http"

	u "github.com/zigamedved/rate-limiter/server/utils"
)

type response struct {
	Message string `json:"message"`
}

func Reply(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	u.SendJson(w, response{Message: "Successful response"})
}
