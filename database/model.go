package database

import "time"

type UserModel struct {
	ID        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt any       `json:"updated_at"`
}

type DomainModel struct {
	ID        int       `json:"id,omitempty"`
	user      int       `json:"user" binding:"required"`
	Domain    string    `json:"domain" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt any       `json:"updated_at"`
}
type DNSRECORDMODEL struct {
	ID        int       `json:"id,omitempty"`
	Domain    int       `json:"domain" binding:"required"`
	Record    string    `json:"record" binding:"required"`
	Type      string    `json:"type" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt any       `json:"updated_at"`
}
