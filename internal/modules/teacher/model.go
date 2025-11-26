package teacher

import "time"

type Teacher struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;unique"`
	SchoolID  uint   `gorm:"not null"`
	ClassID   uint   `gorm:"not null"`
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Subject   string `gorm:"not null"`
	PhoneNo   string `gorm:"not null"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
