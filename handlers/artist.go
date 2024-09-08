package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
