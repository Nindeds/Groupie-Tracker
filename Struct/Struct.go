package Struct

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type LocationIndex struct {
	Index []struct {
		ID        int64    `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type LocationData struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func FetchArtists(url string) ([]Artist, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []Artist

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func FetchLocationIndex(url string) (LocationIndex, error) {
	var locationIndex LocationIndex

	response, err := http.Get(url)
	if err != nil {
		return locationIndex, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&locationIndex); err != nil {
		return locationIndex, err
	}

	return locationIndex, nil
}

func FetchLocationData(url string) (LocationData, error) {
	var locationData LocationData

	response, err := http.Get(url)
	if err != nil {
		return locationData, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&locationData); err != nil {
		return locationData, err
	}

	return locationData, nil
}

func FetchDates(url string) ([]Date, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var datesData struct {
		Index []Date `json:"index"`
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &datesData); err != nil {
		return nil, err
	}

	return datesData.Index, nil
}
