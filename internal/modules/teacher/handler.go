package teacher

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
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
	FirstName string `josn:"first_name"`
	LastName  string `josn:"last_name"`
	Email     string `json:"email"`
	Subject   string `josn:"subject"`
	PhoneNo   string `josn:"phone"`
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
	hashed, _ := utils.HashPassword(password)

	user := &user.User{
		Username: username,
		Password: hashed,
		Role:     "TEACHER",
		SchoolID: body.SchoolID,
	}
	h.UserRepos.CreateUser(user)

	teacher := &Teacher{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		PhoneNo:   body.PhoneNo,
		Subject:   body.Subject,
		SchoolID:  body.SchoolID,
		UserID:    user.ID,
	}
	h.TeacherRepo.CreateTeacher(teacher)

	if teacher.Email != "" {
		utils.SendWelcomeEmail(teacher.Email, username, password)
	}

	utils.WriteJson(w, http.StatusCreated, map[string]interface{}{
		"message": "Teacher created sucessfully.",
		"teacher": teacher,
	})
}
