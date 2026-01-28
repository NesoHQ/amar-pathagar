package services

import (
	"database/sql"
	"math"
)

type MatchingService struct {
	db *sql.DB
}

func NewMatchingService(db *sql.DB) *MatchingService {
	return &MatchingService{db: db}
}

// Calculate distance between two points using Haversine formula
func (m *MatchingService) calculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadius = 6371 // km

	dLat := (lat2 - lat1) * math.Pi / 180
	dLng := (lng2 - lng1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLng/2)*math.Sin(dLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

// Calculate interest match score between user and book
func (m *MatchingService) calculateInterestMatch(userID, bookID string) (float64, error) {
	// Get user interests
	userInterests := make(map[string]float64)
	rows, err := m.db.Query(`
		SELECT interest, weight FROM user_interests WHERE user_id = $1
	`, userID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var interest string
		var weight float64
		if err := rows.Scan(&interest, &weight); err != nil {
			return 0, err
		}
		userInterests[interest] = weight
	}

	// Get book topics
	var topics []string
	err = m.db.QueryRow(`
		SELECT topics FROM books WHERE id = $1
	`, bookID).Scan(&topics)
	if err != nil {
		return 0, err
	}

	// Calculate match score
	var matchScore float64
	for _, topic := range topics {
		if weight, exists := userInterests[topic]; exists {
			matchScore += weight
		}
	}

	// Normalize to 0-100
	if len(topics) > 0 {
		matchScore = (matchScore / float64(len(topics))) * 100
	}

	return matchScore, nil
}

// Calculate priority score for book request
func (m *MatchingService) CalculatePriorityScore(userID, bookID string, holderLat, holderLng float64) (float64, float64, error) {
	// Get user data
	var successScore int
	var userLat, userLng sql.NullFloat64
	err := m.db.QueryRow(`
		SELECT success_score, location_lat, location_lng 
		FROM users WHERE id = $1
	`, userID).Scan(&successScore, &userLat, &userLng)
	if err != nil {
		return 0, 0, err
	}

	// Calculate distance
	var distance float64
	if userLat.Valid && userLng.Valid {
		distance = m.calculateDistance(holderLat, holderLng, userLat.Float64, userLng.Float64)
	} else {
		distance = 10000 // Default high distance if location not set
	}

	// Calculate interest match
	interestScore, err := m.calculateInterestMatch(userID, bookID)
	if err != nil {
		return 0, 0, err
	}

	// Calculate priority score
	// Formula: (SuccessScore * 0.4) + (InterestMatch * 0.3) + (DistanceScore * 0.3)
	// Distance score: closer = higher (inverse relationship)
	distanceScore := 100.0
	if distance > 0 {
		distanceScore = math.Max(0, 100-(distance/10)) // 10km = 10 points reduction
	}

	priorityScore := (float64(successScore)*0.4 + interestScore*0.3 + distanceScore*0.3)

	return priorityScore, interestScore, nil
}

// Select best match from multiple requests
func (m *MatchingService) SelectBestMatch(bookID string) (string, error) {
	query := `
		SELECT br.id, br.user_id, br.priority_score
		FROM book_requests br
		WHERE br.book_id = $1 AND br.status = 'pending'
		ORDER BY br.priority_score DESC, br.requested_at ASC
		LIMIT 1
	`

	var requestID, userID string
	var priorityScore float64
	err := m.db.QueryRow(query, bookID).Scan(&requestID, &userID, &priorityScore)
	if err != nil {
		return "", err
	}

	return requestID, nil
}

// Update request priority scores when book holder changes
func (m *MatchingService) UpdateRequestPriorities(bookID string) error {
	// Get current holder location
	var holderLat, holderLng sql.NullFloat64
	err := m.db.QueryRow(`
		SELECT u.location_lat, u.location_lng
		FROM books b
		JOIN users u ON b.current_holder_id = u.id
		WHERE b.id = $1
	`, bookID).Scan(&holderLat, &holderLng)
	if err != nil {
		return err
	}

	if !holderLat.Valid || !holderLng.Valid {
		return nil // Skip if no location
	}

	// Get all pending requests
	rows, err := m.db.Query(`
		SELECT id, user_id FROM book_requests
		WHERE book_id = $1 AND status = 'pending'
	`, bookID)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Update each request
	for rows.Next() {
		var requestID, userID string
		if err := rows.Scan(&requestID, &userID); err != nil {
			continue
		}

		priorityScore, interestScore, err := m.CalculatePriorityScore(userID, bookID, holderLat.Float64, holderLng.Float64)
		if err != nil {
			continue
		}

		_, err = m.db.Exec(`
			UPDATE book_requests
			SET priority_score = $1, interest_match_score = $2
			WHERE id = $3
		`, priorityScore, interestScore, requestID)
		if err != nil {
			continue
		}
	}

	return nil
}
