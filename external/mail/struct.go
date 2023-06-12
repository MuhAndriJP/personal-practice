package mail

import (
	"os"
	"strconv"
)

var (
	CONFIG_SMTP_HOST    = os.Getenv("MAIL_SMTP_HOST")
	CONFIG_SMTP_PORT, _ = strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))
	CONFIG_EMAIL        = os.Getenv("MAIL_SENDER")
	CONFIG_PASSWORD     = os.Getenv("MAIL_PASSWORD")

	FROM        = "From"
	TO          = "To"
	SUBJECT     = "Subject"
	BODY_FORMAT = "text/html"
)

type MailRequest struct {
	SenderName string `json:"sender_name"`
	To         string `json:"to"`
	CC         string `json:"cc"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Attach     string `json:"attach"`
}
