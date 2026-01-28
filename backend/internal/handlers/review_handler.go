package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
	"github.com/yourusername/online-library/internal/models"
	"github.com/yourusername/online-library/internal/repository"
	"github.com/yourusername/online-library/internal/services"
)

type ReviewHandler struct {
	reviewRepo   *repository.ReviewRepository
	scoreService *services.SuccessScoreService
	notifService *services.NotificationService
}

func NewReviewHandler(reviewRepo *repository.ReviewRepository, scoreService *services.SuccessScoreService, notifService *services.NotificationService) *ReviewHandler {
	return &ReviewHandler{
		reviewRepo:   reviewRepo,
		scoreService: scoreService,
		notifService: notifService,
	}
}

func (h *ReviewHandler) Create(c *gin.Context) {
	reviewerID := c.GetString("user_id")
	var req dto.CreateUserReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	review := &models.UserReview{
		ReviewerID: reviewerID,
		RevieweeID: req.RevieweeID,
		Comment:    req.Comment,
	}

	if req.BookID != "" {
		review.BookID = sql.NullString{String: req.BookID, Valid: true}
	}
	if req.BehaviorRating != nil {
		review.BehaviorRating = sql.NullInt64{Int64: int64(*req.BehaviorRating), Valid: true}
	}
	if req.BookConditionRating != nil {
		review.BookConditionRating = sql.NullInt64{Int64: int64(*req.BookConditionRating), Valid: true}
	}
	if req.CommunicationRating != nil {
		review.CommunicationRating = sql.NullInt64{Int64: int64(*req.CommunicationRating), Valid: true}
	}

	if err := h.reviewRepo.Create(review); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create review"})
		return
	}

	// Calculate average and update success score
	avgRating := 0.0
	count := 0
	if req.BehaviorRating != nil {
		avgRating += float64(*req.BehaviorRating)
		count++
	}
	if req.BookConditionRating != nil {
		avgRating += float64(*req.BookConditionRating)
		count++
	}
	if req.CommunicationRating != nil {
		avgRating += float64(*req.CommunicationRating)
		count++
	}
	if count > 0 {
		avgRating /= float64(count)
		if avgRating >= 4.0 {
			h.scoreService.ProcessPositiveReview(req.RevieweeID, review.ID)
		} else if avgRating < 3.0 {
			h.scoreService.ProcessNegativeReview(req.RevieweeID, review.ID)
		}
	}

	c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) GetByUser(c *gin.Context) {
	userID := c.Param("id")
	reviews, err := h.reviewRepo.GetByReviewee(userID, 50, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}
