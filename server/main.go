package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"terraform-run-dashboard/handlers"
)

func main() {
	// Load .env file (optional — safe to skip if not present)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, continuing without it...")
	}

	router := gin.Default()

	// Enable CORS so frontend (localhost:5173) can access backend (localhost:8080)
	router.Use(cors.Default())

	// Route definitions
	router.GET("/workspaces", handlers.GetWorkspaces)
	router.GET("/workspaces/:id/runs", handlers.GetRuns)
	router.POST("/workspaces/:id/run", handlers.TriggerRun)

	// Determine port to use
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("🚀 Server running on http://localhost:%s\n", port)
	router.Run(":" + port)
}


