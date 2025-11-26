package teacher

type Handler struct {
	TeacherRepo *Repository
}

func NewHandler(teacherRepo *Repository) *Handler {
	return &Handler{
		TeacherRepo: teacherRepo,
	}
}
