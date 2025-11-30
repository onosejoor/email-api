package types

type SendEmailRequest struct {
	To                 []string `json:"to"`
	Subject            string   `json:"subject"`
	HTML               string   `json:"html,omitempty"`
	From               string   `json:"from"`
	GMAIL_USER         string   `json:"gmail_user"`
	GMAIL_APP_PASSWORD string   `json:"gmail_app_password"`
}
