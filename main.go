package main

import (
	"fmt"
	"net/http"
	"os"

	"groupie/handlers"
)

func main() {
	////serve style
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./styles"))))

	///server
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists/", handlers.GetArtist)
	handlers.GetArtistsData()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"

	}

	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe("0.0.0.0:"+PORT, nil)

	///Template
}
