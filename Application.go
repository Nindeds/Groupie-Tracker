package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const json_url = "https://groupietrackers.herokuapp.com/api"

type Data struct {
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
	application := app.New()
	Fenetre := application.NewWindow("Groupie Tracker")
	Barre_de_recherche := widget.NewEntry()
	Bouton_recherche := widget.NewButton("Rechercher", func() {
	})

	buttons := []struct {
		nom      string
		fonction func()
	}{
		{"artiste", func() {
			Fenetre.Close()
		}},
		{"geolocalisation", func() { fmt.Println("Fonctionnalité du bouton 2") }},
		{"dates", func() { fmt.Println("Fonctionnalité du bouton 3") }},
	}

	var BoutonApplication []fyne.CanvasObject
	for _, button := range buttons {
		btn := widget.NewButton(button.nom, button.fonction)
		BoutonApplication = append(BoutonApplication, btn)
	}

	NavBarre := container.NewHBox(BoutonApplication...)
	Fenetre.Resize(fyne.NewSize(1200, 1200))

	AffichageBouton := container.NewVBox(
		NavBarre,
		container.NewVBox(
			Barre_de_recherche,
			Bouton_recherche,
		),
	)
	Fenetre.SetContent(AffichageBouton)
	Fenetre.ShowAndRun()
}

func Application(window fyne.Window) {

	window.ShowAndRun()
}
