package models

import (
	"github.com/nad279444/csrf_token-api/randomstrings"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Username, PasswordHash, Role string
}

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15

//encode into string and send it
func GenerateCSRFSecret() (string, error) {
	return randomstrings.GenerateRandomString(32)
}