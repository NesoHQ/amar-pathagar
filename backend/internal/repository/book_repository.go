package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/yourusername/online-library/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `
		INSERT INTO books (title, author, isbn, cover_url, description, category, tags, topics, physical_code, status, created_by, donated_by, is_donated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(
		query,
		book.Title,
		book.Author,
		book.ISBN,
		book.CoverURL,
		book.Description,
		book.Category,
		pq.Array(book.Tags),
		pq.Array(book.Topics),
		book.PhysicalCode,
		book.Status,
		book.CreatedBy,
		book.DonatedBy,
		book.IsDonated,
	).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

func (r *BookRepository) FindByID(id string) (*models.Book, error) {
	book := &models.Book{}
	query := `
		SELECT id, title, author, COALESCE(isbn, '') as isbn, COALESCE(cover_url, '') as cover_url, 
		       COALESCE(description, '') as description, COALESCE(category, '') as category, 
		       COALESCE(tags, '{}') as tags, COALESCE(topics, '{}') as topics, 
		       physical_code, status, current_holder_id, created_by, donated_by, 
		       is_donated, donation_date, total_reads, average_rating, created_at, updated_at
		FROM books
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.CoverURL,
		&book.Description,
		&book.Category,
		pq.Array(&book.Tags),
		pq.Array(&book.Topics),
		&book.PhysicalCode,
		&book.Status,
		&book.CurrentHolderID,
		&book.CreatedBy,
		&book.DonatedBy,
		&book.IsDonated,
		&book.DonationDate,
		&book.TotalReads,
		&book.AverageRating,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("book not found")
	}
	return book, err
}

func (r *BookRepository) FindAll(filters map[string]interface{}) ([]*models.Book, error) {
	query := `
		SELECT id, title, author, COALESCE(isbn, '') as isbn, COALESCE(cover_url, '') as cover_url, 
		       COALESCE(description, '') as description, COALESCE(category, '') as category, 
		       COALESCE(tags, '{}') as tags, COALESCE(topics, '{}') as topics, 
		       physical_code, status, current_holder_id, created_by, donated_by, 
		       is_donated, donation_date, total_reads, average_rating, created_at, updated_at
		FROM books
		WHERE 1=1
	`
	args := []interface{}{}
	argCount := 1

	if status, ok := filters["status"].(string); ok && status != "" {
		query += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, status)
		argCount++
	}

	if category, ok := filters["category"].(string); ok && category != "" {
		query += fmt.Sprintf(" AND category = $%d", argCount)
		args = append(args, category)
		argCount++
	}

	if search, ok := filters["search"].(string); ok && search != "" {
		query += fmt.Sprintf(" AND (title ILIKE $%d OR author ILIKE $%d)", argCount, argCount)
		args = append(args, "%"+search+"%")
		argCount++
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*models.Book{}
	for rows.Next() {
		book := &models.Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.CoverURL,
			&book.Description,
			&book.Category,
			pq.Array(&book.Tags),
			pq.Array(&book.Topics),
			&book.PhysicalCode,
			&book.Status,
			&book.CurrentHolderID,
			&book.CreatedBy,
			&book.DonatedBy,
			&book.IsDonated,
			&book.DonationDate,
			&book.TotalReads,
			&book.AverageRating,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) Update(book *models.Book) error {
	query := `
		UPDATE books
		SET title = $1, author = $2, isbn = $3, cover_url = $4, description = $5, 
		    category = $6, tags = $7, topics = $8, status = $9, current_holder_id = $10,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $11
	`
	_, err := r.db.Exec(
		query,
		book.Title,
		book.Author,
		book.ISBN,
		book.CoverURL,
		book.Description,
		book.Category,
		pq.Array(book.Tags),
		pq.Array(book.Topics),
		book.Status,
		book.CurrentHolderID,
		book.ID,
	)
	return err
}

func (r *BookRepository) Delete(id string) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
