package admin

import (
	"github.com/TheGauravsahu/school-api/internal/modules/student"
	"github.com/TheGauravsahu/school-api/internal/modules/teacher"
)

type Handler struct {
	StudentHandler *student.Handler
	TeacherHandler *teacher.Handler
}

func NewHandler(studentHandler *student.Handler, teacherHandler *teacher.Handler) *Handler {
	return &Handler{
		StudentHandler: studentHandler,
		TeacherHandler: teacherHandler,
	}
}
