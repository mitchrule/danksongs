package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Values("content-type")[0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		res := ErrorResponse{
			Code:    400,
			Message: "Incorrect content-type",
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Println("error...")
			log.Fatal(err)
		}

		// Returns the associated objectID
		w.Write(payload)
	}

	var playListName string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &playListName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListID, err := actions.CreatePlaylist(playListName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListIDJSON, err := primitive.ObjectID.MarshalJSON(playListID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns the newly created playList ObjectID
	w.WriteHeader(http.StatusCreated)
	w.Write(playListIDJSON)
}

func GetPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hit GetPlaylistHandler"))

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

	var playListID primitive.ObjectID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &playListID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	playList, err := actions.GetPlaylist(playListID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListByte, err := json.Marshal(playList)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns the playlist as a JSON
	w.Write(playListByte)
	w.WriteHeader(http.StatusCreated)

}

func DeletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hit DeletePlaylistHandler"))

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

	var playListID primitive.ObjectID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &playListID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	success, err := actions.DeletePlaylist(playListID)
	if err != nil || !success {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns nothing other than the success statement
	w.WriteHeader(http.StatusOK)
}

func AddSongHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hit AddSongHandler"))

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

	var ids models.SongPLPair

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &ids)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	success, err := actions.AddSong(ids)

	if err != nil || !success {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns nothing other than the success statement
	w.WriteHeader(http.StatusOK)
}

func RemoveSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit RemoveSongHandler"))

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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var ids models.SongPLPair

	err = json.Unmarshal(body, &ids)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	//playList, err := actions.GetPlaylist(playListID)
	success, err := actions.RemoveSong(ids)

	if err != nil || !success {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns nothing other than the success statement
	w.WriteHeader(http.StatusOK)
}
