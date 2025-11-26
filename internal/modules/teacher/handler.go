package teacher

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
	"gorm.io/gorm"
)

type Handler struct {
	TeacherRepo *Repository
	UserRepos   *user.Repository
}

func NewHandler(teacherRepo *Repository, userRepo *user.Repository) *Handler {
	return &Handler{
		TeacherRepo: teacherRepo,
		UserRepos:   userRepo,
	}
}

type CreateTeacherRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Subject   string `json:"subject"`
	PhoneNo   string `json:"phone"`
	SchoolID  uint   `json:"school_id"`
	ClassID   uint   `json:"class_id"`
}

func (h *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var body CreateTeacherRequest
	if ok := utils.ParseJson(w, r, &body); !ok {
		return
	}

	username := utils.GenerateUsername(body.FirstName, body.LastName, int(body.SchoolID))
	password := utils.GeneratePassword()
	hashChan := make(chan string)
	go func() {
		hashed, _ := utils.HashPassword(password)
		hashChan <- hashed
	}()
	hashedPassword := <-hashChan

	var createdTeacher *Teacher
	err := h.TeacherRepo.db.Transaction(func(tx *gorm.DB) error {
		user := &user.User{
			Username: username,
			Password: hashedPassword,
			Role:     "TEACHER",
			SchoolID: body.SchoolID,
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// Create Teacher
		teacher := &Teacher{
			UserID:    user.ID,
			SchoolID:  body.SchoolID,
			Email:     body.Email,
			FirstName: body.FirstName,
			LastName:  body.LastName,
			PhoneNo:   body.PhoneNo,
			Subject:   body.Subject,
			ClassID:   body.ClassID,
		}
		if err := tx.Create(teacher).Error; err != nil {
			return err
		}

		createdTeacher = teacher
		return nil
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Could not create teacher")
		return
	}

	go func(email, username, p string) {
		utils.SendWelcomeEmail(email, username, p)
	}(createdTeacher.Email, username, password)

	utils.WriteJson(w, http.StatusCreated, map[string]interface{}{
		"message": "Teacher created successfully",
		"teacher": createdTeacher,
	})

}
