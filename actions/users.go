package actions

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var SESSION_MINS = time.Duration(5) * time.Minute

// CreateUser adds a user to the database
func CreateUser(user models.User) error {
	log.Println("Creating user...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Hash the users password for storage
	user.Password = getHash([]byte(user.Password))

	// Check the username has not already been used
	err := database.UsersCollection.FindOne(ctx, bson.M{"name": user.Name}).Err()
	if err != mongo.ErrNoDocuments {
		log.Println("Duplicate User...")
		return errors.New("Duplicate user")
	}

	// Place in db otherwise
	insertResult, err := database.UsersCollection.InsertOne(ctx, user)

	if err != nil {
		log.Println("user insert error")
		log.Println(err)
		return err
	}
	log.Println("Inserted: ", user)
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

	// Generate the claims for the JWT token for this session
	expirationTime := time.Now().Add(SESSION_MINS)

	// Create a claim based on user info
	claim := models.Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Place in db otherwise
	insertResult, err := database.UsersCollection.InsertOne(ctx, claim)
	if err != nil {
		log.Println("token insert error")
		log.Println(err)
		return "", err
	}

	//log.Println(claim)
	log.Println("Inserted: ", insertResult)
	jwtToken, err := GenerateJWT(claim)

	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

////////////////////////////////////////////////////////////////////
///////////////////// Helper functions /////////////////////////////
////////////////////////////////////////////////////////////////////

// getHash generates a hash from a given password string
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
