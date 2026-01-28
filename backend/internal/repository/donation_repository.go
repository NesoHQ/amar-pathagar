package repository

import (
	"database/sql"

	"github.com/yourusername/online-library/internal/models"
)

type DonationRepository struct {
	db *sql.DB
}

func NewDonationRepository(db *sql.DB) *DonationRepository {
	return &DonationRepository{db: db}
}

func (r *DonationRepository) Create(donation *models.Donation) error {
	return r.db.QueryRow(`
		INSERT INTO donations (donor_id, donation_type, book_id, amount, currency, message, is_public)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`, donation.DonorID, donation.DonationType, donation.BookID, donation.Amount,
		donation.Currency, donation.Message, donation.IsPublic).
		Scan(&donation.ID, &donation.CreatedAt)
}

func (r *DonationRepository) GetAll(limit, offset int, publicOnly bool) ([]*models.Donation, error) {
	query := `
		SELECT id, donor_id, donation_type, book_id, amount, currency, message, is_public, created_at
		FROM donations
	`
	if publicOnly {
		query += " WHERE is_public = true"
	}
	query += " ORDER BY created_at DESC LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var donations []*models.Donation
	for rows.Next() {
		donation := &models.Donation{}
		err := rows.Scan(&donation.ID, &donation.DonorID, &donation.DonationType, &donation.BookID,
			&donation.Amount, &donation.Currency, &donation.Message, &donation.IsPublic, &donation.CreatedAt)
		if err != nil {
			continue
		}
		donations = append(donations, donation)
	}
	return donations, nil
}

func (r *DonationRepository) GetByDonor(donorID string) ([]*models.Donation, error) {
	rows, err := r.db.Query(`
		SELECT id, donor_id, donation_type, book_id, amount, currency, message, is_public, created_at
		FROM donations
		WHERE donor_id = $1
		ORDER BY created_at DESC
	`, donorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var donations []*models.Donation
	for rows.Next() {
		donation := &models.Donation{}
		err := rows.Scan(&donation.ID, &donation.DonorID, &donation.DonationType, &donation.BookID,
			&donation.Amount, &donation.Currency, &donation.Message, &donation.IsPublic, &donation.CreatedAt)
		if err != nil {
			continue
		}
		donations = append(donations, donation)
	}
	return donations, nil
}
