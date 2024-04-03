package Groupie_Tracker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2/app"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"


func main() {
	artists, err := fetchArtists(baseURL + "/artists")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes :", err)
		return
	}

	Location, err := fetchLocations(baseURL + "/locations")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des localisations:", err)
		return
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
	fmt.Println(Location)
	myWindow.ShowAndRun()
}

func fetch

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

func fetchLocations(url string) ([]string, error) {
	var locationIndex LocationIndex

	// Envoyer une requête GET pour récupérer l'URL des localisations
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Décoder la réponse JSON pour obtenir l'URL des localisations
	err = json.NewDecoder(response.Body).Decode(&locationIndex)
	if err != nil {
		return nil, err
	}

	// Maintenant, vous pouvez envoyer une nouvelle requête GET à l'URL des localisations pour récupérer les données réelles
	locationsResponse, err := http.Get(locationIndex.LocationsURL)
	if err != nil {
		return nil, err
	}
	defer locationsResponse.Body.Close()

	// Décodez la réponse JSON pour obtenir les données de localisation réelles
	var locations []string
	err = json.NewDecoder(locationsResponse.Body).Decode(&locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
