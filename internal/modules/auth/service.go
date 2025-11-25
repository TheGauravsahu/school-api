package auth

import (
	"github.com/TheGauravsahu/school-api/internal/modules/school"
	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
)

type Service struct {
	SchoolRepo *school.Repository
	UserRepo   *user.Repository
}

func NewService(schoolRepo *school.Repository, userRepo *user.Repository) *Service {
	return &Service{
		SchoolRepo: schoolRepo,
		UserRepo:   userRepo,
	}
}

func (s *Service) RegisterSchoolAndAdmin(schoolName, address, logo, username, password string) error {
	// create school
	newSchool := &school.School{
		Name:    schoolName,
		Address: address,
		Logo:    logo,
	}
	if err := s.SchoolRepo.CreateSchool(newSchool); err != nil {
		return err
	}

	hashPass, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// create admin user
	admin := &user.User{
		SchoolID: newSchool.ID,
		Username: username,
		Password: hashPass,
		Role:     "ADMIN",
	}
	if err := s.UserRepo.CreateUser(admin); err != nil {
		return err
	}

	return nil

}
