package dto

type CreateBookRequest struct {
	Title        string   `json:"title" binding:"required"`
	Author       string   `json:"author" binding:"required"`
	ISBN         string   `json:"isbn"`
	CoverURL     string   `json:"cover_url"`
	Description  string   `json:"description"`
	Category     string   `json:"category"`
	Tags         []string `json:"tags"`
	PhysicalCode string   `json:"physical_code" binding:"required"`
}

type UpdateBookRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	ISBN        string   `json:"isbn"`
	CoverURL    string   `json:"cover_url"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
}

type AssignBookRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type FinishReadingRequest struct {
	Notes  string `json:"notes"`
	Rating int    `json:"rating" binding:"min=1,max=5"`
	Review string `json:"review"`
}

type BookListResponse struct {
	Books      []BookDTO `json:"books"`
	Total      int       `json:"total"`
	Page       int       `json:"page"`
	PageSize   int       `json:"page_size"`
	TotalPages int       `json:"total_pages"`
}

type BookDTO struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	ISBN          string   `json:"isbn"`
	CoverURL      string   `json:"cover_url"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	Tags          []string `json:"tags"`
	PhysicalCode  string   `json:"physical_code"`
	Status        string   `json:"status"`
	CurrentHolder *UserDTO `json:"current_holder,omitempty"`
	QueueLength   int      `json:"queue_length"`
	TotalReads    int      `json:"total_reads"`
	AverageRating float64  `json:"average_rating"`
	CreatedAt     string   `json:"created_at"`
}
