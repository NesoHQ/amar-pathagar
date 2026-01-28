package dto

type CreateUserReviewRequest struct {
	RevieweeID          string `json:"reviewee_id" binding:"required"`
	BookID              string `json:"book_id"`
	BehaviorRating      *int   `json:"behavior_rating" binding:"omitempty,min=1,max=5"`
	BookConditionRating *int   `json:"book_condition_rating" binding:"omitempty,min=1,max=5"`
	CommunicationRating *int   `json:"communication_rating" binding:"omitempty,min=1,max=5"`
	Comment             string `json:"comment"`
}

type UserReviewResponse struct {
	ID                  string  `json:"id"`
	ReviewerID          string  `json:"reviewer_id"`
	ReviewerName        string  `json:"reviewer_name"`
	RevieweeID          string  `json:"reviewee_id"`
	RevieweeName        string  `json:"reviewee_name"`
	BookID              *string `json:"book_id"`
	BookTitle           *string `json:"book_title"`
	BehaviorRating      *int    `json:"behavior_rating"`
	BookConditionRating *int    `json:"book_condition_rating"`
	CommunicationRating *int    `json:"communication_rating"`
	AverageRating       float64 `json:"average_rating"`
	Comment             string  `json:"comment"`
	CreatedAt           string  `json:"created_at"`
}
