package config

import (
	"log"

	"github.com/TheGauravsahu/school-api/internal/modules/school"
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
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")
}
