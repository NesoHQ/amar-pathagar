package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID              string          `json:"id"`
	Username        string          `json:"username"`
	Email           string          `json:"email"`
	PasswordHash    string          `json:"-"`
	FullName        string          `json:"full_name"`
	Role            string          `json:"role"`
	AvatarURL       string          `json:"avatar_url"`
	Bio             string          `json:"bio"`
	LocationLat     sql.NullFloat64 `json:"location_lat,omitempty"`
	LocationLng     sql.NullFloat64 `json:"location_lng,omitempty"`
	LocationAddress string          `json:"location_address"`
	SuccessScore    int             `json:"success_score"`
	BooksShared     int             `json:"books_shared"`
	BooksReceived   int             `json:"books_received"`
	ReviewsReceived int             `json:"reviews_received"`
	IdeasPosted     int             `json:"ideas_posted"`
	TotalUpvotes    int             `json:"total_upvotes"`
	TotalDownvotes  int             `json:"total_downvotes"`
	IsDonor         bool            `json:"is_donor"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleMember UserRole = "member"
)
