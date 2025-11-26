package attendance

import (
	"fmt"
	"sync"
)

type AttendanceJob struct {
	StudentID   uint   `json:"student_id"`
	SchoolID    uint   `json:"school_id"`
	Date        string `json:"date"`
	Status      bool   `json:"status"`
	ParentEmail string `json:"parent_email,omitempty"`
	StudentName string `json:"student_name,omitempty"`
}

type AttendanceResult struct {
	Job   AttendanceJob
	Error error
}

func StartAttendanceWorkerPool(
	workerCount int,
	jobs <-chan AttendanceJob,
	repo *Repository,
) <-chan AttendanceResult {
	results := make(chan AttendanceResult, 100)
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for w := 0; w < workerCount; w++ {
		go func(wID int) {
			defer wg.Done()
			for job := range jobs {
				a := &Attendance{
					StudentID: job.StudentID,
					SchoolID:  job.SchoolID,
					Date:      job.Date,
					Status:    job.Status,
				}

				err := repo.CreateAttendance(a)
				results <- AttendanceResult{Job: job, Error: err}
				if err != nil {
					fmt.Printf("[attendance worker %d] error for student %d: %v\n", wID, job.StudentID, err)
				}
			}
		}(w)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
