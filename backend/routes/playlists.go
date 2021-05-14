package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		return
	}

	// Changed it to a json
	var playListData models.PlaylistData

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &playListData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playListID, err := actions.CreatePlaylist(playListData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playListIDJSON, err := primitive.ObjectID.MarshalJSON(playListID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
	}

	var playListID primitive.ObjectID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &playListID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playList, err := actions.GetPlaylist(playListID)
	if err != nil {

		log.Println(err)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	playListByte, err := json.Marshal(playList)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		// Returns the playlist as a JSON
		w.WriteHeader(http.StatusOK)
		w.Write(playListByte)
	}
}

// TODO
func GetRecentPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
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

	playLists, err := actions.GetRecentPlaylists()
	if err != nil {

		log.Println(err)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	playListByte, err := json.Marshal(playLists)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		// Returns the playlist as a JSON
		w.WriteHeader(http.StatusOK)
		w.Write(playListByte)
	}
}

func SearchPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	var searchTerm string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &searchTerm)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playLists, err := actions.SearchPlaylists(searchTerm)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	playListByte, err := json.Marshal(playLists)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		// Returns the playlist as a JSON
		w.WriteHeader(http.StatusOK)
		w.Write(playListByte)
	}
}

func DeletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Println("-----Error In Reading-----")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &playListID)
	if err != nil {
		log.Println("-----Error In Unmarshelling-----")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success, err := actions.DeletePlaylist(playListID)
	if err != nil || !success {
		log.Println("-----Error In Action-----")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Returns nothing other than the success statement
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Playlist Deleted"))

}

func AddSongHandler(w http.ResponseWriter, r *http.Request) {

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
		return
	}

	params := mux.Vars(r)
	playlistID, err := primitive.ObjectIDFromHex(params["playlistid"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var song models.Song

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &song)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success, err := actions.AddSong(playlistID, song)

	if err != nil || !success {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Returns nothing other than the success statement
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Song Addded"))
}

func RemoveSongHandler(w http.ResponseWriter, r *http.Request) {

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
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ids models.SongPLPair

	err = json.Unmarshal(body, &ids)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//playList, err := actions.GetPlaylist(playListID)
	success, err := actions.RemoveSong(ids)
	if err != nil || !success {

		log.Println(err)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Song Removed"))
}
