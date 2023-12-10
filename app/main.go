package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/controllers"
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
 
      // 自動マイグレーション
      // Memoモデルの構造体の通りのスキーマを構築(id、content、created_at、updated_at、deleted_at)
      db.AutoMigrate(&models.Memo{})
 
      // モデルとコントローラの初期化
      // モデルはデータベースとのやり取りを担当し、コントローラはクライアントからのリクエストを処理し、モデルを通じてデータベースとやり取りをします。
      memoModel := models.NewMemoModel(db)
      memoController := controllers.NewMemoController(memoModel)
 
      // ルーティング設定
      r := gin.Default()
      r.GET("/memos", memoController.GetMemos)
      r.GET("/memos/:id", memoController.GetMemo)
      r.POST("/memos", memoController.CreateMemo)
      r.PUT("/memos/:id", memoController.UpdateMemo)
      r.DELETE("/memos/:id", memoController.DeleteMemo)
 
      // サーバ起動
      r.Run()
}