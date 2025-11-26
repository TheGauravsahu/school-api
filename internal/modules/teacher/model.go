package teacher

import "time"

type Teacher struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;unique"`
	SchoolID  uint   `gorm:"not null"`
	ClassID   uint   `gorm:"not null"`
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Subject   string `gorm:"not null"`
	PhoneNo   string `gorm:"not null" json:"phone"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
