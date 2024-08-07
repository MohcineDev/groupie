package main

import (
	"cmp"
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

	// PORT := os.Getenv("PORT")
	PORT := cmp.Or(os.Getenv("PORT"), "3000")
	// if PORT == "" {
	// 	PORT = "8000"

	// }

	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe(":"+PORT, nil)

	///Template
}
