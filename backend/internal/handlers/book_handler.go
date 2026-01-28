package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
	"github.com/yourusername/online-library/internal/models"
	"github.com/yourusername/online-library/internal/repository"
)

type BookHandler struct {
	bookRepo *repository.BookRepository
}

func NewBookHandler(bookRepo *repository.BookRepository) *BookHandler {
	return &BookHandler{bookRepo: bookRepo}
}

func (h *BookHandler) GetAll(c *gin.Context) {
	filters := map[string]interface{}{
		"status":   c.Query("status"),
		"category": c.Query("category"),
		"search":   c.Query("search"),
	}

	books, err := h.bookRepo.FindAll(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Books retrieved successfully", books))
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.bookRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Book not found"))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Book retrieved successfully", book))
}

func (h *BookHandler) Create(c *gin.Context) {
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	userID, _ := c.Get("user_id")

	book := &models.Book{
		Title:        req.Title,
		Author:       req.Author,
		ISBN:         req.ISBN,
		CoverURL:     req.CoverURL,
		Description:  req.Description,
		Category:     req.Category,
		Tags:         req.Tags,
		Topics:       req.Topics,
		PhysicalCode: req.PhysicalCode,
		Status:       models.StatusAvailable,
		IsDonated:    req.IsDonated,
	}

	if userID != nil {
		book.CreatedBy = sql.NullString{String: userID.(string), Valid: true}
		if req.IsDonated {
			book.DonatedBy = sql.NullString{String: userID.(string), Valid: true}
		}
	}

	if err := h.bookRepo.Create(book); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse("Book created successfully", book))
}

func (h *BookHandler) Update(c *gin.Context) {
	id := c.Param("id")

	book, err := h.bookRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Book not found"))
		return
	}

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.ISBN != "" {
		book.ISBN = req.ISBN
	}
	if req.CoverURL != "" {
		book.CoverURL = req.CoverURL
	}
	if req.Description != "" {
		book.Description = req.Description
	}
	if req.Category != "" {
		book.Category = req.Category
	}
	if req.Tags != nil {
		book.Tags = req.Tags
	}
	if req.Topics != nil {
		book.Topics = req.Topics
	}
	if req.Status != "" {
		book.Status = models.BookStatus(req.Status)
	}

	if err := h.bookRepo.Update(book); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Book updated successfully", book))
}

func (h *BookHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.bookRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Book deleted successfully", nil))
}
