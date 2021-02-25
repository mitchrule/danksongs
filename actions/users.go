package actions

import (
	"log"

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
