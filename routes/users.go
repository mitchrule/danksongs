package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

		w.Write(payload)
	}

	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = actions.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
