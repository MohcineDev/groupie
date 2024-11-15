package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func HandleHome(res http.ResponseWriter, req *http.Request) {
	file := "./templates/index.html"

	if req.URL.Path != "/" {
		// http.NotFound(res, req)
		ParseAndExecute("templates/404.html", res)
		return
	}
	if req.Method != http.MethodGet {
		http.Error(res, "methode not Allowed!!!!!!!!", http.StatusMethodNotAllowed)
		return

	}
	////Get artist data
	GetArtistsData()

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(res, ArtistsObj)
	if err != nil {
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
}

// //
func ParseAndExecute(name string, res http.ResponseWriter) {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		http.Error(res, "Error parsing the file ", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		fmt.Println("Error when executing the template", err)
	}
}
