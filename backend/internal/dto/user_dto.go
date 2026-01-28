package dto

type UpdateProfileRequest struct {
	FullName        string   `json:"full_name"`
	Bio             string   `json:"bio"`
	AvatarURL       string   `json:"avatar_url"`
	LocationLat     *float64 `json:"location_lat"`
	LocationLng     *float64 `json:"location_lng"`
	LocationAddress string   `json:"location_address"`
}

type UserPublicProfile struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	FullName        string `json:"full_name"`
	AvatarURL       string `json:"avatar_url"`
	Bio             string `json:"bio"`
	SuccessScore    int    `json:"success_score"`
	BooksShared     int    `json:"books_shared"`
	BooksReceived   int    `json:"books_received"`
	ReviewsReceived int    `json:"reviews_received"`
	IdeasPosted     int    `json:"ideas_posted"`
	IsDonor         bool   `json:"is_donor"`
	JoinedAt        string `json:"joined_at"`
}

type AddInterestsRequest struct {
	Interests []string `json:"interests" binding:"required"`
}

type LeaderboardResponse struct {
	TopReaders     []UserPublicProfile `json:"top_readers"`
	TopSharers     []UserPublicProfile `json:"top_sharers"`
	TopDonors      []UserPublicProfile `json:"top_donors"`
	HighestScores  []UserPublicProfile `json:"highest_scores"`
	TopIdeaWriters []UserPublicProfile `json:"top_idea_writers"`
}
