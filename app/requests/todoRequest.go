package requests

import "time"

type GetTodoOutput struct {
      ID uint `json:"id"`
      Title string `json:"title"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `json:"state"`
      Users []AuthOutput `json:"users"`
}

type CreateTodoInput struct {
      Title string `json:"title" binding:"required"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `json:"state"`
      Email string `json:"email" binding:"required"`
}
 
type UpdateTodoInput struct {
      ID uint `json:"id"`
      Title string `json:"title"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `json:"state"`
}

type CreateUserInput struct {
      Name string `json:"name" binding:"required"`
      Email string `json:"email" binding:"required"`
      Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
      Name string `json:"name"`
      Email string `json:"email"`
      Password string `json:"password"`
}

type AuthInput struct {
      Email string `json:"email" binding:"required"`
      Password string `json:"password" binding:"required"`
}

type AuthOutput struct {
      ID uint `json:"id"`
      Name string `json:"name"`
      Email string `json:"email"`
}