package dto

type CreateBookRequestRequest struct {
	BookID string `json:"book_id" binding:"required"`
}

type BookRequestResponse struct {
	ID                 string   `json:"id"`
	BookID             string   `json:"book_id"`
	BookTitle          string   `json:"book_title"`
	UserID             string   `json:"user_id"`
	Username           string   `json:"username"`
	Status             string   `json:"status"`
	PriorityScore      float64  `json:"priority_score"`
	InterestMatchScore float64  `json:"interest_match_score"`
	DistanceKm         *float64 `json:"distance_km"`
	RequestedAt        string   `json:"requested_at"`
	ProcessedAt        *string  `json:"processed_at"`
	DueDate            *string  `json:"due_date"`
}

type ProcessRequestRequest struct {
	RequestID string `json:"request_id" binding:"required"`
	Action    string `json:"action" binding:"required,oneof=approve reject"`
	DueDays   int    `json:"due_days"`
}
