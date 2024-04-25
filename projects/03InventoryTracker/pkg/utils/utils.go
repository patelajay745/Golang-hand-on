package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/models"
)

var secretKey = []byte("Secret")

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// VerifyToken verifies the JWT token and returns the claims if valid.
func VerifyToken(tokenString string) (*models.CustomClaims, error) {
	claims := &models.CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ValidationError{}
	}

	return claims, nil
}

func GenerateToken(userId int, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return tokenString, nil
}
