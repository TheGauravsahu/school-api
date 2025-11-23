package class

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique"`
	TeacherID uint   `gorm:"not null"`
}