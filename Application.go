package main

import (
	"Groupie-Tracker/Struct"
	"fmt"
	"fyne.io/fyne/v2/app"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

func main() {
	artists, err := Struct.FetchArtists(baseURL + "/artists")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes :", err)
		return
	}

	Date, err := Struct.FetchDates(baseURL + "/dates")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des dates:", err)
		return
	}

	locationIndex, err := Struct.FetchLocationIndex(baseURL + "/locations")
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'index des lieux:", err)
		return
	}

	for _, loc := range locationIndex.Index {
		locationData, err := Struct.FetchLocationData(loc.Dates)
		if err != nil {
			fmt.Printf("Erreur lors de la récupération des données pour le lieu avec l'ID %d : %v\n", loc.ID, err)
			continue
		}

		fmt.Printf("ID: %d\n", loc.ID)
		fmt.Println("Locations:")
		for _, l := range locationData.Locations {
			fmt.Printf("- %s\n", l)
		}
		fmt.Printf("Dates: %s\n\n", locationData.Dates)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")

	var artistNames []string
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}
	fmt.Println(artistNames)
	fmt.Println(Date)
	myWindow.ShowAndRun()
}
