package utils

import "time"

type NotificationJob struct {
	ToEmail     string
	StudentName string
	SchoolName  string
	Message     string
}

func StartNotificationWokerPool(
	n int,
	jobs <-chan NotificationJob,
) {
	for i := 0; i < n; i++ {
		go func(wID int) {
			for job := range jobs {
				if job.ToEmail != "" {
					subject := "Attendance update"
					body := job.Message
					_ = SendAbsenceEmail(job.ToEmail, subject, body)
				}
				// small sleep  to avoid hammering providers
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
}
