package actions

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	expirationTime := time.Now().Add(5 * time.Minute)

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

	log.Println("Inserted: ", insertResult)
	jwtToken, err := GenerateJWT(claim)

	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

// Validate user token checks their JTW token is valid before acessing the API
func ValidateUserToken(tknStr string) (bool, error) {
	// We can obtain the session token from the requests cookies, which come with every request

	// Initialize a new instance of `Claims`
	claim := &models.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claim, func(token *jwt.Token) (interface{}, error) {
		// I dont like that for this to work we have to return our secret key
		// Will fix if I can
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	log.Panicln(claim)

	if err != nil {
		log.Println("Error...")
		if err == jwt.ErrSignatureInvalid {
			//return
			log.Println("Signature Invalid")
			return false, err
		}
		//return
		return false, err
	}
	if !tkn.Valid {
		log.Println("Token Invalid")
		return false, err
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token

	return true, nil

	//w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
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
// Delete if other method is better

func GenerateJWT(claim models.Claims) (string, error) {

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(secretKey)

	log.Println(tokenString)
	if err != nil {
		log.Println("Error in JWT token generation")
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
