package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// Parser
type Parser interface {
	Parse(tokenString string) (*jwt.Token, error)
}
