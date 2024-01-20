package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"app/models"
	"app/pkg/utils"
	"app/requests"

	"github.com/gin-gonic/gin"
)
 
type TodoController struct {
      Model *models.TodoModel
}
 
// NewTodoController関数はTodoModelを引数として受け取り、それを使用してTodoControllerを初期化します。
// これは依存性注入の一例で、テストやモックの作成が容易になります。この方式を使用すると、テスト中に実際のデータベースを使用する代わりにモックデータベースを注入できます。これにより、テストの可読性とメンテナンス性が向上します。
func NewTodoController(m *models.TodoModel) *TodoController {
      return &TodoController{Model: m}
}
 
// gin.ContextはGinの中心的な部分で、リクエストとレスポンスの情報を含んでいます
func (mc *TodoController) GetTodos(c *gin.Context) {
      // models/todo.goのGetAll関数で全件取得
      todos, err := mc.Model.GetTodoAll()
      if err != nil {
            // 500エラーを返す
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
      // JSONメソッドは、HTTPレスポンスをJSON形式で生成するためのメソッド
      // gin.HはGinが提供する便利な関数で、map[string]interface{}型のマップを短く書くためのものです。この場合、"data": todosはクライアントに返すJSONのキーと値を設定しています。
      c.JSON(http.StatusOK, gin.H{"data": todos})
}
 
func (mc *TodoController) GetTodo(c *gin.Context) {
      // strconv.Atoi→文字列を整数に変換
      id, err := strconv.Atoi(c.Param("id"))
      if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
      }
 
      // GetByID関数はuint型を引数として受け取るのでuinit型に変換
      // uint型: 0および正の整数のみを表現できます
      todo, err := mc.Model.GetTodoByID(uint(id))
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
 
      c.JSON(http.StatusOK, gin.H{"data": todo})
}
 
func (mc *TodoController) CreateTodo(c *gin.Context) {

      var input requests.CreateTodoInput

      // ShouldBindJSONメソッドは、HTTPリクエストのボディからJSONデータを構造体またはマップにバインドするためのものです。
      // map[string]interface{}{"name": "John","age": 30,"email": "john@example.com",}← マップにバインド
      if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
      }
 
      // 入力されたcontentを引数に
      todo, err := mc.Model.CreateTodo(input)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
 
      c.JSON(http.StatusOK, gin.H{"data": todo})
}
 
func (mc *TodoController) UpdateTodo(c *gin.Context) {
      id, err := strconv.Atoi(c.Param("id"))
      if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
      }
 
      var input requests.UpdateTodoInput
      if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
      }
 
      todo, err := mc.Model.UpdateTodo(uint(id), input)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
 
      c.JSON(http.StatusOK, gin.H{"data": todo})
}
 
func (mc *TodoController) DeleteTodo(c *gin.Context) {
      id, err := strconv.Atoi(c.Param("id"))
      if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
      }
 
      if err := mc.Model.DeleteTodo(uint(id)); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
 
      c.JSON(http.StatusOK, gin.H{"data": true})
}

func (mc *TodoController) GetUsers(c *gin.Context) {
      users, err := mc.Model.GetAllUser()
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
      c.JSON(http.StatusOK, gin.H{"data": users})
}

func (mc *TodoController) GetUser(c *gin.Context) {
      email := c.Param("email")
      fmt.Printf("%+v\n", email)
 
      user, err := mc.Model.GetUserByEmail(email)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
 
      c.JSON(http.StatusOK, gin.H{"data": user})
}

func (mc *TodoController) UpdateUser(c *gin.Context) {
      email := c.Param("email")
      fmt.Printf("%+v\n", email)
 
      var input requests.UpdateUserInput
      if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
      }

      user, err := mc.Model.UpdateUser(email, input)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }

      c.JSON(http.StatusOK, gin.H{"data": user})
}

func (mc *TodoController) DeleteUser(c *gin.Context) {
      email := c.Param("email")
      fmt.Printf("%+v\n", email)
 
      if err := mc.Model.DeleteUserByEmail(email); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
 
      c.JSON(http.StatusOK, gin.H{"data": true})
}

// signUp login logout
func (mc *TodoController) SignUp(c *gin.Context) {
      var input requests.CreateUserInput
      if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
      }
      
      user, err := mc.Model.CreateUser(input)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }
      fmt.Printf("%+v\n", user)

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

      // Cookieの有効期限を設定
      cookieMaxAge := 60 * 60 * 24 * 30 // 30日

    // Cookieにトークンをセット
	c.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)

      c.JSON(http.StatusOK, gin.H{"data": user})
}

func (mc *TodoController) Login(c *gin.Context) {
      var input requests.LoginInput
      if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
      }
      
      user, err := mc.Model.LoginUser(input)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }

      err = mc.Model.VerifyPassword(user, input.Password)
      if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
      }

      token, err := utils.GenerateToken(user.ID)
      if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                  "message": "Failed to login",
            })
            return
      }

      // Cookieの有効期限を設定
      cookieMaxAge := 60 * 60 * 24 * 30 // 30日
      // Cookieにトークンをセット
      c.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)

      c.JSON(http.StatusOK, gin.H{
            "data": map[string]interface{}{
                  "message": "login success",
                  "user": user,
            },
      })
}