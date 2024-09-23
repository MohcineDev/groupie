package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func Search(res http.ResponseWriter, req *http.Request) {
	fromDate, _ := strconv.Atoi(req.FormValue("fromDate"))
	toDate, _ := strconv.Atoi(req.FormValue("toDate"))
	fmt.Println("fromDate : ", fromDate)
	fmt.Println("toDate : ", toDate)

	r := ArtistsObj
	for i, val := range r {
		if !(val.CreationDate >= fromDate) {
			// if !(val.CreationDate >= fromDate && val.CreationDate <= toDate) {
			// r = append(r[:i], r[i+1:]...)
			// i--
			r[i] = r[len(r)-1]
			r = r[:len(r)-1]
			i--
		}
	}

	/*
		fromAlbum, _ := strconv.Atoi(req.FormValue("fromAlbum"))
		toAlbum, _ := strconv.Atoi(req.FormValue("toAlbum"))

		for i, val := range r {
			first, _ := strconv.Atoi(val.FirstAlbum)
			if !(first >= fromAlbum && first <= toAlbum) {
				r = append(r[:i], r[i+1:]...)
			}
		}
	*/
	// member_1, _ := strconv.Atoi(req.FormValue("member_1"))
	// member_2, _ := strconv.Atoi(req.FormValue("member_2"))
	// member_3, _ := strconv.Atoi(req.FormValue("member_3"))
	// member_4, _ := strconv.Atoi(req.FormValue("member_4"))
	// member_5, _ := strconv.Atoi(req.FormValue("member_5"))
	// member_6, _ := strconv.Atoi(req.FormValue("member_6"))
	// member_7, _ := strconv.Atoi(req.FormValue("member_7"))

	// switch 0 {
	// case member_1:
	// 	r = append(r, val)
	// case member_2:
	// 	r = append(r, val)
	// case member_3:
	// 	r = append(r, val)
	// case member_4:
	// 	r = append(r, val)
	// case member_5:
	// 	r = append(r, val)
	// case member_6:
	// 	r = append(r, val)
	// case member_7:
	// 	r = append(r, val)
	// }

	/*location := req.FormValue("location")
	tf := false
	for _, val := range LocFilter.Index {
		for _, v := range val.Locations {
			if v == location {
				tf = true
				break
			}
		}
		if tf {
			tf = false
			r = append(r, ArtistsObj[val.ID])
		}
	}*/
	// fmt.Println(r)

	GetArtistsLocations()

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
	type Data struct {
		ArtistsObj        []Res
		FilteredLocations []string
	}
	data := &Data{r, FilteredLocations}

	err = tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, "Internal Server Error!!!!!!!!", http.StatusInternalServerError)
		return
	}
}
