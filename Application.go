package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/app"
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

const baseURL = "https://groupietrackers.herokuapp.com/api"

func main() {
	artists, err := fetchArtists(baseURL + "/artists")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes :", err)
		return
	}

	Date, err := fetchDates(baseURL + "/dates")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des localisations:", err)
		return
	}
	locationIndex, err := fetchLocationIndex(baseURL + "/locations")
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'index des lieux:", err)
		return
	}

	// Récupération des données pour chaque lieu
	for _, loc := range locationIndex.Index {
		locationData, err := fetchLocationData(loc.Dates)
		if err != nil {
			fmt.Printf("Erreur lors de la récupération des données pour le lieu avec l'ID %d : %v\n", loc.ID, err)
			continue
		}

		// Affichage des données
		fmt.Printf("ID: %d\n", loc.ID)
		fmt.Println("Locations:")
		for _, l := range locationData.Locations {
			fmt.Printf("- %s\n", l)
		}
		fmt.Printf("Dates: %s\n\n", locationData.Dates)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")

	/*searchInput := widget.NewEntry()
	searchbutton := widget.NewButton("Rechercher", func() {

	})

	*/

	var artistNames []string
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}
	fmt.Println(Date)
	myWindow.ShowAndRun()
}

func fetchArtists(url string) ([]Artist, error) {
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

func fetchLocationIndex(url string) (LocationIndex, error) {
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

func fetchLocationData(url string) (LocationData, error) {
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

func fetchDates(url string) ([]Date, error) {
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
