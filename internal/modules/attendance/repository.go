package attendance

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateAttendance(a *Attendance) error {
	result := r.db.Model(&Attendance{}).
		Where("student_id = ? AND date = ?", a.StudentID, a.Date).
		Updates(map[string]interface{}{
			"status":     a.Status,
			"updated_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return r.db.Create(a).Error
	}
	return nil
}

func (r *Repository) GetAttendanceByStudent(stuID uint, from, to time.Time) ([]Attendance, error) {
	var list []Attendance
	err := r.db.Where("student_id = ? AND date BETWEEN ? AND ?", stuID, from, to).Find(&list).Error
	return list, err
}

func (r *Repository) GetAttendanceBySchool(schoolID uint, date time.Time) ([]Attendance, error) {
	var list []Attendance
	err := r.db.Where("school_id = ? AND date = ?", schoolID, date).Find(&list).Error
	return list, err
}
