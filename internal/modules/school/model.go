package school

import "time"

type School struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique"`
	Address   string `gorm:"not null"`
	Logo      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt  time.Time
}
