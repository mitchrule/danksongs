package main

import (
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/routes"
)

func main() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	database.InitDatabase()

	router := routes.NewRouter()

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening on 8080...")
	s.ListenAndServe()

}
