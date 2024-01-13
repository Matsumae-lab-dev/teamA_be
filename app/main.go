package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/controllers"
	"app/models"
	seeder "app/seeds"
)
 
func main() {
      // DB接続設定
      dsn := "user=gorm password=gorm dbname=gorm host=db port=5432 sslmode=disable"
      db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
      if err != nil {
            panic("failed to connect database")
      }
 
      fmt.Println("Connection Opened to Database")
 
      // 自動マイグレーション
      // Todoモデルの構造体の通りのスキーマを構築
      db.AutoMigrate(&models.Todo{}, &models.User{})
      seeder.Seeder(db)
      
      // モデルとコントローラの初期化
      // モデルはデータベースとのやり取りを担当し、コントローラはクライアントからのリクエストを処理し、モデルを通じてデータベースとやり取りをします。
      todoModel := models.NewTodoModel(db)
      todoController := controllers.NewTodoController(todoModel)
      
      // ルーティング設定
      r := gin.Default()
      r.GET("/todos", todoController.GetTodos)
      r.GET("/todos/:id", todoController.GetTodo)
      r.POST("/todos", todoController.CreateTodo)
      r.PUT("/todos/:id", todoController.UpdateTodo)
      r.DELETE("/todos/:id", todoController.DeleteTodo)

      // サーバ起動
      r.Run()
}