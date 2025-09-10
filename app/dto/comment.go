package dto

type CommentCreateRequest struct {
	AuthorName string `json:"author_name" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type CommentResponse struct {
	ID         uint64 `json:"id"`
	BlogID     uint64 `json:"blog_id"`
	AuthorName string `json:"author_name"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
}
