package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/usermanagementapi/handler"
	"github.com/joshua468/usermanagementapi/middleware"
	"github.com/joshua468/usermanagementapi/model"
	"github.com/joshua468/usermanagementapi/repository"
	"github.com/joshua468/usermanagementapi/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Initialize repository
	var userRepository repository.UserRepository
	userRepository = *repository.NewUserRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepository)

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)

	// Setup Gin router
	r := gin.Default()

	// Use logging middleware
	r.Use(middleware.LoggingMiddleware())

	// Define routes
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.POST("/", userHandler.CreateUser)
	}

	// Auto migrate the User model
	err = model.Migrate(db)
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	// Start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error:", err)
	}
}
