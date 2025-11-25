package student

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
)

type Handler struct {
	UserRepo    *user.Repository
	StudentRepo *Repository
}

type CreateStudentRequest struct {
	SchoolID  uint   `json:"school_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassID   uint   `json:"class_id"`
	Section   string `json:"section"`
	RollNo    int    `json:"roll_no"`
}

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var body CreateStudentRequest
	if ok := utils.ParseJson(w, r, &body); !ok {
		return
	}

	hashed, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}

	newUser := &user.User{
		Username: body.Username,
		Password: hashed,
		Role:     "STUDENT",
		SchoolID: body.SchoolID,
	}

	if err := h.UserRepo.CreateUser(newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	student := &Student{
		UserID:    newUser.ID,
		SchoolID:  body.SchoolID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Age:       body.Age,
		ClassID:   body.ClassID,
		Section:   body.Section,
		RollNo:    body.RollNo,
	}

	if err := h.StudentRepo.CreateStudent(student); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]any{
		"message": "Student created successfully.",
		"student": student,
	})

}

func (h *Handler) ImportStudents(w http.ResponseWriter, r *http.Request) {
	var jobsSlice []ImportJob
	if ok := utils.ParseJson(w, r, &jobsSlice); !ok {
		return
	}

	if len(jobsSlice) == 0 {
		utils.WriteError(w, http.StatusBadRequest, "no student data provided")
		return
	}

	// create channel
	jobs := make(chan ImportJob, len(jobsSlice))
	results := StartWokerPool(10, jobs, h.StudentRepo, h.UserRepo)

	for _, j := range jobsSlice {
		jobs <- j
	}
	close(jobs)

	var failed int
	var processed int
	var errs []string
	for res := range results {
		processed++
		if res.Error != nil {
			failed++
			errs = append(errs, fmt.Sprintf("%s %s: %v", res.Job.FirstName, res.Job.LastName, res.Error))
		}
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"processed": processed,
		"failed":    failed,
		"errors":    errs,
	})
}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	schoolIDParam := r.URL.Query().Get("schoolId")

	// fetch by school
	if schoolIDParam != "" {
		schoolID, err := strconv.Atoi(schoolIDParam)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, "Invalid schoolId")
			return
		}

		students, err := h.StudentRepo.FindBySchool(uint(schoolID))
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Could not fetch students.")
			return
		}

		utils.WriteJson(w, http.StatusOK, map[string]any{
			"message":  "Fetched students by school.",
			"students": students,
		})
		return
	}

	students, err := h.StudentRepo.FindAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Could not fetch students.")
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]any{
		"message":  "Fetched all students.",
		"students": students,
	})
}

func (h *Handler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	student, err := h.StudentRepo.GetStudentByID(uint(id))
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Student not found.")
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]any{
		"message": "Student retrived successfully.",
		"student": student,
	})
}

func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	var body map[string]interface{}
	if ok := utils.ParseJson(w, r, &body); !ok {
		return
	}

	if err := h.StudentRepo.UpdateStudent(uint(id), body); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to update student")
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]any{"message": "Updated successfully"})
}

func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.StudentRepo.DeleteStudent(uint(id)); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})
}
