package teacher

type Teacher struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;unique"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Subject   string `gorm:"not null"`
}