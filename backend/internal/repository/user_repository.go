package repository

import (
	"database/sql"
	"fmt"

	"github.com/yourusername/online-library/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash, full_name, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash, full_name, role, 
		       COALESCE(avatar_url, '') as avatar_url, 
		       COALESCE(bio, '') as bio,
		       location_lat, location_lng, 
		       COALESCE(location_address, '') as location_address,
		       COALESCE(success_score, 100) as success_score,
		       COALESCE(books_shared, 0) as books_shared,
		       COALESCE(books_received, 0) as books_received,
		       COALESCE(reviews_received, 0) as reviews_received,
		       COALESCE(ideas_posted, 0) as ideas_posted,
		       COALESCE(total_upvotes, 0) as total_upvotes,
		       COALESCE(total_downvotes, 0) as total_downvotes,
		       COALESCE(is_donor, false) as is_donor,
		       created_at, updated_at
		FROM users
		WHERE username = $1
	`
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.Role,
		&user.AvatarURL,
		&user.Bio,
		&user.LocationLat,
		&user.LocationLng,
		&user.LocationAddress,
		&user.SuccessScore,
		&user.BooksShared,
		&user.BooksReceived,
		&user.ReviewsReceived,
		&user.IdeasPosted,
		&user.TotalUpvotes,
		&user.TotalDownvotes,
		&user.IsDonor,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash, full_name, role, 
		       COALESCE(avatar_url, '') as avatar_url, 
		       COALESCE(bio, '') as bio,
		       location_lat, location_lng, 
		       COALESCE(location_address, '') as location_address,
		       COALESCE(success_score, 100) as success_score,
		       COALESCE(books_shared, 0) as books_shared,
		       COALESCE(books_received, 0) as books_received,
		       COALESCE(reviews_received, 0) as reviews_received,
		       COALESCE(ideas_posted, 0) as ideas_posted,
		       COALESCE(total_upvotes, 0) as total_upvotes,
		       COALESCE(total_downvotes, 0) as total_downvotes,
		       COALESCE(is_donor, false) as is_donor,
		       created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.Role,
		&user.AvatarURL,
		&user.Bio,
		&user.LocationLat,
		&user.LocationLng,
		&user.LocationAddress,
		&user.SuccessScore,
		&user.BooksShared,
		&user.BooksReceived,
		&user.ReviewsReceived,
		&user.IdeasPosted,
		&user.TotalUpvotes,
		&user.TotalDownvotes,
		&user.IsDonor,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash, full_name, role, 
		       COALESCE(avatar_url, '') as avatar_url, 
		       COALESCE(bio, '') as bio,
		       location_lat, location_lng, 
		       COALESCE(location_address, '') as location_address,
		       COALESCE(success_score, 100) as success_score,
		       COALESCE(books_shared, 0) as books_shared,
		       COALESCE(books_received, 0) as books_received,
		       COALESCE(reviews_received, 0) as reviews_received,
		       COALESCE(ideas_posted, 0) as ideas_posted,
		       COALESCE(total_upvotes, 0) as total_upvotes,
		       COALESCE(total_downvotes, 0) as total_downvotes,
		       COALESCE(is_donor, false) as is_donor,
		       created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.Role,
		&user.AvatarURL,
		&user.Bio,
		&user.LocationLat,
		&user.LocationLng,
		&user.LocationAddress,
		&user.SuccessScore,
		&user.BooksShared,
		&user.BooksReceived,
		&user.ReviewsReceived,
		&user.IdeasPosted,
		&user.TotalUpvotes,
		&user.TotalDownvotes,
		&user.IsDonor,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}
