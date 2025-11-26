package attendance

import (
	"fmt"
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/utils"
)

type Handler struct {
	AttendanceRepo *Repository
}

func NewHandler(attenRepo *Repository) *Handler {
	return &Handler{
		AttendanceRepo: attenRepo,
	}
}

func (h *Handler) MarkAttendance(w http.ResponseWriter, r *http.Request) {
	var jobsInput []AttendanceJob
	if ok := utils.ParseJson(w, r, &jobsInput); !ok {
		return
	}
	if len(jobsInput) == 0 {
		utils.WriteError(w, http.StatusBadRequest, "no attendance data provided")
		return
	}

	// jobs channel
	jobs := make(chan AttendanceJob, len(jobsInput))
	results := StartAttendanceWorkerPool(10, jobs, h.AttendanceRepo)

	// notifications
	notifJobs := make(chan utils.NotificationJob, len(jobsInput))
	utils.StartNotificationWokerPool(20, notifJobs)

	// send jobs to workers
	for _, j := range jobsInput {
		jobs <- j
	}
	close(jobs)

	processed := 0
	failed := 0
	var errors []string
	var absentCount int

	for res := range results {
		processed++
		if res.Error != nil {
			failed++
			errors = append(errors, fmt.Sprintf("student %d: %v", res.Job.StudentID, res.Error))
			continue
		}

		// absent
		if !res.Job.Status {
			absentCount++
			msg := fmt.Sprintf("Your child %s is absent today.", res.Job.StudentName)
			notifJobs <- utils.NotificationJob{
				ToEmail:     res.Job.ParentEmail,
				StudentName: res.Job.StudentName,
				Message:     msg,
			}
		}
	}

	// close notification queue (workers will finish)
	close(notifJobs)

	utils.WriteJson(w, http.StatusOK, map[string]any{
		"processed": processed,
		"failed":    failed,
		"absent":    absentCount,
		"errors":    errors,
	})
}
