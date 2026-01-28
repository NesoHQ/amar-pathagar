package repository

import (
	"database/sql"

	"github.com/yourusername/online-library/internal/models"
)

type IdeaRepository struct {
	db *sql.DB
}

func NewIdeaRepository(db *sql.DB) *IdeaRepository {
	return &IdeaRepository{db: db}
}

func (r *IdeaRepository) Create(idea *models.ReadingIdea) error {
	return r.db.QueryRow(`
		INSERT INTO reading_ideas (book_id, user_id, title, content)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`, idea.BookID, idea.UserID, idea.Title, idea.Content).Scan(&idea.ID, &idea.CreatedAt, &idea.UpdatedAt)
}

func (r *IdeaRepository) GetByID(id string) (*models.ReadingIdea, error) {
	idea := &models.ReadingIdea{}
	err := r.db.QueryRow(`
		SELECT id, book_id, user_id, title, content, upvotes, downvotes, created_at, updated_at
		FROM reading_ideas WHERE id = $1
	`, id).Scan(&idea.ID, &idea.BookID, &idea.UserID, &idea.Title, &idea.Content,
		&idea.Upvotes, &idea.Downvotes, &idea.CreatedAt, &idea.UpdatedAt)
	return idea, err
}

func (r *IdeaRepository) GetByBook(bookID string, limit, offset int) ([]*models.ReadingIdea, error) {
	rows, err := r.db.Query(`
		SELECT id, book_id, user_id, title, content, upvotes, downvotes, created_at, updated_at
		FROM reading_ideas
		WHERE book_id = $1
		ORDER BY (upvotes - downvotes) DESC, created_at DESC
		LIMIT $2 OFFSET $3
	`, bookID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ideas []*models.ReadingIdea
	for rows.Next() {
		idea := &models.ReadingIdea{}
		err := rows.Scan(&idea.ID, &idea.BookID, &idea.UserID, &idea.Title, &idea.Content,
			&idea.Upvotes, &idea.Downvotes, &idea.CreatedAt, &idea.UpdatedAt)
		if err != nil {
			continue
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func (r *IdeaRepository) Vote(ideaID, userID, voteType string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Remove existing vote
	_, err = tx.Exec(`DELETE FROM idea_votes WHERE idea_id = $1 AND user_id = $2`, ideaID, userID)
	if err != nil {
		return err
	}

	// Add new vote
	_, err = tx.Exec(`
		INSERT INTO idea_votes (idea_id, user_id, vote_type)
		VALUES ($1, $2, $3)
	`, ideaID, userID, voteType)
	if err != nil {
		return err
	}

	// Update counts
	_, err = tx.Exec(`
		UPDATE reading_ideas
		SET upvotes = (SELECT COUNT(*) FROM idea_votes WHERE idea_id = $1 AND vote_type = 'upvote'),
		    downvotes = (SELECT COUNT(*) FROM idea_votes WHERE idea_id = $1 AND vote_type = 'downvote')
		WHERE id = $1
	`, ideaID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *IdeaRepository) GetUserVote(ideaID, userID string) (string, error) {
	var voteType string
	err := r.db.QueryRow(`
		SELECT vote_type FROM idea_votes WHERE idea_id = $1 AND user_id = $2
	`, ideaID, userID).Scan(&voteType)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return voteType, err
}
