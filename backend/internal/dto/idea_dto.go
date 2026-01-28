package dto

type CreateIdeaRequest struct {
	BookID  string `json:"book_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type VoteIdeaRequest struct {
	VoteType string `json:"vote_type" binding:"required,oneof=upvote downvote"`
}

type IdeaResponse struct {
	ID        string `json:"id"`
	BookID    string `json:"book_id"`
	BookTitle string `json:"book_title"`
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Upvotes   int    `json:"upvotes"`
	Downvotes int    `json:"downvotes"`
	NetScore  int    `json:"net_score"`
	CreatedAt string `json:"created_at"`
}
