package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"app/controllers"
	"app/models"
	"app/pkg/middleware"
)
 
func main() {
      // DB接続設定
      dsn := "user=gorm password=gorm dbname=gorm host=db port=5432 sslmode=disable"
      db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
      })

      if err != nil {
            panic("failed to connect database")
      }
 
      fmt.Println("Connection Opened to Database")
 
      // 自動マイグレーション
      // Todoモデルの構造体の通りのスキーマを構築
      db.AutoMigrate(&models.Todo{}, &models.User{})
      // seeder.Seeder(db)
      
      // モデルとコントローラの初期化
      // モデルはデータベースとのやり取りを担当し、コントローラはクライアントからのリクエストを処理し、モデルを通じてデータベースとやり取りをします。
      todoModel := models.NewTodoModel(db)
      todoController := controllers.NewTodoController(todoModel)
      
      // ルーティング設定
      r := gin.Default()
      api := r.Group("/api")
      api.Use(middleware.AuthMiddleware)
      {
            api.GET("/todos", todoController.GetTodos)
            api.GET("/todos/:id", todoController.GetTodo)
            api.POST("/todos", todoController.CreateTodo)
            api.PUT("/todos/:id", todoController.UpdateTodo)
            api.DELETE("/todos/:id", todoController.DeleteTodo)
      
            api.GET("/users", todoController.GetUsers)
            api.GET("/users/:email", todoController.GetUser)
            api.PUT("/users/:email", todoController.UpdateUser)
            api.DELETE("/users/:email", todoController.DeleteUser)
      }

      // auth group
      auth := r.Group("/auth")
      
      auth.POST("/signup", todoController.SignUp)
      auth.POST("/login", todoController.Login)

      // サーバ起動
      r.Run()
}