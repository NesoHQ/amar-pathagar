package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/dto"
	"github.com/yourusername/online-library/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	response, err := h.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse("User registered successfully", response))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	response, err := h.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Login successful", response))
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")
	c.JSON(http.StatusOK, dto.SuccessResponse("User info", gin.H{
		"user_id": userID,
	}))
}
