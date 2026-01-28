package services

import (
	"database/sql"
	"fmt"
)

type NotificationService struct {
	db *sql.DB
}

func NewNotificationService(db *sql.DB) *NotificationService {
	return &NotificationService{db: db}
}

func (n *NotificationService) Create(userID, notifType, title, message, link string) error {
	_, err := n.db.Exec(`
		INSERT INTO notifications (user_id, type, title, message, link)
		VALUES ($1, $2, $3, $4, $5)
	`, userID, notifType, title, message, link)
	return err
}

func (n *NotificationService) NotifyBookAvailable(userID, bookID, bookTitle string) error {
	return n.Create(
		userID,
		"book_available",
		"Book Available",
		fmt.Sprintf("The book '%s' is now available for you!", bookTitle),
		fmt.Sprintf("/books/%s", bookID),
	)
}

func (n *NotificationService) NotifyRequestApproved(userID, bookID, bookTitle string) error {
	return n.Create(
		userID,
		"request_approved",
		"Request Approved",
		fmt.Sprintf("Your request for '%s' has been approved!", bookTitle),
		fmt.Sprintf("/books/%s", bookID),
	)
}

func (n *NotificationService) NotifyRequestRejected(userID, bookID, bookTitle string) error {
	return n.Create(
		userID,
		"request_rejected",
		"Request Not Approved",
		fmt.Sprintf("Your request for '%s' was not approved this time.", bookTitle),
		fmt.Sprintf("/books/%s", bookID),
	)
}

func (n *NotificationService) NotifyReturnDue(userID, bookID, bookTitle string, daysLeft int) error {
	return n.Create(
		userID,
		"return_due",
		"Book Return Reminder",
		fmt.Sprintf("Please return '%s' in %d days.", bookTitle, daysLeft),
		fmt.Sprintf("/books/%s", bookID),
	)
}

func (n *NotificationService) NotifyReviewReceived(userID, reviewerName string) error {
	return n.Create(
		userID,
		"review_received",
		"New Review",
		fmt.Sprintf("%s left you a review!", reviewerName),
		"/profile/reviews",
	)
}

func (n *NotificationService) NotifySuccessScoreChange(userID string, change int, reason string) error {
	action := "increased"
	if change < 0 {
		action = "decreased"
	}
	return n.Create(
		userID,
		"success_score_change",
		"Success Score Updated",
		fmt.Sprintf("Your success score %s by %d points: %s", action, abs(change), reason),
		"/profile",
	)
}

func (n *NotificationService) NotifyIdeaVote(userID, voterName, ideaTitle string, isUpvote bool) error {
	voteType := "upvoted"
	if !isUpvote {
		voteType = "downvoted"
	}
	return n.Create(
		userID,
		"idea_vote",
		"Idea Vote",
		fmt.Sprintf("%s %s your idea '%s'", voterName, voteType, ideaTitle),
		"/ideas",
	)
}

func (n *NotificationService) GetUserNotifications(userID string, limit int, unreadOnly bool) ([]map[string]interface{}, error) {
	query := `
		SELECT id, type, title, message, link, is_read, created_at
		FROM notifications
		WHERE user_id = $1
	`
	if unreadOnly {
		query += " AND is_read = false"
	}
	query += " ORDER BY created_at DESC LIMIT $2"

	rows, err := n.db.Query(query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []map[string]interface{}
	for rows.Next() {
		var id, notifType, title, message, link string
		var isRead bool
		var createdAt string

		err := rows.Scan(&id, &notifType, &title, &message, &link, &isRead, &createdAt)
		if err != nil {
			continue
		}

		notifications = append(notifications, map[string]interface{}{
			"id":         id,
			"type":       notifType,
			"title":      title,
			"message":    message,
			"link":       link,
			"is_read":    isRead,
			"created_at": createdAt,
		})
	}

	return notifications, nil
}

func (n *NotificationService) MarkAsRead(notificationID string) error {
	_, err := n.db.Exec(`
		UPDATE notifications SET is_read = true WHERE id = $1
	`, notificationID)
	return err
}

func (n *NotificationService) MarkAllAsRead(userID string) error {
	_, err := n.db.Exec(`
		UPDATE notifications SET is_read = true WHERE user_id = $1
	`, userID)
	return err
}

func (n *NotificationService) GetUnreadCount(userID string) (int, error) {
	var count int
	err := n.db.QueryRow(`
		SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND is_read = false
	`, userID).Scan(&count)
	return count, err
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
