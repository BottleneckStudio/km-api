package auth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

type JWTParser struct {
	keySet *jwk.Set
}

func New(keySet *jwk.Set) *JWTParser {
	if keySet == nil {
		return nil
	}
	return &JWTParser{
		keySet: keySet,
	}
}

// Parse parses a token in string form to an actual JWT Token
func (p *JWTParser) Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		keys := p.keySet.LookupKeyID(kid)
		if len(keys) == 0 {
			return nil, fmt.Errorf("key %v not found", kid)
		}
		return keys[0].Materialize()
	})

	return token, err
}
