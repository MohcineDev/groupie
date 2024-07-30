package handlers

import (
	"net/http"
	"text/template"
)

func HandleHome(res http.ResponseWriter, req *http.Request) {
	file := "index.html"

	if req.URL.Path != "/" {
		http.NotFound(res, req)
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
		/////500
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(res, ArtistsObj)
	if err != nil {
		/////500

		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
}
