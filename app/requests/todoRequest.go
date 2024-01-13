package requests

import "time"

type CreateTodoInput struct {
      Title string `json:"title" binding:"required"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `json:"state"`
      Email string `json:"email" binding:"required"`
}
 
type UpdateTodoInput struct {
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