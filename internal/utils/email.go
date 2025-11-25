package utils

import (
	"fmt"
	"time"
)

func SendWelcomeEmail(toEmail, username, password string) {
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[EMAIL] To: %s | Username: %s | Password: %s\n", toEmail, username, password)
		// TODO: integrate SMTP, handle retries, templating, HTML, etc.
	}()
}
