package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Structs used within MongoDB

// User - Models a user and their associated credentials
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password"`
}

// A Vote With the user assoicated with it
type Vote struct {
	VoterID primitive.ObjectID `bson:"_voterid,omitempty"`
}

// Claims - for the JWT token verification
type Claims struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username"`
	jwt.StandardClaims
}

// Song - Models a song
type Song struct {
	ID     primitive.ObjectID `bson:"_id"`
	Title  string             `bson:"title,omitempty"`
	Artist string             `bson:"artist,omitempty"`
	URL    string             `bson:"url,omitempty"`
	Votes  []Vote             `bson:"votes"`
}

// Playlist - Models a list of songs to be voted on
type Playlist struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`

	// TODO: Change the model into something that can be returned in a json
	Songs          []Song  `bson:"songs"`
	VoteThreshold  uint16  `bson:"votethreshold,omitempty"`
	VotePreportion float64 `bson:"votepreportion,omitempty"`
}

// Structs used to capture data from requests

// PlaylistData - Captures data for create playlist
type PlaylistData struct {
	PlayListName   string
	VoteThreshold  uint16
	VotePreportion float64
}

// SongPLPair - A song ID and an associated Playlist ID for Add/Remove Song to work
// with
type SongPLPair struct {
	SongID     primitive.ObjectID
	PlaylistID primitive.ObjectID
}
