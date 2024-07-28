package handlers

import (
	"net/http"
	"text/template"
)

func HandleHome(res http.ResponseWriter, req *http.Request) {
	file := "index.html"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(res, ArtistsObj)
	if err != nil {
		panic(err)
	}
}
