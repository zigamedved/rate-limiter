package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	h "github.com/zigamedved/rate-limiter/server/handler"
	ratelimiter "github.com/zigamedved/rate-limiter/server/rateLimiter"
)

const (
	port              = 8080
	maxRequests       = 5
	timeFrameDuration = 5
)

func main() {

	rl := ratelimiter.NewRateLimiter(maxRequests, timeFrameDuration*time.Second)

	mux := http.NewServeMux()

	mux.Handle("/", rl.HandleRequest(http.HandlerFunc(h.Reply)))

	log.Printf("Listening on :%v...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
