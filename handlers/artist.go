package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func GetArtist(res http.ResponseWriter, req *http.Request) {
	file := "artist.html"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.URL.Path)
	elem := strings.Split(string(req.URL.Path), "/")
	id, _ := strconv.Atoi(elem[len(elem)-1])
	if id > 51 || id < 0 {
		http.NotFound(res, req)
		return
	}
	err = tmpl.Execute(res, ArtistsObj[id-1])
	if err != nil {
		fmt.Println(err)
	}
}
