package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/models"
)
 
func main() {
      // DB接続設定
      dsn := "user=gorm password=gorm dbname=gorm host=db port=5432 sslmode=disable"
      db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
      if err != nil {
            panic("failed to connect database")
      }
 
      fmt.Println("Connection Opened to Database")
 
      // シーダーを実行
      if err := Seeder(db); err != nil {
            fmt.Println("Seeder error:", err)
            return
      }
}
 
// Seeder 関数はデータベースに初期データを投入するための関数です。
func Seeder(db *gorm.DB) error {
      // Memo モデルを使用してデータを作成
      memos := []models.Memo{
            {Content: "Memo 1"},
            {Content: "Memo 2"},
            {Content: "Memo 3"},
      }
 
      // データをデータベースに保存
      for _, memo := range memos {
            if err := db.Create(&memo).Error; err != nil {
                  return err
            }
      }
 
      fmt.Println("Seeder executed successfully.")
 
      return nil
}