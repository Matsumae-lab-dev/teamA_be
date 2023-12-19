package models

import (
	"app/requests"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
      //　フィールドは主キーとして機能し、gorm:"primary_key" タグによって指定されています。このフィールドは uint 型で、データベース上のレコードを一意に識別します。
    ID        uint       `gorm:"primary_key" json:"id"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
      // 論理削除をサポートするためのもので、gorm.DeletedAt 型で定義されています。このフィールドが NULL でない場合、レコードは削除されたことを示します。
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
 
type Todo struct {
      ID      uint   `gorm:"primary_key" json:"id"`
      Title   string `gorm:"not null" json:"title"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `gorm:"not null" json:"state"`
}
 
type TodoModel struct {
      DB *gorm.DB
}

// NewTodoModel 関数は TodoModel のコンストラクタ関数です。この関数は、*gorm.DB 型の引数を受け取り、その引数を使って新しい TodoModel インスタンスを生成して返します。
func NewTodoModel(db *gorm.DB) *TodoModel {
      return &TodoModel{DB: db}
}
func (m *TodoModel) GetAll() ([]Todo, error) {
      var todos []Todo
      // m.DB.Find(&todos) は GORM を使用してデータベースからメモを検索します。検索結果は todos スライスに格納されます。
      if err := m.DB.Find(&todos).Error; err != nil {
            return nil, err
      }
      fmt.Println(todos)
      return todos, nil
}
 
func (m *TodoModel) GetByID(id uint) (Todo, error) {
      var todo Todo
      // First：指定されたモデルに基づいて最初のレコードを検索します。
            // Where: 指定された条件に基づいてレコードをフィルタリングします。
      if err := m.DB.Where("id = ?", id).First(&todo).Error; err != nil {
            return Todo{}, err
      }
      return todo, nil
}
 
func (m *TodoModel) Create(todo requests.CreateTodoInput) (Todo, error) {
      fmt.Printf("%+v\n", todo)
      newTodo := Todo{
            Title:       todo.Title,
            Description: todo.Description,
            Category:    todo.Category,
            Deadline:    todo.Deadline,
            State:       todo.State,
      }
      if err := m.DB.Create(&newTodo).Error; err != nil {
            return Todo{}, err
      }
      return newTodo, nil
}
 
func (m *TodoModel) Update(id uint, todo requests.UpdateTodoInput) (Todo, error) {
      existingTodo, err := m.GetByID(id)
      if err != nil {
            return Todo{}, err
      }
      updatedTodo := requests.UpdateTodoInput{
            Title:       todo.Title,
            Description: todo.Description,
            Category:    todo.Category,
            Deadline:    todo.Deadline,
            State:       todo.State,
      }
      if err := m.DB.Model(&existingTodo).Updates(updatedTodo).Error; err != nil {
            return Todo{}, err
      }
      return existingTodo, nil
}
 
func (m *TodoModel) Delete(id uint) error {
      todo, err := m.GetByID(id)
      if err != nil {
            return err
      }
      return m.DB.Delete(&todo).Error
}