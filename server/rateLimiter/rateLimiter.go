package ratelimiter

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	u "github.com/zigamedved/rate-limiter/server/utils"
)

type rateLimiter struct {
	mu                 sync.Mutex
	clientFirstRequest map[string]time.Time
	requestCount       map[string]int
	maxRequests        int
	timeFrame          time.Duration
}

func (rl *rateLimiter) HandleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		clientID := u.GetClientId(r)
		if clientID == "" {
			log.Printf("Couldn't get clientId from request!")
			u.SendErrorJson(w, fmt.Errorf("missing clientId query parameter"), http.StatusBadRequest)
			return
		}

		rl.mu.Lock()
		defer rl.mu.Unlock()

		currentReqTime := time.Now()
		clientFirstReq := rl.clientFirstRequest[clientID]
		fmt.Printf("client first request %v", clientFirstReq)
		// if client performed at least 1 request && his first request was more than 5 seconds ago, we delete his entries
		if !clientFirstReq.IsZero() && currentReqTime.Sub(clientFirstReq) > rl.timeFrame {
			rl.clientFirstRequest[clientID] = time.Time{}
			rl.requestCount[clientID] = 0
		}
		if clientFirstReq.IsZero() {
			rl.clientFirstRequest[clientID] = currentReqTime
		}
		rl.requestCount[clientID] = rl.requestCount[clientID] + 1

		if rl.requestCount[clientID] > rl.maxRequests {
			log.Printf("Too many requests from clientId=%v", clientID)

			u.SendErrorJson(w, fmt.Errorf("too many requests from client"), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func NewRateLimiter(maxRequests int, timeFrame time.Duration) *rateLimiter {
	return &rateLimiter{
		clientFirstRequest: make(map[string]time.Time),
		requestCount:       make(map[string]int),
		maxRequests:        maxRequests,
		timeFrame:          timeFrame,
	}
}
