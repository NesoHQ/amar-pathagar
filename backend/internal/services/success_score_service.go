package services

import (
	"database/sql"
	"fmt"
	"time"
)

type SuccessScoreService struct {
	db *sql.DB
}

func NewSuccessScoreService(db *sql.DB) *SuccessScoreService {
	return &SuccessScoreService{db: db}
}

const (
	ScoreReturnOnTime   = 10
	ScoreReturnLate     = -15
	ScorePositiveReview = 5
	ScoreNegativeReview = -10
	ScoreIdeaPosted     = 3
	ScoreIdeaUpvote     = 1
	ScoreIdeaDownvote   = -1
	ScoreLostBook       = -50
	ScoreBookDonated    = 20
	ScoreMoneyDonated   = 10
)

func (s *SuccessScoreService) UpdateScore(userID string, change int, reason string, refType string, refID *string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update user success score
	_, err = tx.Exec(`
		UPDATE users 
		SET success_score = success_score + $1,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`, change, userID)
	if err != nil {
		return err
	}

	// Record history
	var refIDVal interface{} = nil
	if refID != nil {
		refIDVal = *refID
	}

	_, err = tx.Exec(`
		INSERT INTO success_score_history (user_id, change_amount, reason, reference_type, reference_id)
		VALUES ($1, $2, $3, $4, $5)
	`, userID, change, reason, refType, refIDVal)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *SuccessScoreService) ProcessReturnOnTime(userID, bookID string) error {
	return s.UpdateScore(userID, ScoreReturnOnTime, "Returned book on time", "book", &bookID)
}

func (s *SuccessScoreService) ProcessReturnLate(userID, bookID string) error {
	return s.UpdateScore(userID, ScoreReturnLate, "Returned book late", "book", &bookID)
}

func (s *SuccessScoreService) ProcessPositiveReview(userID, reviewID string) error {
	return s.UpdateScore(userID, ScorePositiveReview, "Received positive review", "review", &reviewID)
}

func (s *SuccessScoreService) ProcessNegativeReview(userID, reviewID string) error {
	return s.UpdateScore(userID, ScoreNegativeReview, "Received negative review", "review", &reviewID)
}

func (s *SuccessScoreService) ProcessIdeaPosted(userID, ideaID string) error {
	return s.UpdateScore(userID, ScoreIdeaPosted, "Posted reading idea", "idea", &ideaID)
}

func (s *SuccessScoreService) ProcessIdeaUpvote(userID, ideaID string) error {
	return s.UpdateScore(userID, ScoreIdeaUpvote, "Idea received upvote", "idea", &ideaID)
}

func (s *SuccessScoreService) ProcessIdeaDownvote(userID, ideaID string) error {
	return s.UpdateScore(userID, ScoreIdeaDownvote, "Idea received downvote", "idea", &ideaID)
}

func (s *SuccessScoreService) ProcessLostBook(userID, bookID string) error {
	return s.UpdateScore(userID, ScoreLostBook, "Lost book", "book", &bookID)
}

func (s *SuccessScoreService) ProcessBookDonation(userID, donationID string) error {
	return s.UpdateScore(userID, ScoreBookDonated, "Donated book", "donation", &donationID)
}

func (s *SuccessScoreService) ProcessMoneyDonation(userID, donationID string) error {
	return s.UpdateScore(userID, ScoreMoneyDonated, "Made financial contribution", "donation", &donationID)
}

func (s *SuccessScoreService) GetScoreHistory(userID string, limit int) ([]map[string]interface{}, error) {
	query := `
		SELECT id, change_amount, reason, reference_type, reference_id, created_at
		FROM success_score_history
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := s.db.Query(query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []map[string]interface{}
	for rows.Next() {
		var id, reason, refType string
		var changeAmount int
		var refID sql.NullString
		var createdAt time.Time

		err := rows.Scan(&id, &changeAmount, &reason, &refType, &refID, &createdAt)
		if err != nil {
			return nil, err
		}

		entry := map[string]interface{}{
			"id":             id,
			"change_amount":  changeAmount,
			"reason":         reason,
			"reference_type": refType,
			"created_at":     createdAt.Format(time.RFC3339),
		}
		if refID.Valid {
			entry["reference_id"] = refID.String
		}
		history = append(history, entry)
	}

	return history, nil
}

func (s *SuccessScoreService) CanUserRequestBook(userID string) (bool, string, error) {
	var successScore int
	err := s.db.QueryRow("SELECT success_score FROM users WHERE id = $1", userID).Scan(&successScore)
	if err != nil {
		return false, "", err
	}

	if successScore < 20 {
		return false, fmt.Sprintf("Your success score (%d) is too low. Minimum required: 20", successScore), nil
	}

	return true, "", nil
}
