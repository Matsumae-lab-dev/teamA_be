package seeder

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"app/models"
)

// Seeder 関数はデータベースに初期データを投入するための関数です。
func Seeder(db *gorm.DB) error {
      // Todo モデルを使用してデータを作成
      todos := []models.Todo{
            {Title: "title1", Description: "description1", Category: "category1", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
            {Title: "title2", Description: "description2", Category: "category2", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
            {Title: "title3", Description: "description3", Category: "category3", Deadline: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), State: false},
      }

      user := []models.User{
            {Name: "name1", Email: "email1", Password: "password1"},
            {Name: "name2", Email: "email2", Password: "password2"},
            {Name: "name3", Email: "email3", Password: "password3"},
      }
 
      // データをデータベースに保存
      for _, user := range user {
            if err := db.Create(&user).Error; err != nil {
                  return err
            }
      }
      for _, todo := range todos {
            if err := db.Create(&todo).Error; err != nil {
                  return err
            }
      }
 
      fmt.Println("Seeder executed successfully.")
 
      return nil
}