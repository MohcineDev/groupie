package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type locations struct {
	LocationsD []string `json:"locations"`
}
type Data struct {
	Old Res
	Loc locations
}

func GetArtist(res http.ResponseWriter, req *http.Request) {
	var newlocation locations
	file := "artist.html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println(err)
	}
	// id, _ := fmt.Println("id : ", id)
	elem := strings.Split(string(req.URL.Path), "/")
	id, _ := strconv.Atoi(elem[len(elem)-1])
	LocationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%v", id)

	locationsData, err1 := http.Get(LocationURL)

	if err1 != nil {
		fmt.Println(err1)
	}

	content, err2 := io.ReadAll(locationsData.Body)

	json.Unmarshal(content, &newlocation)
	fmt.Println("\n\ncontent: ", newlocation)
	if err2 != nil {
		fmt.Println(err2)
	}

	if id > 51 || id < 0 {
		http.NotFound(res, req)
		return
	}
	result := Data{
		Old: ArtistsObj[id-1],
		Loc: newlocation,
	}

	err = tmpl.Execute(res, result)
	if err != nil {
		fmt.Println(err)
	}
}
