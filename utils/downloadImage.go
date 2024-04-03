package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Télécharge une image depuis une URL
func DownloadImage(url string) string {

	// Effectue une requête HTTP GET pour récupérer l'image à partir de l'URL.
	resp, err := http.Get(url)
	if err != nil {
		// En cas d'erreur affiche l'erreur.
		panic(err)
	}
	defer resp.Body.Close()

	// Crée un fichier temporaire pour stocker l'image téléchargée.
	file, err := ioutil.TempFile("", "image")
	if err != nil {
		// En cas d'erreur affiche l'erreur.
		panic(err)
	}
	defer file.Close()

	// Copie le contenu de la réponse HTTP (l'image téléchargée) dans le fichier temporaire.
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		// En cas d'erreur affiche l'erreur.
		panic(err)
	}

	// Return fichier temporaire créé pour l'image téléchargée.
	return file.Name()
}
