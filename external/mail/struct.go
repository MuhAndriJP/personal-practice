package mail

import (
	"os"
	"strconv"
)

var (
	CONFIG_SMTP_HOST    = os.Getenv("MAIL_SMTP_HOST")
	CONFIG_SMTP_PORT, _ = strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))
	CONFIG_USERNAME     = os.Getenv("MAIL_USERNAME")
	CONFIG_EMAIL        = os.Getenv("MAIL_SENDER")
	CONFIG_PASSWORD     = os.Getenv("MAIL_PASSWORD")

	FROM        = "From"
	TO          = "To"
	SUBJECT     = "Subject"
	BODY_FORMAT = "text/html"

	SENDER_NAME   = "rockxhast 1 <" + os.Getenv("MAIL_SENDER") + ">"
	SUBJECT_EMAIL = "Subject Email"
)

type MailRequest struct {
	Recipient string `json:"recipient"`
	Body      string `json:"body"`
	CC        string `json:"cc"`
}
