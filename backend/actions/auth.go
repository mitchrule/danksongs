package actions

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchrule/danksongs/common"
	"github.com/mitchrule/danksongs/models"
)

// Validate user token checks their JTW token is valid before acessing the API and retured an
// updated token
func ValidateUserJWT(tknStr string) (string, error) {
	// We can obtain the session token from the requests cookies, which come with every request

	// Initialize a new instance of `Claims`
	claim := models.Claims{}

	// Parse the JWT string and store the result in `claims`.
	tkn, err := jwt.ParseWithClaims(tknStr, &claim, func(token *jwt.Token) (interface{}, error) {
		// I dont like that for this to work we have to return our secret key
		// Will fix if I can
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		log.Println("Error with jwt signature...")
		if err == jwt.ErrSignatureInvalid {
			//return
			log.Println("Signature Invalid")
			return "", err
		}
		//return
		return "", err
	}
	if !tkn.Valid {
		log.Println("Token Invalid")
		return "", err
	}

	// Create a new token with the claims gathered from the previous token
	expirationTime := time.Now().Add(common.SESSION_MINS)
	claim.ExpiresAt = expirationTime.Unix()
	newToken, err := GenerateJWT(claim)

	if err != nil {
		return "", err
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	return newToken, nil
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
