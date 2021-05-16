package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Handles logout request
func LogoutUserHandler(w http.ResponseWriter, r *http.Request) {

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

	// @TODO Fetch the JWT token from the user cookie
	// instead of passing it through as a function
	var userID primitive.ObjectID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	err = json.Unmarshal(body, &userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	success, err := actions.LogoutUser(userID)
	if err != nil || !success {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// Remove the users token cookie

		// Remove the users authorisation header

		// Direct them to the homepage

	}
	// w.WriteHeader(http.StatusExpectationFailed)
	// w.Write([]byte("Hit LogoutUserHandler"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusExpectationFailed)
	w.Write([]byte("Hit DeleteUserHandler"))
}
