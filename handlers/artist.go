package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

type locations struct {
	LocationsD []string `json:"locations"`
}
type dates struct {
	DatesD []string `json:"dates"`
}
type relation struct {
	RelationD map[string][]string `json:"datesLocations"`
}

type Data struct {
	Old      Res
	Loc      locations
	Date     dates
	Relation relation
}

var (
	newlocation locations
	Dates       dates
	Relation    relation
)

func GetArtist(res http.ResponseWriter, req *http.Request) {
	// elem := strings.Split(string(req.URL.Path), "/")
	id, _ := strconv.Atoi(req.URL.Query().Get("q"))
	fmt.Println("id : ", id)
	if req.Method != http.MethodGet {
		http.Error(res, "methode not Allowed!!!!!!!!", http.StatusMethodNotAllowed)
		return
	}

	// check id
	if id > len(ArtistsObj) || id < 1 {

		http.NotFound(res, req)

		return
	}

	file := "artist.html"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		/////500

		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
	}

	Relation.RelationD = make(map[string][]string)
	er := getArtistData(id, "https://groupietrackers.herokuapp.com/api/locations/", "locations")
	er = getArtistData(id, "https://groupietrackers.herokuapp.com/api/dates/", "dates")
	er = getArtistData(id, "https://groupietrackers.herokuapp.com/api/relation/", "relation")

	if er != nil {
		/////500
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
	result := Data{
		Old:      ArtistsObj[id-1],
		Loc:      newlocation,
		Date:     Dates,
		Relation: Relation,
	}
	err = tmpl.Execute(res, result)
	if err != nil {
		/////500
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
}

func getArtistData(id int, url string, info string) error {
	DatesURL := fmt.Sprint(url, id)

	DatesData, err := http.Get(DatesURL)
	if err != nil {
		return err
	}

	content, err1 := io.ReadAll(DatesData.Body)

	if err1 != nil {
		return err
	}
	if info == "locations" {
		return json.Unmarshal(content, &newlocation)
	} else if info == "dates" {
		return json.Unmarshal(content, &Dates)
	} else if info == "relation" {
		return json.Unmarshal(content, &Relation)
	}
	return nil
}
