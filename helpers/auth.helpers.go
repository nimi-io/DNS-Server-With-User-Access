package helpers

import (
	"dns-user/database"
    "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

var secretKey = []byte("secret-key")

func GenerateToken(user database.UserModel) (string, error){
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": user.ID,
        "email": user.Email,
        "username": user.Username,
    })
    tokenString, err := token.SignedString(secretKey)
    return tokenString, err
}