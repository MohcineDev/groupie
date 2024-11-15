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

	fromAlbum, _ := strconv.Atoi(req.FormValue("fromAlbum"))
	toAlbum, _ := strconv.Atoi(req.FormValue("toAlbum"))

	r := make([]Res, len(ArtistsObj))

	copy(r, ArtistsObj)

	// for i, val := range r {
	for i := 0; i < len(r); i++ {
		first, _ := strconv.Atoi(r[i].FirstAlbum[len(r[i].FirstAlbum)-4:])
		if (r[i].CreationDate < fromDate || r[i].CreationDate > toDate) || (first < fromAlbum || first > toAlbum) {
			r = append(r[:i], r[i+1:]...)
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

	if member_1 == 0 {
		checkMembers(&r, 1)
	}

	if member_2 == 0 {
		checkMembers(&r, 2)
	}

	if member_3 == 0 {
		checkMembers(&r, 3)
	}

	if member_4 == 0 {
		checkMembers(&r, 4)
	}

	if member_5 == 0 {
		checkMembers(&r, 5)
	}

	if member_6 == 0 {
		checkMembers(&r, 6)
	}

	if member_7 == 0 {
		checkMembers(&r, 7)
	}

	if member_8 == 0 {
		checkMembers(&r, 8)
	}
	fmt.Println("\n\n----------  ", ArtistsObj[0].Location)
	// location := req.FormValue("location")
	// ArtistsObj.Location = LocFilter
	// for i := 0; i < len(r); i++ {

	// }
	// fmt.Println(r)

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

func checkMembers(r *[]Res, num int) {
	for i := 0; i < len(*r); i++ {
		if len((*r)[i].Members) == num {
			*r = append((*r)[:i], (*r)[i+1:]...)
			i--
		}
	}
}
