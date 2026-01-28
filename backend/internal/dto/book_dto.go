package dto

type CreateBookRequest struct {
	Title        string   `json:"title" binding:"required"`
	Author       string   `json:"author" binding:"required"`
	ISBN         string   `json:"isbn"`
	CoverURL     string   `json:"cover_url"`
	Description  string   `json:"description"`
	Category     string   `json:"category"`
	Tags         []string `json:"tags"`
	Topics       []string `json:"topics"`
	PhysicalCode string   `json:"physical_code" binding:"required"`
	IsDonated    bool     `json:"is_donated"`
}

type UpdateBookRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	ISBN        string   `json:"isbn"`
	CoverURL    string   `json:"cover_url"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Topics      []string `json:"topics"`
	Status      string   `json:"status"`
}

type BookRequestRequest struct {
	BookID string `json:"book_id" binding:"required"`
}
