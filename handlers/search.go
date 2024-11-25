package handlers

import (
	"net/http"
	"strconv"
	"text/template"
)

func Search(res http.ResponseWriter, req *http.Request) {
	GetArtistsData()
	GetArtistsLocations()

	for i := 0; i < len(ArtistsObj); i++ {
		ArtistsObj[i].Location.Locations = LocFilter.Index[i].Locations
	}

	fromDate, _ := strconv.Atoi(req.FormValue("fromDate"))
	toDate, _ := strconv.Atoi(req.FormValue("toDate"))

	fromAlbum, _ := strconv.Atoi(req.FormValue("fromAlbum"))
	toAlbum, _ := strconv.Atoi(req.FormValue("toAlbum"))

	filteredRes := make([]Res, len(ArtistsObj))
	copy(filteredRes, ArtistsObj)

	for i := 0; i < len(filteredRes); i++ {
		first, _ := strconv.Atoi(filteredRes[i].FirstAlbum[len(filteredRes[i].FirstAlbum)-4:])
		if (filteredRes[i].CreationDate < fromDate || filteredRes[i].CreationDate > toDate) || (first < fromAlbum || first > toAlbum) {
			filteredRes = append(filteredRes[:i], filteredRes[i+1:]...)
			i--
		}
	}

	member_1, _ := strconv.Atoi(req.FormValue("member_1"))
	member_2, _ := strconv.Atoi(req.FormValue("member_2"))
	member_3, _ := strconv.Atoi(req.FormValue("member_3"))
	member_4, _ := strconv.Atoi(req.FormValue("member_4"))
	member_5, _ := strconv.Atoi(req.FormValue("member_5"))
	member_6, _ := strconv.Atoi(req.FormValue("member_6"))
	member_7, _ := strconv.Atoi(req.FormValue("member_7"))
	member_8, _ := strconv.Atoi(req.FormValue("member_8"))

	if member_1 != 0 || member_2 != 0 || member_3 != 0 || member_4 != 0 || member_5 != 0 || member_6 != 0 || member_7 != 0 || member_8 != 0 {
		if member_1 == 0 {
			checkMembers(&filteredRes, 1)
		}

		if member_2 == 0 {
			checkMembers(&filteredRes, 2)
		}

		if member_3 == 0 {
			checkMembers(&filteredRes, 3)
		}

		if member_4 == 0 {
			checkMembers(&filteredRes, 4)
		}

		if member_5 == 0 {
			checkMembers(&filteredRes, 5)
		}

		if member_6 == 0 {
			checkMembers(&filteredRes, 6)
		}

		if member_7 == 0 {
			checkMembers(&filteredRes, 7)
		}

		if member_8 == 0 {
			checkMembers(&filteredRes, 8)
		}
	}

	location := req.FormValue("location")
	if location != "all" {
		for i := 0; i < len(filteredRes); i++ {
			tf := false
			for j := 0; j < len(filteredRes[i].Location.Locations); j++ {
				if location == filteredRes[i].Location.Locations[j] {
					tf = true
					break
				}
			}
			if !tf {
				filteredRes = append(filteredRes[:i], filteredRes[i+1:]...)
				i--
			}
		}
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
	type Data struct {
		ArtistsObj        []Res
		FilteredLocations []string
	}
	data := &Data{filteredRes, FilteredLocations}

	err = tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
}

func checkMembers(r *[]Res, num int) {
	for i := 0; i < len(*r); i++ {
		if len((*r)[i].Members) == num {
			*r = append((*r)[:i], (*r)[i+1:]...)
			i--
		}
	}
}
