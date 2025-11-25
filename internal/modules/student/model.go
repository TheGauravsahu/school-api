package student

type Student struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Email     string `gom:"not null"`
	UserID    uint   `gorm:"not null;unique"`
	SchoolID  uint   `gorm:"not null"`
	ClassID   uint   `gorm:"not null"`
	Section   string
	RollNo    int `gorm:"not null"`
}
