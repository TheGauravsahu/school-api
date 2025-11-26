package teacher

import (
	"fmt"
	"sync"

	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
	"gorm.io/gorm"
)

type ImportJob struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Subject   string `json:"subject"`
	PhoneNo   string `json:"phone"`
	SchoolID  uint   `json:"school_id"`
	ClassID   uint   `json:"class_id"`
}

type JobResult struct {
	Job   ImportJob
	Error error
}

func StartWokerPool(
	workerCount int,
	jobs <-chan ImportJob,
	teacherRepo *Repository,
	userRepo *user.Repository,
) <-chan JobResult {
	results := make(chan JobResult, 100)
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for w := 0; w < workerCount; w++ {
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				err := processSingleTeacher(job, teacherRepo, userRepo)
				results <- JobResult{Job: job, Error: err}
				if err != nil {
					fmt.Printf("[worker %d] error: %v\n", workerID, err)
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

func processSingleTeacher(
	job ImportJob,
	teacherRepo *Repository,
	userRepo *user.Repository) error {
	// generate username & password
	username := utils.GenerateUsername(job.FirstName, job.LastName, int(job.SchoolID))
	password := utils.GeneratePassword()

	// hash
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	teacherRepo.db.Transaction(func(tx *gorm.DB) error {
		// create new user
		user := &user.User{
			SchoolID: job.SchoolID,
			Username: username,
			Password: hashed,
			Role:     "TEACHER",
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// create teacher
		newTeacher := &Teacher{
			UserID:    user.ID,
			SchoolID:  job.SchoolID,
			Email:     job.Email,
			FirstName: job.FirstName,
			LastName:  job.LastName,
			PhoneNo:   job.PhoneNo,
			Subject:   job.Subject,
			ClassID:   job.ClassID,
		}

		if err := tx.Create(newTeacher).Error; err != nil {
			return err
		}

		return nil
	})

	if job.Email != "" {
		utils.SendWelcomeEmail(job.Email, username, password)
	}

	return nil
}
