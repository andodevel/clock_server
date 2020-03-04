package helpers

import (
	"crypto/rand"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/andodevel/clock_server/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	// JWTSecrect ...
	JWTSecrect = "serEct" // TODO: Move to config
	// JWTCookieKey ...
	JWTCookieKey = "JWTCookie"
)

// JWTClaims ...
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateJWTToken ...
func CreateJWTToken(user *models.User) (string, error) {
	claims := JWTClaims{
		user.Username,
		jwt.StandardClaims{
			Id:        strconv.Itoa(user.ID),
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(JWTSecrect))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseJWTToken ...
func ParseJWTToken(token string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	rawToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecrect), nil
	})
	if rawToken == nil && err != nil {
		return nil, err
	}
	err = claims.Valid()
	if err != nil {
		log.Println("Invalid JWT claims, error: " + err.Error())
		return nil, err
	}

	return claims, nil
}

// IsValidJWTToken ...
func IsValidJWTToken(token string) bool {
	claims := &JWTClaims{}

	rawToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecrect), nil
	})
	if (rawToken == nil && err != nil) || claims.Valid() != nil {
		return false
	}

	return true
}

// GenerateAccessToken ...
func GenerateAccessToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
