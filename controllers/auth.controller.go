package controllers

import (
	"context"
	model "dns-user/database"
	help "dns-user/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user model.UserModel

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
		log.Println(user)

	hashedPassword, err := help.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	checkQuery := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err = model.DB.QueryRowContext(ctx, checkQuery, user.Email).Scan(&count)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}
	

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	err = model.DB.QueryRowContext(ctx, query, user.Username, user.Email, hashedPassword).Scan(&user.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
