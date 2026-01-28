package repository

import (
	"database/sql"

	"github.com/yourusername/online-library/internal/models"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(review *models.UserReview) error {
	return r.db.QueryRow(`
		INSERT INTO user_reviews (reviewer_id, reviewee_id, book_id, behavior_rating, 
			book_condition_rating, communication_rating, comment)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`, review.ReviewerID, review.RevieweeID, review.BookID, review.BehaviorRating,
		review.BookConditionRating, review.CommunicationRating, review.Comment).
		Scan(&review.ID, &review.CreatedAt)
}

func (r *ReviewRepository) GetByReviewee(revieweeID string, limit, offset int) ([]*models.UserReview, error) {
	rows, err := r.db.Query(`
		SELECT id, reviewer_id, reviewee_id, book_id, behavior_rating, 
			book_condition_rating, communication_rating, comment, created_at
		FROM user_reviews
		WHERE reviewee_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, revieweeID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*models.UserReview
	for rows.Next() {
		review := &models.UserReview{}
		err := rows.Scan(&review.ID, &review.ReviewerID, &review.RevieweeID, &review.BookID,
			&review.BehaviorRating, &review.BookConditionRating, &review.CommunicationRating,
			&review.Comment, &review.CreatedAt)
		if err != nil {
			continue
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (r *ReviewRepository) GetAverageRatings(revieweeID string) (map[string]float64, error) {
	var behaviorAvg, conditionAvg, communicationAvg sql.NullFloat64
	err := r.db.QueryRow(`
		SELECT 
			AVG(behavior_rating) as behavior_avg,
			AVG(book_condition_rating) as condition_avg,
			AVG(communication_rating) as communication_avg
		FROM user_reviews
		WHERE reviewee_id = $1
	`, revieweeID).Scan(&behaviorAvg, &conditionAvg, &communicationAvg)

	if err != nil {
		return nil, err
	}

	return map[string]float64{
		"behavior":      behaviorAvg.Float64,
		"condition":     conditionAvg.Float64,
		"communication": communicationAvg.Float64,
	}, nil
}
