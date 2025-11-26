package config

import (
	"log"

	"github.com/TheGauravsahu/school-api/internal/modules/attendance"
	"github.com/TheGauravsahu/school-api/internal/modules/school"
	"github.com/TheGauravsahu/school-api/internal/modules/student"
	"github.com/TheGauravsahu/school-api/internal/modules/teacher"
	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("school.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	// AutoMigrate ALL models
	err = DB.AutoMigrate(
		&school.School{},
		&user.User{},
		&student.Student{},
		&attendance.Attendance{},
		&teacher.Teacher{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")
}
