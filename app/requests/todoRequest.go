package requests

import "time"

type CreateTodoInput struct {
      Title string `json:"title" binding:"required"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `json:"state"`
}
 
type UpdateTodoInput struct {
      Title string `json:"title"`
      Description string `json:"description"`
      Category string `json:"category"`
      Deadline time.Time `json:"deadline"`
      State bool `json:"state"`
}