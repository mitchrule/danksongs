package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/actions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit CreatePlaylistHandler"))
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

		// Returns the associated objectID
		w.Write(payload)
	}

	var playListName string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &playListName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListID, err := actions.CreatePlaylist(playListName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListIDJSON, err := primitive.ObjectID.MarshalJSON(playListID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns the newly created playList ObjectID
	w.Write(playListIDJSON)
	w.WriteHeader(http.StatusCreated)

}

func GetPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit GetPlaylistHandler"))

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
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &playListID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	playList, err := actions.GetPlaylist(playListID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListByte, err := json.Marshal(playList)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns the playlist as a JSON
	w.Write(playListByte)
	w.WriteHeader(http.StatusCreated)

}

func DeletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit DeletePlaylistHandler"))

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
		w.WriteHeader(http.StatusBadRequest)
	}

	//primitive.ObjectID.JSON.Unmarshal(body, &playListID)

	err = json.Unmarshal(body, &playListID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	playList, err := actions.GetPlaylist(playListID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	playListByte, err := json.Marshal(playList)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns nothing other than the success statement 
	w.Write(playListByte)
	w.WriteHeader(http.StatusCreated)

}

func AddSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit AddSongHandler"))

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

	//var playListID primitive.ObjectID
	//var songID 	   primitive.ObjectID

	var ids := {
		playListID: primitive.ObjectID,
		songID:		primitive.ObjectID,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//playList, err := actions.GetPlaylist(playListID)
	success, err := actions.AddSong(ids.playListID,ids.songID)

	if err != nil || !success {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns nothing other than the success statement 
	w.WriteHeader(http.StatusCreated)
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

	var ids := {
		playListID: primitive.ObjectID,
		songID:		primitive.ObjectID,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//playList, err := actions.GetPlaylist(playListID)
	success, err := actions.RemoveSong(ids.playListID,ids.songID)

	if err != nil || !success {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Returns nothing other than the success statement 
	w.WriteHeader(http.StatusCreated)
}
