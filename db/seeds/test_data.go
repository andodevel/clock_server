package seeds

import (
	"github.com/andodevel/go-echo-template/models"
	"github.com/jinzhu/gorm"
)

// SeedTestData ...Seed test data: Users + ...
func SeedTestData(DB *gorm.DB) error {
	// Users
	var users = []models.User{
		models.User{Name: "An Do", Username: "andodevel"},
	}

	for _, user := range users {
		DB.Create(&user)
	}

	return nil
}
