package attendance

type Attendance struct {
	StudentID uint   `json:"student_id"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}