package middleware

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var (
	limiter *rate.Limiter
	clients map[string]*rate.Limiter
	mu      sync.Mutex
)

func init() {
	limiter = rate.NewLimiter(rate.Limit(100), 100)
	clients = make(map[string]*rate.Limiter)
}

func getClientLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, ok := clients[ip]
	if !ok {
		limiter = rate.NewLimiter(rate.Limit(10), 10)
		clients[ip] = limiter
	}
	return limiter
}

func RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// global
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// per client
		ip := r.RemoteAddr
		clientLimiter := getClientLimiter(ip)
		if !clientLimiter.Allow() {
			http.Error(w, fmt.Sprintf("Rate limit exceeded for %s", ip), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	}
}
