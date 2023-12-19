package migrate

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"app/models"
)

// Migrate 関数はデータベースのマイグレーションを行うための関数です。
func Migrate(db *gorm.DB) error {
	// Todo モデルを使用してデータベースのテーブルを作成
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}

	todos := []models.Todo{
		{ID: 1, Title: "title1", Description: "description1", Category: "category1", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
		{ID: 2, Title: "title2", Description: "description2", Category: "category2", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
		{ID: 3, Title: "title3", Description: "description3", Category: "category3", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
	}

	// データをデータベースに保存
	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			return err
		}
	}

	fmt.Println("Migration did run successfully")

	return nil
}