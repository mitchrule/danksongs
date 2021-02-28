package actions

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser adds a user to the database
func CreateUser(user models.User) error {
	log.Println("Creating user...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Hash the users password for storage
	user.Password = getHash([]byte(user.Password))

	insertResult, err := database.UsersCollection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	log.Println("Inserted: ", insertResult)
	return nil
}

//LoginUser loggs in the user provided they provide the correct details
func LoginUser(user models.User) (string, error) {
	log.Println("Logging in user...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the details of the user from db
	var dbUser models.User
	err := database.UsersCollection.FindOne(ctx, bson.M{"name": user.Name}).Decode(&dbUser)

	if err != nil {
		return "", err
	}

	// Compare passwords
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		return "", passErr
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

///////////////////////////////////////////////////
/// Helper functions
///////////////////////////////////////////////////

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
