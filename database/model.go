package database

import "time"

type UserModel struct {
    ID       int    `json:"id,omitempty"`
    Username string `json:"username"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
