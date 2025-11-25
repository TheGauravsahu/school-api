package student

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateStudent(student *Student) error {
	return r.db.Create(student).Error
}

func (r *Repository) GetStudentByID(id uint) (*Student, error) {
	var stu Student
	if err := r.db.First(&stu, id).Error; err != nil {
		return nil, err
	}
	return &stu, nil
}

func (r *Repository) FindBySchool(schoolID uint) ([]Student, error) {
	var students []Student
	if err := r.db.Where("school_id = ?", schoolID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *Repository) FindBySchoolAndClass(schoolID uint, classID uint) ([]Student, error) {
	var students []Student
	if err := r.db.Where("school_id = ? AND class_id = ?", schoolID, classID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *Repository) FindAll() ([]Student, error) {
	var students []Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *Repository) DeleteStudent(id uint) error {
	return r.db.Delete(&Student{}, id).Error
}

func (r *Repository) UpdateStudent(id uint, updates map[string]interface{}) error {
	return r.db.Model(&Student{}).Where("id = ?", id).Updates(updates).Error
}
