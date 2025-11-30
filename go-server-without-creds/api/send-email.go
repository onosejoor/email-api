package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/onosejoor/email-api/go-server-without-creds/pkg/mailer"
	"github.com/onosejoor/email-api/go-server-without-creds/types"
	"golang.org/x/time/rate"
)

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload JsonResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

var (
	limiters = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getLimiter(key string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limiters[key]
	if !exists {
		limiter = rate.NewLimiter(1, 5)
		limiters[key] = limiter
	}
	return limiter
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, JsonResponse{false, "Method not allowed"})
		return
	}

	key := r.RemoteAddr
	if !getLimiter(key).Allow() {
		writeJSON(w, http.StatusTooManyRequests, JsonResponse{false, "Too many requests"})
		return
	}

	token := r.Header.Get("X-API-KEY")
	if token == "" || token != os.Getenv("EMAIL_API_TOKEN") {
		writeJSON(w, http.StatusUnauthorized, JsonResponse{false, "Unauthorized. Invalid Header Token"})
		return
	}

	var req types.SendEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, JsonResponse{false, "Invalid JSON body"})
		return
	}

	if len(req.To) == 0 || req.Subject == "" || req.HTML == "" {
		writeJSON(w, http.StatusBadRequest, JsonResponse{false, "Missing required fields"})
		return
	}

	if err := mailer.SendEmail(req); err != nil {
		writeJSON(w, http.StatusInternalServerError, JsonResponse{false, "Failed to send email: " + err.Error()})
		return
	}

	log.Println("Email Sent successfully to " + strings.Join(req.To, ",") + " from " + req.From)
	writeJSON(w, http.StatusOK, JsonResponse{true, "Email sent successfully"})
}
