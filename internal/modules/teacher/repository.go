package teacher

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateTeacher(t *Teacher) error {
	return r.db.Create(t).Error
}

func (r *Repository) FindById(id uint) (Teacher, error) {
	var teacher Teacher
	err := r.db.First(&teacher).Error
	return teacher, err
}

func (r *Repository) FindAll() ([]Teacher, error) {
	var teachers []Teacher
	if err := r.db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *Repository) FindBySchool(schoolID uint) ([]Teacher, error) {
	var teachers []Teacher
	if err := r.db.Where("school_id = ?", schoolID).Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *Repository) DeleteTeacher(id uint) error {
	return r.db.Delete(&Teacher{}, id).Error
}

func (r *Repository) UpdateTeacher(id uint, updates map[string]interface{}) error {
	return r.db.Model(&Teacher{}).Where("id = ?", id).Updates(updates).Error
}
