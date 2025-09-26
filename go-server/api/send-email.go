package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	gomail "gopkg.in/mail.v2"
)

var (
	dialer *gomail.Dialer
	sender gomail.SendCloser
	once   sync.Once
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

	if err := SendEmail(req.From, req.To, req.Subject, req.HTML); err != nil {
		writeJSON(w, http.StatusInternalServerError, JsonResponse{false, "Failed to send email: " + err.Error()})
		return
	}

	log.Println("Email Sent successfully to " + req.To + " from " + req.From)
	writeJSON(w, http.StatusOK, JsonResponse{true, "Email sent successfully"})
}

func initDialer() {
	dialer = gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("GMAIL_USER"), os.Getenv("GMAIL_APP_PASSWORD"))
	var err error
	sender, err = dialer.Dial()
	if err != nil {
		log.Fatalf("Failed to dial SMTP: %v", err)
	}
}

func SendEmail(from, to, subject, body string) error {
	once.Do(initDialer)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", from, os.Getenv("GMAIL_USER")))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return gomail.Send(sender, m)
}
