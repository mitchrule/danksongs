package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
)

// CreateUserHandler manages the http request
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check to see if the request is the correct format
	if r.Header.Values("content-type")[0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		res := ErrorResponse{
			Code:    400,
			Message: "Incorrect content-type",
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Writing error payload..")
		w.Write(payload)
	}

	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("--Internal error in unmarshalling--")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = actions.CreateUser(user)
	if err != nil {
		log.Println("--Internal error in action--")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err == nil {
		w.WriteHeader(http.StatusCreated)
	}

}

// LoginUserHandler manages the http request
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check to see if the request is the correct format
	if r.Header.Values("content-type")[0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		res := ErrorResponse{
			Code:    400,
			Message: "Incorrect content-type",
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(payload)
	}

	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	jwtToken, err := actions.LoginUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err == nil {

		bearerToken := "Bearer " + jwtToken

		// Sets auth header in the response
		w.Header().Add("Authorization", bearerToken)

		// Set the JWT token as a cookie in the user browser incase the header is not consistant
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   jwtToken,
			Expires: time.Now().Add(actions.SESSION_MINS),
		})

		w.WriteHeader(http.StatusCreated)
	}
}

// Middleware for JWT authentication
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenString string
		var bearerToken string

		// Catchall while I figure out how to use the authorisation header exclusively
		if r.Header.Get("Authorization") != "" {
			// Get the token from the auth header
			bearerToken := r.Header.Get("Authorization")
			tokenParts := strings.Split(bearerToken, " ")
			tokenString = tokenParts[1]
			log.Println("Token Retrieved From Header...")
		} else {
			// Get the token from the cookie supplied
			cookie, err := r.Cookie("token")
			tokenString = cookie.Value
			//log.Println("Cookie:", bearerToken)
			if err != nil {
				if err == http.ErrNoCookie {
					// If the cookie is not set, return an unauthorized status
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				// For any other type of error, return a bad request status
				w.WriteHeader(http.StatusBadRequest)
				log.Panicln(err)
				return
			}
			log.Println("Token Retrieved From Cookie...")
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
				Expires: time.Now().Add(actions.SESSION_MINS),
			})
			log.Println("new token set...")
			next.ServeHTTP(w, r)
		} else {
			log.Println(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
