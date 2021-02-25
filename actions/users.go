package actions

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchrule/danksongs/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser adds a user to the database
func CreateUser(user models.User) error {
	return nil
}

/// Helper functions

// getHash generates a hash from a given password string
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// GenerateJWT generates a JWT token for a particuar session
func GenerateJWT() (string, error) {

	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
