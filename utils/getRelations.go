package utils

import (
	"Groupie-Tracker/data_structure"
	"encoding/json"
	"net/http"
	"time"
)

// Récupère les relations entre artistes et les ajoute aux informations d'artistes existantes.
func GetRelations(artists []data_structure.Artist) []data_structure.Artist {

	var relations data_structure.Relations

	// Crée une copie de la liste d'artistes pour mettre à jour les données.
	var newArtists []data_structure.Artist = artists

	// Effectue une requête HTTP GET pour récupérer les relations depuis l'API.
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		// En cas d'erreur affiche l'erreur.
		panic(err)
	}
	defer resp.Body.Close()

	// Décode les données JSON de la réponse HTTP dans la structure de relations.
	err = json.NewDecoder(resp.Body).Decode(&relations)
	if err != nil {
		// En cas d'erreur affiche l'erreur.
		panic(err)
	}

	// Parcourt les relations et ajoute les concerts aux artistes correspondants.
	for _, element := range relations.Relations {
		// Parcourt chaque ville et les dates de concerts associées dans l'élément.
		for city, dates := range element.DatesLocations {
			// Parcourt chaque date de concert dans la liste des dates de la ville.
			for _, dateString := range dates {
				// Parse la date au format "02-01-2006".
				date, err := time.Parse("02-01-2006", dateString)
				if err != nil {
					// En cas d'erreur lors du parsing de la date, panique et affiche l'erreur.
					panic(err)
				}

				// Ajoute le concert au tableau approprié (passé ou futur) de l'artiste.
				if date.Before(time.Now()) {
					// Si la date est antérieure à la date actuelle, ajoute le concert aux concerts passés.
					newArtists[element.ID-1].PastConcert = append(newArtists[element.ID-1].PastConcert, data_structure.Concert{city, dateString})
				} else {
					// Sinon, ajoute le concert aux concerts futurs.
					newArtists[element.ID-1].FuturConcert = append(newArtists[element.ID-1].PastConcert, data_structure.Concert{city, dateString})
				}
			}
		}
	}

	// Renvoie la liste d'artistes mise à jour avec les informations sur les concerts.
	return newArtists
}
