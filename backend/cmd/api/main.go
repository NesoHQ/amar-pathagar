package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-library/internal/config"
	"github.com/yourusername/online-library/internal/database"
	"github.com/yourusername/online-library/internal/handlers"
	"github.com/yourusername/online-library/internal/middleware"
	"github.com/yourusername/online-library/internal/repository"
	"github.com/yourusername/online-library/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	db, err := database.Connect(cfg.Database.ConnectionString())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.DB)
	ideaRepo := repository.NewIdeaRepository(db.DB)
	donationRepo := repository.NewDonationRepository(db.DB)
	reviewRepo := repository.NewReviewRepository(db.DB)
	bookmarkRepo := repository.NewBookmarkRepository(db.DB)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret)
	successScoreService := services.NewSuccessScoreService(db.DB)
	notificationService := services.NewNotificationService(db.DB)
	// matchingService := services.NewMatchingService(db.DB) // Available for future use

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(db.DB)
	ideaHandler := handlers.NewIdeaHandler(ideaRepo, successScoreService, notificationService)
	donationHandler := handlers.NewDonationHandler(donationRepo, successScoreService, db.DB)
	reviewHandler := handlers.NewReviewHandler(reviewRepo, successScoreService, notificationService)
	bookmarkHandler := handlers.NewBookmarkHandler(bookmarkRepo)

	// Setup router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Public routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(authService))
	{
		api.GET("/me", authHandler.Me)

		// User routes
		api.GET("/users/:id/profile", userHandler.GetPublicProfile)
		api.GET("/users/:id/reviews", reviewHandler.GetByUser)
		api.PUT("/users/profile", userHandler.UpdateProfile)
		api.POST("/users/interests", userHandler.AddInterests)
		api.GET("/leaderboard", userHandler.GetLeaderboard)

		// Reading ideas routes
		api.POST("/ideas", ideaHandler.Create)
		api.GET("/books/:bookId/ideas", ideaHandler.GetByBook)
		api.POST("/ideas/:id/vote", ideaHandler.Vote)

		// Review routes
		api.POST("/reviews", reviewHandler.Create)

		// Donation routes
		api.POST("/donations", donationHandler.Create)
		api.GET("/donations", donationHandler.GetAll)

		// Bookmark routes
		api.POST("/bookmarks", bookmarkHandler.Create)
		api.DELETE("/bookmarks/:bookId", bookmarkHandler.Delete)
		api.GET("/bookmarks", bookmarkHandler.GetByUser)
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
