package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Res struct {
	ID           int      `json:"id"`
	Img          string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     Location
}

var ArtistsObj []Res

type LocationsData struct {
	Index []Location `json:"index"`
}

var LocFilter LocationsData

func GetArtistsData() {
	const URL = "https://groupietrackers.herokuapp.com/api/artists"

	// query api
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	////read all data from the incoming byte stream
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// json.NGetArtistsDataewDecoder(res).Decode(any)

	json.Unmarshal(resData, &ArtistsObj)

	// fmt.Println("after : ", &ArtistsObj)
	// var a *[]Res
	a := &ArtistsObj
	*a = append(ArtistsObj[0:20], ArtistsObj[24:]...)
	ArtistsObj = *a
}

// /for filter select dropdown
var FilteredLocations []string

func GetArtistsLocations() {
	const URL = "https://groupietrackers.herokuapp.com/api/locations"
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	a, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	er := json.Unmarshal(a, &LocFilter)

	if er != nil {
		fmt.Println("er : ", er)
	}
	allLocations := make(map[string]bool)
	for _, v := range LocFilter.Index {
		for _, k := range v.Locations {
			if !allLocations[k] {
				allLocations[k] = true
				FilteredLocations = append(FilteredLocations, k)
			}
		}
	}
}
