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
	http.HandleFunc("GET /api/admin/students", middlewares.AuthMiddleware(h.StudentHandler.GetAllStudents, "ADMIN", "TEACHER"))
	http.HandleFunc("GET /api/admin/students/{id}", middlewares.AuthMiddleware(h.StudentHandler.GetStudentById, "ADMIN", "TEACHER"))
	http.HandleFunc("PUT /api/admin/students/{id}", middlewares.AuthMiddleware(h.StudentHandler.UpdateStudent, "ADMIN"))
	http.HandleFunc("DELETE /api/admin/students/{id}", middlewares.AuthMiddleware(h.StudentHandler.DeleteStudent, "ADMIN"))
}
