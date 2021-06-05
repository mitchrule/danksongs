package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/common"
)

// Middleware for JWT authentication
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenString string
		var bearerToken string
		cookie, err := r.Cookie("token")

		// Catchall while I figure out how to use the authorisation header exclusively
		if r.Header.Get("Authorization") != "" {
			// Get the token from the auth header
			bearerToken := r.Header.Get("Authorization")
			tokenParts := strings.Split(bearerToken, " ")
			tokenString = tokenParts[1]
			log.Println("Token Retrieved From Header...")
		} else if err != http.ErrNoCookie {
			// Get the token from the cookie supplied
			log.Println("Cookie:", cookie)
			tokenString = cookie.Value

			if err != nil {
				// For any other type of error, return a bad request status
				w.WriteHeader(http.StatusBadRequest)
				log.Println(err)
				return
			}
			log.Println("Token Retrieved From Cookie...")
		} else {
			// No token so refuse request
			log.Println("No token...")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if tokenString == "" {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Token string not found")
		}

		// Validate the token is from the user specified
		newToken, err := actions.ValidateUserJWT(tokenString)

		// Continue the http request to the api if this succeeds
		if err == nil && newToken != "" {

			// Set the auth header and cookie to the new JWT token here
			// Sets auth header in the response
			w.Header().Add("Authorization", bearerToken)

			// Set the JWT token as a cookie in the user browser incase the header is not consistant
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   newToken,
				Expires: time.Now().Add(common.SESSION_MINS),
			})
			log.Println("new token set...")
			next.ServeHTTP(w, r)
		} else {
			log.Println(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
