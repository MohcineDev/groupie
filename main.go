package main

import (
	"fmt"
	"net/http"

	"groupie/handlers"
)

func main() {
	///server paths
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists", handlers.GetArtist)

	////serve style
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./styles"))))

	// run server
	const PORT = "8000"
	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe("localhost:"+PORT, nil)
}
