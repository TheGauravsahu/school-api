package admin

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/middlewares"
	"github.com/TheGauravsahu/school-api/internal/modules/student"
)

type Handler struct {
	StudentHandler *student.Handler
}

func Router(h *Handler) {
	http.HandleFunc("POST /api/admin/students", middlewares.AuthMiddleware(h.StudentHandler.CreateStudent, "ADMIN", "TEACHER"))
	http.HandleFunc("POST /api/admin/students/import", middlewares.AuthMiddleware(h.StudentHandler.ImportStudents, "ADMIN", "TEACHER"))
}
