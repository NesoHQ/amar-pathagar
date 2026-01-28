package dto

type CreateBookmarkRequest struct {
	BookID        string `json:"book_id" binding:"required"`
	BookmarkType  string `json:"bookmark_type" binding:"required,oneof=like bookmark priority"`
	PriorityLevel int    `json:"priority_level"`
}

type BookmarkResponse struct {
	ID            string `json:"id"`
	BookID        string `json:"book_id"`
	BookTitle     string `json:"book_title"`
	BookmarkType  string `json:"bookmark_type"`
	PriorityLevel int    `json:"priority_level"`
	CreatedAt     string `json:"created_at"`
}
