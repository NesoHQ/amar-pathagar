package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
	"github.com/yourusername/online-library/internal/models"
	"github.com/yourusername/online-library/internal/repository"
)

type BookmarkHandler struct {
	bookmarkRepo *repository.BookmarkRepository
}

func NewBookmarkHandler(bookmarkRepo *repository.BookmarkRepository) *BookmarkHandler {
	return &BookmarkHandler{bookmarkRepo: bookmarkRepo}
}

func (h *BookmarkHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var req dto.CreateBookmarkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	bookmark := &models.UserBookmark{
		UserID:        userID,
		BookID:        req.BookID,
		BookmarkType:  req.BookmarkType,
		PriorityLevel: req.PriorityLevel,
	}

	if err := h.bookmarkRepo.Create(bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create bookmark"})
		return
	}

	c.JSON(http.StatusCreated, bookmark)
}

func (h *BookmarkHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	bookID := c.Param("bookId")
	bookmarkType := c.Query("type")

	if err := h.bookmarkRepo.Delete(userID, bookID, bookmarkType); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete bookmark"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Bookmark removed", nil))
}

func (h *BookmarkHandler) GetByUser(c *gin.Context) {
	userID := c.GetString("user_id")
	bookmarkType := c.Query("type")

	bookmarks, err := h.bookmarkRepo.GetByUser(userID, bookmarkType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to fetch bookmarks"})
		return
	}

	c.JSON(http.StatusOK, bookmarks)
}
