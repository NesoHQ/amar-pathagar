package repository

import (
	"database/sql"

	"github.com/yourusername/online-library/internal/models"
)

type BookmarkRepository struct {
	db *sql.DB
}

func NewBookmarkRepository(db *sql.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

func (r *BookmarkRepository) Create(bookmark *models.UserBookmark) error {
	return r.db.QueryRow(`
		INSERT INTO user_bookmarks (user_id, book_id, bookmark_type, priority_level)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, book_id, bookmark_type) 
		DO UPDATE SET priority_level = $4
		RETURNING id, created_at
	`, bookmark.UserID, bookmark.BookID, bookmark.BookmarkType, bookmark.PriorityLevel).
		Scan(&bookmark.ID, &bookmark.CreatedAt)
}

func (r *BookmarkRepository) Delete(userID, bookID, bookmarkType string) error {
	_, err := r.db.Exec(`
		DELETE FROM user_bookmarks 
		WHERE user_id = $1 AND book_id = $2 AND bookmark_type = $3
	`, userID, bookID, bookmarkType)
	return err
}

func (r *BookmarkRepository) GetByUser(userID, bookmarkType string) ([]*models.UserBookmark, error) {
	query := `
		SELECT id, user_id, book_id, bookmark_type, priority_level, created_at
		FROM user_bookmarks
		WHERE user_id = $1
	`
	args := []interface{}{userID}

	if bookmarkType != "" {
		query += " AND bookmark_type = $2"
		args = append(args, bookmarkType)
	}

	query += " ORDER BY priority_level DESC, created_at DESC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []*models.UserBookmark
	for rows.Next() {
		bookmark := &models.UserBookmark{}
		err := rows.Scan(&bookmark.ID, &bookmark.UserID, &bookmark.BookID, &bookmark.BookmarkType,
			&bookmark.PriorityLevel, &bookmark.CreatedAt)
		if err != nil {
			continue
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}

func (r *BookmarkRepository) CheckExists(userID, bookID, bookmarkType string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM user_bookmarks 
		WHERE user_id = $1 AND book_id = $2 AND bookmark_type = $3)
	`, userID, bookID, bookmarkType).Scan(&exists)
	return exists, err
}
