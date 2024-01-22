package models

import (
	"app/requests"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation" // 追加
	"github.com/go-ozzo/ozzo-validation/is"         // 追加

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
      Users    []*User `gorm:"many2many:user_todos;"`
}

type User struct {
      ID      uint   `gorm:"primary_key" json:"id"`
      Name   string `gorm:"not null" json:"name"`
      Email  string `gorm:"unique;not null" json:"email"`
      Password string `gorm:"not null" json:"password"`
}
 
type TodoModel struct {
      DB *gorm.DB
}

// NewTodoModel 関数は TodoModel のコンストラクタ関数です。この関数は、*gorm.DB 型の引数を受け取り、その引数を使って新しい TodoModel インスタンスを生成して返します。
func NewTodoModel(db *gorm.DB) *TodoModel {
      return &TodoModel{DB: db}
}

func (m *TodoModel) GetTodoAll() ([]Todo, error) {
      var todos []Todo
      // m.DB.Find(&todos) は GORM を使用してデータベースからメモを検索します。検索結果は todos スライスに格納されます。
      if err := m.DB.Preload("Users").Find(&todos).Error; err != nil {
            return nil, err
      }
      fmt.Println(todos)
      return todos, nil
}
 
func (m *TodoModel) GetTodoByID(id uint) (Todo, error) {
      var todo Todo
      // First：指定されたモデルに基づいて最初のレコードを検索します。
      // Where: 指定された条件に基づいてレコードをフィルタリングします。
      if err := m.DB.Preload("Users").Where("id = ?", id).First(&todo).Error; err != nil {
            return Todo{}, err
      }
      return todo, nil
}
 
func (m *TodoModel) CreateTodo(todo requests.CreateTodoInput) (Todo, error) {
      fmt.Printf("%+v\n", todo)
      newTodo := Todo{
            Title:       todo.Title,
            Description: todo.Description,
            Category:    todo.Category,
            Deadline:    todo.Deadline,
            State:       todo.State,
      }

      relationUser,err := m.GetUserByEmail(todo.Email)
      if err != nil {
            return Todo{}, err
      }
      if err := m.DB.Create(&newTodo).Error; err != nil {
            return Todo{}, err
      }
      // 中間テーブルの関係を作成
      if err := m.DB.Model(&newTodo).Association("Users").Append(&User{
            ID: relationUser.ID, 
            Email: relationUser.Email,
            Name: relationUser.Name,
            Password: relationUser.Password,
            }); err != nil {
            return Todo{}, err
      }
      return newTodo, nil
}
 
func (m *TodoModel) UpdateTodo(id uint, todo requests.UpdateTodoInput) (Todo, error) {
      existingTodo, err := m.GetTodoByID(id)
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
 
func (m *TodoModel) DeleteTodo(id uint) error {
      todo, err := m.GetTodoByID(id)
      if err != nil {
            return err
      }
      return m.DB.Delete(&todo).Error
}

func (m *TodoModel) CreateUser(user requests.CreateUserInput) (User, error) {
      fmt.Printf("%+v\n", user)

      // 既存のユーザーが存在するか確認
      _, err := m.GetUserByEmail(user.Email)
      if err == nil {
            // 既存のユーザーが見つかった場合はエラーを返す
            return User{}, fmt.Errorf("User with email %s already exists", user.Email)
      }
      
      newUser := User{
            Name:       user.Name,
            Email:      user.Email,
            Password:    user.Password,
      }

      // バリデーション
      if err := newUser.ValidateUser(); err != nil {
            return User{}, err
      }

      if err := m.DB.Create(&newUser).Error; err != nil {
            return User{}, err
      }
      return newUser, nil
}

func (m *TodoModel) GetAllUser() ([]User, error) {
      var users []User
      if err := m.DB.Find(&users).Error; err != nil {
            return nil, err
      }
      fmt.Println(users)
      return users, nil
}

func (m *TodoModel) GetUserByEmail(email string) (User, error) {
      var user User
      // First：指定されたモデルに基づいて最初のレコードを検索します。
            // Where: 指定された条件に基づいてレコードをフィルタリングします。
      if err := m.DB.Where("email = ?", email).First(&user).Error; err != nil {
            return User{}, err
      }
      return user, nil
}

func (m *TodoModel) UpdateUser(email string, user requests.UpdateUserInput) (User, error) { 
      existingUser, err := m.GetUserByEmail(email)
      if err != nil {
            return User{}, err
      }
      updatedUser := requests.UpdateUserInput{
            Name:       user.Name,
            Email: user.Email,
            Password:    user.Password,
      }
      if err := m.DB.Model(&existingUser).Updates(updatedUser).Error; err != nil {
            return User{}, err
      }
      return existingUser, nil
}

func (m *TodoModel) DeleteUserByEmail(email string) error {
      user, err := m.GetUserByEmail(email)
      if err != nil {
            return err
      }
      return m.DB.Delete(&user).Error
}

func (user *User) ValidateUser() error {
      err := validation.ValidateStruct(user,
            validation.Field(&user.Name,
                  validation.Required.Error("Name is requred"),
                  validation.Length(1, 255).Error("Name is too long"),
            ),
            validation.Field(&user.Email,
                  validation.Required.Error("Email is required"),
                  is.Email.Error("Email is invalid format"),
            ),
            validation.Field(&user.Password,
                  validation.Required.Error("Password is required"),
                  validation.Length(8, 255).Error("Password is less than 7 chars or more than 256 chars"),
            ),
      )
      return err
}


func (m *TodoModel) LoginUser(user requests.AuthInput) (User, error) {
      var loginUser User
      if err := m.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&loginUser).Error; err != nil {
            return User{}, err
      }
      return loginUser, nil
}

func (m *TodoModel) VerifyPassword(user User, password string) error {
      if user.Password != password {
            return fmt.Errorf("Password is invalid")
      }
      return nil
}

func (m *TodoModel) ConvertTodoToOutput(todo Todo) (requests.GetTodoOutput) {
      var users []requests.AuthOutput
      for _, user := range todo.Users {
            users = append(users, requests.AuthOutput{
                  ID:    user.ID,
                  Name:  user.Name,
                  Email: user.Email,
            })
      }
      return requests.GetTodoOutput{
            ID:          todo.ID,
            Title:       todo.Title,
            Description: todo.Description,
            Category:    todo.Category,
            Deadline:    todo.Deadline,
            State:       todo.State,
            Users:       users,
      }
}

func (m *TodoModel) ConvertTodosToOutput(todos []Todo) ([]requests.GetTodoOutput) {
      var output []requests.GetTodoOutput
      for _, todo := range todos {
            output = append(output, m.ConvertTodoToOutput(todo))
      }
      return output
}