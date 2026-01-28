package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
	"github.com/yourusername/online-library/internal/models"
	"github.com/yourusername/online-library/internal/repository"
	"github.com/yourusername/online-library/internal/services"
)

type IdeaHandler struct {
	ideaRepo     *repository.IdeaRepository
	scoreService *services.SuccessScoreService
	notifService *services.NotificationService
}

func NewIdeaHandler(ideaRepo *repository.IdeaRepository, scoreService *services.SuccessScoreService, notifService *services.NotificationService) *IdeaHandler {
	return &IdeaHandler{
		ideaRepo:     ideaRepo,
		scoreService: scoreService,
		notifService: notifService,
	}
}

func (h *IdeaHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var req dto.CreateIdeaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	idea := &models.ReadingIdea{
		BookID:  req.BookID,
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.ideaRepo.Create(idea); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create idea"})
		return
	}

	// Update success score
	h.scoreService.ProcessIdeaPosted(userID, idea.ID)

	c.JSON(http.StatusCreated, idea)
}

func (h *IdeaHandler) GetByBook(c *gin.Context) {
	bookID := c.Param("bookId")
	ideas, err := h.ideaRepo.GetByBook(bookID, 50, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to fetch ideas"})
		return
	}
	c.JSON(http.StatusOK, ideas)
}

func (h *IdeaHandler) Vote(c *gin.Context) {
	userID := c.GetString("user_id")
	ideaID := c.Param("id")
	var req dto.VoteIdeaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.ideaRepo.Vote(ideaID, userID, req.VoteType); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to vote"})
		return
	}

	// Get idea author
	idea, _ := h.ideaRepo.GetByID(ideaID)
	if idea != nil && idea.UserID != userID {
		if req.VoteType == "upvote" {
			h.scoreService.ProcessIdeaUpvote(idea.UserID, ideaID)
		} else {
			h.scoreService.ProcessIdeaDownvote(idea.UserID, ideaID)
		}
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Vote recorded", nil))
}
