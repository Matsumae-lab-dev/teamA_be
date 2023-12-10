package requests

// contetnフィールドは必須
type CreateMemoInput struct {
      Content string `json:"content" binding:"required"`
}
 
type UpdateMemoInput struct {
      Content string `json:"content"`
}