package student

import (
	"fmt"
	"sync"

	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
)

type ImportJob struct {
	SchoolID  uint   `json:"school_id"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassID   uint   `json:"class_id"`
	Section   string `json:"section"`
	RollNo    int    `json:"roll_no"`
	Email     string `json:"email"`
}

type JobResult struct {
	Job   ImportJob
	Error error
}

func StartWokerPool(
	workerCount int,
	jobs <-chan ImportJob,
	studentRepo *Repository,
	userRepo *user.Repository,
) <-chan JobResult {
	results := make(chan JobResult, 100)
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for w := 0; w < workerCount; w++ {
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				err := processSingleStudent(job, studentRepo, userRepo)
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

func processSingleStudent(
	job ImportJob,
	studentRepo *Repository,
	userRepo *user.Repository) error {
	// generate username & password
	username := utils.GenerateUsername(job.FirstName, job.LastName, job.RollNo)
	password := utils.GeneratePassword()

	// hash
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// create new user
	newUser := &user.User{
		SchoolID: job.SchoolID,
		Username: username,
		Password: hashed,
		Role:     "STUDENT",
	}

	if err := userRepo.CreateUser(newUser); err != nil {
		return err
	}

	// create student
	newStudent := &Student{
		UserID:    newUser.ID,
		SchoolID:  job.SchoolID,
		FirstName: job.FirstName,
		LastName:  job.LastName,
		Age:       job.Age,
		ClassID:   job.ClassID,
		Section:   job.Section,
		RollNo:    job.RollNo,
		Email:     job.Email,
	}

	if err := studentRepo.CreateStudent(newStudent); err != nil {
		return err
	}

	if job.Email != "" {
		utils.SendWelcomeEmail(job.Email, username, password)
	}

	return nil
}
