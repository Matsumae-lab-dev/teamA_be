package models

import (
	"time"

	"gorm.io/gorm"
)

//
type BaseModel struct {
      //　フィールドは主キーとして機能し、gorm:"primary_key" タグによって指定されています。このフィールドは uint 型で、データベース上のレコードを一意に識別します。
    ID        uint       `gorm:"primary_key" json:"id"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
      // 論理削除をサポートするためのもので、gorm.DeletedAt 型で定義されています。このフィールドが NULL でない場合、レコードは削除されたことを示します。
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
 
type Memo struct {
      BaseModel
      Content string `json:"content"`
}
 
type MemoModel struct {
      DB *gorm.DB
}
 
// NewMemoModel 関数は MemoModel のコンストラクタ関数です。この関数は、*gorm.DB 型の引数を受け取り、その引数を使って新しい MemoModel インスタンスを生成して返します。
func NewMemoModel(db *gorm.DB) *MemoModel {
      return &MemoModel{DB: db}
}
 
func (m *MemoModel) GetAll() ([]Memo, error) {
      var memos []Memo
      // m.DB.Find(&memos) は GORM を使用してデータベースからメモを検索します。検索結果は memos スライスに格納されます。
      if err := m.DB.Find(&memos).Error; err != nil {
            return nil, err
      }
      return memos, nil
}
 
func (m *MemoModel) GetByID(id uint) (Memo, error) {
      var memo Memo
      // First：指定されたモデルに基づいて最初のレコードを検索します。
            // Where: 指定された条件に基づいてレコードをフィルタリングします。
      if err := m.DB.Where("id = ?", id).First(&memo).Error; err != nil {
            return Memo{}, err
      }
      return memo, nil
}
 
func (m *MemoModel) Create(content string) (Memo, error) {
      memo := Memo{Content: content}
      // Create：新しいレコードを作成します。
      if err := m.DB.Create(&memo).Error; err != nil {
            return Memo{}, err
      }
      return memo, nil
}
 
func (m *MemoModel) Update(id uint, content string) (Memo, error) {
      memo, err := m.GetByID(id)
      if err != nil {
            return Memo{}, err
      }
      if err := m.DB.Model(&memo).Update("content", content).Error; err != nil {
            return Memo{}, err
      }
      return memo, nil
}
 
func (m *MemoModel) Delete(id uint) error {
      memo, err := m.GetByID(id)
      if err != nil {
            return err
      }
      return m.DB.Delete(&memo).Error
}