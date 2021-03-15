package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"token":"` + jwtToken + `"}`))
	}
}

// Middleware for JWT authentication
func AuthenticateJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Returive the cookie from the http request
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Extract the token from the http cookie
		tokenString := c.Value
		log.Println(tokenString)

		// Split the token string into two parts
		tokenParts := strings.Split(tokenString, " ")

		log.Panicln("Current Token Parts: ", tokenParts)
		// #TODO: Use token parts in the action instead of the routes

		if tokenString == "" {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Token string not found")
		}

		// Validate the token is from the user specified
		success, err := actions.ValidateUserToken(tokenString)

		// Continue the http request to the api if this succeeds
		if err == nil && success {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
