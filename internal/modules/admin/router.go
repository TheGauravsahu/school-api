package admin

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/modules/student"
)

type Handler struct {
	StudentHandler *student.Handler
}

func Router(h *Handler) {
	http.HandleFunc("POST /api/admin/students", h.StudentHandler.CreateStudent)
}
