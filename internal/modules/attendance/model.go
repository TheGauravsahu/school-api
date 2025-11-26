package attendance

import "time"

type Attendance struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	StudentID uint   `json:"student_id" gorm:"not null"`
	SchoolID  uint   `gorm:"not null;index" json:"school_id"`
	Date      string `json:"date" gorm:"not null"`
	Status    bool   `json:"status"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
