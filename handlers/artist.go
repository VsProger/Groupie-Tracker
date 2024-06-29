package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

type Artist struct {
	ID               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	LocationsJson    string   `json:"locations"`
	ConcertDatesJson string   `json:"concertDates"`
	RelationsJson    string   `json:"relations"`
}

type DatesLocations map[string][]string

type Relations struct {
	ID             int
	DatesLocations DatesLocations
}

// type ConcertDates struct {
// 	ID           int      `json:"id"`
// 	ConcertDates []string `json:"id:dates"`
// }

type FetchArtistsFunc func() ([]Artist, error)

var FetchArtists FetchArtistsFunc = fetchArtists

func fetchArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var artists []Artist
	if jsonErr := json.Unmarshal(body, &artists); jsonErr != nil {
		return nil, jsonErr
	}

	return artists, nil
}

func fetchRelations(id string) (DatesLocations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var relations *Relations
	if jsonErr := json.Unmarshal(body, &relations); jsonErr != nil {
		return nil, jsonErr
	}
	datesLocations := relations.DatesLocations

	return datesLocations, nil
}

// func fetchDates() (Locations, error){}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		displayError(w, "Method Not Allowed", "The method is not allowed for the requested URL.", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		displayError(w, "Invalid Artist ID", "Id cannot be empty. Providing an Id is required.", http.StatusBadRequest)
		return
	}
	artistIDStr := parts[2]

	for _, char := range artistIDStr {
		if !unicode.IsDigit(char) {
			displayError(w, "Invalid Artist ID", "Id contains invalid characters. Only numerical values are allowed.", http.StatusBadRequest)
			return
		}
	}

	if strings.HasPrefix(artistIDStr, "0") && artistIDStr != "0" {
		displayError(w, "Invalid Format", "Id has an invalid format. Although '00001' is numerically equal to 1, its formatting does not match the expected input.", http.StatusBadRequest)
		return
	}

	if len(artistIDStr) > 10 {
		displayError(w, "Invalid Artist ID", "Id is too long. The input must be shorter.", http.StatusBadRequest)
		return
	}

	if strings.Contains(artistIDStr, " ") {
		displayError(w, "Invalid Artist ID", "Id should not contain spaces. The input must be an integer.", http.StatusBadRequest)
		return
	}

	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		displayError(w, "Invalid Artist ID", "Id must be a number. The entered value is not a number, which is a client error.", http.StatusBadRequest)
		return
	}
	if artistID == 0 {
		displayError(w, "Invalid Artist ID", "Id cannot be zero. The input must be a number from 1 to 52.", http.StatusBadRequest)
		return
	}
	if artistID <= 0 {
		displayError(w, "Invalid Artist ID", "Negative ID values are not allowed. This is a client error.", http.StatusBadRequest)
		return
	}
	if artistID > 52 {
		displayError(w, "Artist Not Found", "Id is out of the valid range (1-52). This artist does not exist.", http.StatusNotFound)
		return
	}

	if strings.Contains(artistIDStr, ".") {
		displayError(w, "Invalid Artist ID", "Id must be an integer. The entered value is a fractional number.", http.StatusBadRequest)
		return
	}

	artists, err := FetchArtists()
	if err != nil {
		displayError(w, "Internal Server Error", "There was an error fetching the artists.", http.StatusInternalServerError)
		return
	}

	var artist Artist
	found := false
	for _, a := range artists {
		if a.ID == artistID {
			artist = a
			found = true
			break
		}
	}
	if !found {
		displayError(w, "Artist Not Found", "The artist with the given ID does not exist.", http.StatusNotFound)
		return
	}

	datesLocations, err := fetchRelations(artistIDStr)
	// fmt.Println(datesLocations, err)

	data := struct {
		Artist         Artist
		DatesLocations DatesLocations
	}{
		Artist:         artist,
		DatesLocations: datesLocations,
	}
	// fmt.Println(data)

	err = tpl.ExecuteTemplate(w, "artist.html", data)
}
