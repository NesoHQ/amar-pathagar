package dto

type CreateDonationRequest struct {
	DonationType string   `json:"donation_type" binding:"required,oneof=book money"`
	BookID       *string  `json:"book_id"`
	Amount       *float64 `json:"amount"`
	Currency     string   `json:"currency"`
	Message      string   `json:"message"`
	IsPublic     bool     `json:"is_public"`
}

type DonationResponse struct {
	ID           string   `json:"id"`
	DonorID      string   `json:"donor_id"`
	DonorName    string   `json:"donor_name"`
	DonationType string   `json:"donation_type"`
	BookID       *string  `json:"book_id"`
	BookTitle    *string  `json:"book_title"`
	Amount       *float64 `json:"amount"`
	Currency     string   `json:"currency"`
	Message      string   `json:"message"`
	CreatedAt    string   `json:"created_at"`
}
