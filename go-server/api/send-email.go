package handler

import (
	"encoding/json"
	"net/http"

	"github.com/onosejoor/email-api/mailer"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string `json:"html,omitempty"`
	From    string `json:"from"`
}

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload JsonResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, JsonResponse{false, "Method not allowed"})
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, JsonResponse{false, "Invalid JSON body"})
		return
	}

	if req.To == "" || req.Subject == "" || req.HTML == "" {
		writeJSON(w, http.StatusBadRequest, JsonResponse{false, "Missing required fields"})
		return
	}

	if err := mailer.SendEmail(req.From, req.To, req.Subject, req.HTML); err != nil {
		writeJSON(w, http.StatusInternalServerError, JsonResponse{false, "Failed to send email: " + err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, JsonResponse{true, "Email sent successfully"})
}
