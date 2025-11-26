package admin

import "github.com/TheGauravsahu/school-api/internal/modules/student"

type Handler struct {
	StudentHandler *student.Handler
}

func NewHandler(studentHandler *student.Handler) *Handler {
	return &Handler{
		StudentHandler: studentHandler,
	}
}
