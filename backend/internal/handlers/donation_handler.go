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

type DonationHandler struct {
	donationRepo *repository.DonationRepository
	scoreService *services.SuccessScoreService
	db           *sql.DB
}

func NewDonationHandler(donationRepo *repository.DonationRepository, scoreService *services.SuccessScoreService, db *sql.DB) *DonationHandler {
	return &DonationHandler{
		donationRepo: donationRepo,
		scoreService: scoreService,
		db:           db,
	}
}

func (h *DonationHandler) Create(c *gin.Context) {
	donorID := c.GetString("user_id")
	var req dto.CreateDonationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	donation := &models.Donation{
		DonorID:      donorID,
		DonationType: req.DonationType,
		Currency:     req.Currency,
		Message:      req.Message,
		IsPublic:     req.IsPublic,
	}

	if req.BookID != nil {
		donation.BookID = sql.NullString{String: *req.BookID, Valid: true}
	}
	if req.Amount != nil {
		donation.Amount = sql.NullFloat64{Float64: *req.Amount, Valid: true}
	}

	if err := h.donationRepo.Create(donation); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to record donation"})
		return
	}

	// Update user as donor
	h.db.Exec("UPDATE users SET is_donor = true WHERE id = $1", donorID)

	// Update success score
	if req.DonationType == "book" {
		h.scoreService.ProcessBookDonation(donorID, donation.ID)
	} else {
		h.scoreService.ProcessMoneyDonation(donorID, donation.ID)
	}

	c.JSON(http.StatusCreated, donation)
}

func (h *DonationHandler) GetAll(c *gin.Context) {
	donations, err := h.donationRepo.GetAll(100, 0, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to fetch donations"})
		return
	}
	c.JSON(http.StatusOK, donations)
}
