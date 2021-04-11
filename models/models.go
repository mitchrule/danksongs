package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User - Models a user and their associated credentials
type User struct {
	ID       primitive.ObjectID `bson:"id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password"`
}

// A Vote With the user assoicated with it
type Vote struct {
	ID      primitive.ObjectID `bson:"id,omitempty"`
	VoterID primitive.ObjectID `bson:"name,omitempty"`
	Time    primitive.DateTime `bson:"time,omitempty"`
}

// Claims - for the JWT token verification
type Claims struct {
	ID       primitive.ObjectID `bson:"id,omitempty"`
	Username string             `json:"username"`
	jwt.StandardClaims
}

// Song - Models a song
type Song struct {
	ID     primitive.ObjectID          `bson:"id,omitempty"`
	Title  string                      `bson:"title,omitempty"`
	Artist string                      `bson:"artist,omitempty"`
	URL    string                      `bson:"url,omitempty"`
	Votes  map[primitive.ObjectID]Vote `bson:"votes,omitempty"`
}

// Playlist - Models a list of songs to be voted on
type Playlist struct {
	ID   primitive.ObjectID `bson:"id,omitempty"`
	Name string             `bson:"string,omitempty"`

	// TODO: Change the model into something that can be returned in a json
	Songs         map[primitive.ObjectID]Song `bson:"songs,omitempty"`
	VoteThreshold float64                     `bson:"votethreshold,omitempty"`
}

// SongPLPair - A song ID and an associated Playlist ID for Add/Remove Song to work
// with
type SongPLPair struct {
	SongID     primitive.ObjectID `bson:"id,omitempty"`
	PlaylistID primitive.ObjectID `bson:"id,omitempty"`
}
