package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetPublicProfile(c *gin.Context) {
	userID := c.Param("id")

	var profile dto.UserPublicProfile
	err := h.db.QueryRow(`
		SELECT id, username, full_name, avatar_url, bio, success_score, 
			books_shared, books_received, reviews_received, ideas_posted, is_donor, created_at
		FROM users WHERE id = $1
	`, userID).Scan(&profile.ID, &profile.Username, &profile.FullName, &profile.AvatarURL,
		&profile.Bio, &profile.SuccessScore, &profile.BooksShared, &profile.BooksReceived,
		&profile.ReviewsReceived, &profile.IdeasPosted, &profile.IsDonor, &profile.JoinedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	var req dto.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	_, err := h.db.Exec(`
		UPDATE users 
		SET full_name = $1, bio = $2, avatar_url = $3, 
		    location_lat = $4, location_lng = $5, location_address = $6,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
	`, req.FullName, req.Bio, req.AvatarURL, req.LocationLat, req.LocationLng,
		req.LocationAddress, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Profile updated successfully", nil))
}

func (h *UserHandler) AddInterests(c *gin.Context) {
	userID := c.GetString("user_id")
	var req dto.AddInterestsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	tx, err := h.db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Database error"})
		return
	}
	defer tx.Rollback()

	for _, interest := range req.Interests {
		_, err := tx.Exec(`
			INSERT INTO user_interests (user_id, interest)
			VALUES ($1, $2)
			ON CONFLICT (user_id, interest) DO NOTHING
		`, userID, interest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to add interests"})
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to save interests"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Interests added successfully", nil))
}

func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	limit := 10

	// Top readers
	topReaders, _ := h.getTopUsers("books_received", limit)

	// Top sharers
	topSharers, _ := h.getTopUsers("books_shared", limit)

	// Top donors
	topDonors, _ := h.getTopDonors(limit)

	// Highest scores
	highestScores, _ := h.getTopUsers("success_score", limit)

	// Top idea writers
	topIdeaWriters, _ := h.getTopUsers("ideas_posted", limit)

	response := dto.LeaderboardResponse{
		TopReaders:     topReaders,
		TopSharers:     topSharers,
		TopDonors:      topDonors,
		HighestScores:  highestScores,
		TopIdeaWriters: topIdeaWriters,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) getTopUsers(orderBy string, limit int) ([]dto.UserPublicProfile, error) {
	query := `
		SELECT id, username, full_name, avatar_url, bio, success_score, 
			books_shared, books_received, reviews_received, ideas_posted, is_donor, created_at
		FROM users
		ORDER BY ` + orderBy + ` DESC
		LIMIT $1
	`

	rows, err := h.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.UserPublicProfile
	for rows.Next() {
		var user dto.UserPublicProfile
		err := rows.Scan(&user.ID, &user.Username, &user.FullName, &user.AvatarURL,
			&user.Bio, &user.SuccessScore, &user.BooksShared, &user.BooksReceived,
			&user.ReviewsReceived, &user.IdeasPosted, &user.IsDonor, &user.JoinedAt)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (h *UserHandler) getTopDonors(limit int) ([]dto.UserPublicProfile, error) {
	query := `
		SELECT DISTINCT u.id, u.username, u.full_name, u.avatar_url, u.bio, u.success_score, 
			u.books_shared, u.books_received, u.reviews_received, u.ideas_posted, u.is_donor, u.created_at
		FROM users u
		JOIN donations d ON u.id = d.donor_id
		WHERE d.is_public = true
		GROUP BY u.id
		ORDER BY COUNT(d.id) DESC
		LIMIT $1
	`

	rows, err := h.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.UserPublicProfile
	for rows.Next() {
		var user dto.UserPublicProfile
		err := rows.Scan(&user.ID, &user.Username, &user.FullName, &user.AvatarURL,
			&user.Bio, &user.SuccessScore, &user.BooksShared, &user.BooksReceived,
			&user.ReviewsReceived, &user.IdeasPosted, &user.IsDonor, &user.JoinedAt)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
