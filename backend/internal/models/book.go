package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Author          string         `json:"author"`
	ISBN            string         `json:"isbn"`
	CoverURL        string         `json:"cover_url"`
	Description     string         `json:"description"`
	Category        string         `json:"category"`
	Tags            []string       `json:"tags"`
	Topics          []string       `json:"topics"`
	PhysicalCode    string         `json:"physical_code"`
	Status          BookStatus     `json:"status"`
	CurrentHolderID sql.NullString `json:"current_holder_id"`
	CurrentHolder   *User          `json:"current_holder,omitempty"`
	CreatedBy       sql.NullString `json:"created_by"`
	DonatedBy       sql.NullString `json:"donated_by"`
	IsDonated       bool           `json:"is_donated"`
	DonationDate    sql.NullTime   `json:"donation_date"`
	TotalReads      int            `json:"total_reads"`
	AverageRating   float64        `json:"average_rating"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type BookStatus string

const (
	StatusAvailable BookStatus = "available"
	StatusReading   BookStatus = "reading"
	StatusReserved  BookStatus = "reserved"
	StatusRequested BookStatus = "requested"
)

type ReadingHistory struct {
	ID           string        `json:"id"`
	BookID       string        `json:"book_id"`
	Book         *Book         `json:"book,omitempty"`
	ReaderID     string        `json:"reader_id"`
	Reader       *User         `json:"reader,omitempty"`
	StartDate    time.Time     `json:"start_date"`
	EndDate      sql.NullTime  `json:"end_date"`
	DurationDays sql.NullInt64 `json:"duration_days"`
	Notes        string        `json:"notes"`
	Rating       sql.NullInt64 `json:"rating"`
	Review       string        `json:"review"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type WaitingQueue struct {
	ID       string    `json:"id"`
	BookID   string    `json:"book_id"`
	Book     *Book     `json:"book,omitempty"`
	UserID   string    `json:"user_id"`
	User     *User     `json:"user,omitempty"`
	Position int       `json:"position"`
	JoinedAt time.Time `json:"joined_at"`
	Notified bool      `json:"notified"`
}

type BookRequest struct {
	ID                 string          `json:"id"`
	BookID             string          `json:"book_id"`
	Book               *Book           `json:"book,omitempty"`
	UserID             string          `json:"user_id"`
	User               *User           `json:"user,omitempty"`
	Status             string          `json:"status"`
	PriorityScore      float64         `json:"priority_score"`
	InterestMatchScore float64         `json:"interest_match_score"`
	DistanceKm         sql.NullFloat64 `json:"distance_km"`
	RequestedAt        time.Time       `json:"requested_at"`
	ProcessedAt        sql.NullTime    `json:"processed_at"`
	DueDate            sql.NullTime    `json:"due_date"`
}

type ReadingIdea struct {
	ID        string    `json:"id"`
	BookID    string    `json:"book_id"`
	Book      *Book     `json:"book,omitempty"`
	UserID    string    `json:"user_id"`
	User      *User     `json:"user,omitempty"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Upvotes   int       `json:"upvotes"`
	Downvotes int       `json:"downvotes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IdeaVote struct {
	ID        string    `json:"id"`
	IdeaID    string    `json:"idea_id"`
	UserID    string    `json:"user_id"`
	VoteType  string    `json:"vote_type"`
	CreatedAt time.Time `json:"created_at"`
}

type UserReview struct {
	ID                  string         `json:"id"`
	ReviewerID          string         `json:"reviewer_id"`
	Reviewer            *User          `json:"reviewer,omitempty"`
	RevieweeID          string         `json:"reviewee_id"`
	Reviewee            *User          `json:"reviewee,omitempty"`
	BookID              sql.NullString `json:"book_id"`
	BehaviorRating      sql.NullInt64  `json:"behavior_rating"`
	BookConditionRating sql.NullInt64  `json:"book_condition_rating"`
	CommunicationRating sql.NullInt64  `json:"communication_rating"`
	Comment             string         `json:"comment"`
	CreatedAt           time.Time      `json:"created_at"`
}

type Donation struct {
	ID           string          `json:"id"`
	DonorID      string          `json:"donor_id"`
	Donor        *User           `json:"donor,omitempty"`
	DonationType string          `json:"donation_type"`
	BookID       sql.NullString  `json:"book_id"`
	Book         *Book           `json:"book,omitempty"`
	Amount       sql.NullFloat64 `json:"amount"`
	Currency     string          `json:"currency"`
	Message      string          `json:"message"`
	IsPublic     bool            `json:"is_public"`
	CreatedAt    time.Time       `json:"created_at"`
}

type UserInterest struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Interest  string    `json:"interest"`
	Weight    float64   `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}

type UserBookmark struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	BookID        string    `json:"book_id"`
	Book          *Book     `json:"book,omitempty"`
	BookmarkType  string    `json:"bookmark_type"`
	PriorityLevel int       `json:"priority_level"`
	CreatedAt     time.Time `json:"created_at"`
}

type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Link      string    `json:"link"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type SuccessScoreHistory struct {
	ID            string         `json:"id"`
	UserID        string         `json:"user_id"`
	ChangeAmount  int            `json:"change_amount"`
	Reason        string         `json:"reason"`
	ReferenceType string         `json:"reference_type"`
	ReferenceID   sql.NullString `json:"reference_id"`
	CreatedAt     time.Time      `json:"created_at"`
}
