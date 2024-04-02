package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/artists"

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
	artists, err := fetchArtists(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes :", err)
		return
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")

	var artistNames []string
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}

	// Créer une liste des noms
	artistList := widget.NewList(
		func() int {
			return len(artistNames)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(index int, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(artistNames[index])
		},
	)

	// Afficher la liste des artistes dans une fenêtre
	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Artistes :"),
		container.NewScroll(artistList),
	))
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
