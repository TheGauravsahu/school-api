package attendance

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/middlewares"
)

func Router(h *Handler) {
	http.HandleFunc("POST /api/attendance/mark", middlewares.AuthMiddleware(h.MarkAttendance, "ADMIN", "TEACER"))
}
