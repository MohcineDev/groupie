package main

import (
	"fmt"
	"net/http"
	"os"

	"groupie/handlers"
)

func main() {
	///server paths
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists", handlers.GetArtist)

	////serve style

	http.HandleFunc("/styles/", checkFile)   
	// run server
	const PORT = "8000"
	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe("localhost:"+PORT, nil)
}

func checkFile(res http.ResponseWriter, req *http.Request) {
	// cssFile := http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./styles"))))
	cssFile := http.StripPrefix("/styles", http.FileServer(http.Dir("./styles")))

	_, err := os.ReadFile("./" + req.URL.Path)
	if err != nil {
		handlers.ParseAndExecute("templates/404.html", res)
		return
	}
	cssFile.ServeHTTP(res, req)
}
