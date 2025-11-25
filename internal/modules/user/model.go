package user

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SchoolID     uint      `gorm:"not null" json:"school_id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Password     string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"type:text CHECK(role IN ('ADMIN','SUPERADMIN','TEACHER','STUDENT','PARENT'));not null" json:"role"`
	RefreshToken string    `gorm:"type:text" json:"-"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
