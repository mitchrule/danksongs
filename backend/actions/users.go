package actions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchrule/danksongs/common"
	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

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
		return errors.New("duplicate user")
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
	expirationTime := time.Now().Add(common.SESSION_MINS)

	// Create a claim based on user info
	claim := models.Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Place in db otherwise
	insertResult, err := database.JWTCollection.InsertOne(ctx, claim)
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

// LogoutUser takes a userID
func LogoutUser(tknStr string) (bool, error) {
	// TODO if other backend work is required
	return false, nil
}

// DeleteUser removes a user from the Database based on
// the userID provided and logs the user out in the process
func DeleteUser(tknStr string) (bool, error) {

	// Locate the users details within the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Take the JWT token provided by the user and parse it to retrieve the associated user
	// info
	user, err := GetUserFromToken(tknStr)

	if err != nil {
		return false, err
	}

	// @TODO find and delete the associated JWT claim
	// Retrieve the details of the user from db based on their name
	var dbUser models.User
	err = database.UsersCollection.FindOneAndDelete(ctx, bson.M{"name": user.Name}).Decode(&dbUser)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ChangePassword(user models.NewUser) (bool, error) {
	log.Println("Chainging Password...")

	if user.Password1 != user.Password2 {
		return false, fmt.Errorf("Error: Passwords do not match")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var oldUser models.User

	// Generate new user entry with changed password and store in db
	var userWithNewPWord models.User

	userWithNewPWord.ID = user.ID
	userWithNewPWord.Name = user.Name
	userWithNewPWord.Password = user.Password1

	err := database.UsersCollection.FindOneAndReplace(ctx, bson.M{"name": user.Name}, userWithNewPWord).Decode(&oldUser)

	log.Println("Changed Pasword for: ", oldUser.Name)

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetUserFromToken extracts the username from a JWT token
func GetUserFromToken(tokenString string) (models.User, error) {
	var username string

	// decode the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = fmt.Sprintf("%v", claims["username"])
	} else {
		log.Println(err)
		return models.User{}, err
	}

	// get user by username from db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err = database.UsersCollection.FindOne(ctx, bson.D{{"name", username}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
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
