package main

import (
	// routes "Auth-API/routes"
	db "dns-user/database"
	"dns-user/routes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var startTime time.Time

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// if os.Getenv("GIN_MODE") != gin.ReleaseMode {
	// 	// os.Setenv("GIN_MODE", gin.ReleaseMode)
	// }
	// gin.SetMode(gin.ReleaseMode)

	startTime = time.Now()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// connStr := os.Getenv("DB_CONN_STR")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db.DatabaseConnection(connStr)

	router := gin.New()
	router.Use(gin.Logger())

	apiVersion1Group := router.Group("/api/v1")

	routes.AuthRoutes(apiVersion1Group)
	routes.UserRoutes(apiVersion1Group)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
	uptime := time.Since(startTime)

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2",
			"message": "This is a message from api-2",
			"status":  "success",
			"code":    200,
			"error":   nil,
			"data":    nil,
			"uptime": gin.H{
				"seconds": uptime.Seconds(),
				"minutes": uptime.Minutes(),
				"hours":   uptime.Hours(),
			},
		})
	})
	router.Run(":" + port)
}
