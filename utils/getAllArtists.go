package utils

import (
	"Groupie-Tracker/data_structure"
	"encoding/json"
	"net/http"
)

// Récupère la liste de tous les artistes depuis l'API et la retourne.
func GetAllArtists() []data_structure.Artist {

	// Déclare une variable pour stocker la liste des artistes.
	var artists []data_structure.Artist

	
	// Effectue une requête HTTP GET pour récupérer les données des artistes depuis l'API.
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		// En cas d'erreur lors de la requête, panique et affiche l'erreur.
		panic(err)
	}
	defer resp.Body.Close()

	// Décode les données JSON de la réponse HTTP dans la liste des artistes.
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		// En cas d'erreur renvoie nil.
		return nil
	}

	// Return la liste des artistes récupérée depuis l'API.
	return artists
}
