package migrate

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"app/models"
)

// 使ってない
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Todo{},&models.User{}); err != nil {
		return err
	}

	todos := []models.Todo{
		{ID: 1, Title: "title1", Description: "description1", Category: "category1", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
		{ID: 2, Title: "title2", Description: "description2", Category: "category2", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
		{ID: 3, Title: "title3", Description: "description3", Category: "category3", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
	}

	users := []models.User{
		{ID: 1, Name: "name1", Email: "email1", Password: "password1"},
		{ID: 2, Name: "name2", Email: "email2", Password: "password2"},
		{ID: 3, Name: "name3", Email: "email3", Password: "password3"},
	}

	// データをデータベースに保存
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			return err
		}
	}

	fmt.Println("Migration did run successfully")

	return nil
}