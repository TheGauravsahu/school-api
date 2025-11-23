package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *User) error {
	// check if exists
	existingUser, err := r.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// create new user
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByUsername(username string) (*User, error) {
	var user User
	result := r.db.Where("username = ?", username).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}


