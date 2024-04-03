package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/widget"
	"net/http"

	"fyne.io/fyne/v2/app"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

type Location struct {
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

func main() {
	artists, err := fetchArtists(baseURL + "/artists")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes :", err)
		return
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")

	searchInput := widget.NewEntry()
	searchbutton := widget.NewButton("Rechercher", func() {

	})

	var artistNames []string
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}
	fmt.Println(artists)

	myWindow.ShowAndRun()
}

func fetchArtists(apiURL string) ([]Artist, error) {
	var artists []Artist

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func fetchLocations(path string) ([]Location, error) {
	var data []Location

	response, err := http.Get(baseURL + path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
