package main

import (
	// "C"

	"net/http"

	"Groupie-Tracker/data_structure"
	"Groupie-Tracker/utils"
	"net/url"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

func main() {
	// Récupère la liste de tous les newListArtists.
	var artists []data_structure.Artist = utils.GetAllArtists()

	// Déclaration de variables pour stocker les noms d'newListArtists et les concerts passés/futurs.
	var listArtistName []string
	var listConcertPast []string
	var listConcertFutur []string

	// Met à jour les relations entre les newListArtists et leurs concerts passés/futurs.
	utils.GetRelations(artists)

	// Crée une nouvelle application GUI.
	appli := app.New()

	// Configure le thème de l'application.
	appli.Settings().SetTheme(theme.DarkTheme())

	// Crée une nouvelle fenêtre avec un titre.
	window := appli.NewWindow("GROUPIE TRACKER")

	// Définit la taille de la fenêtre.
	window.Resize(fyne.NewSize(1000, 700))

	// Crée un menu principal pour l'application.
	mainMenu := fyne.NewMainMenu(
		// Menu "Support" avec des éléments comme "A propos de", "Documentation", et "Sponsor".
		fyne.NewMenu("Support",
			fyne.NewMenuItem("A propos de", func() {
				u, _ := url.Parse("https://meilleurs-albums.com/principaux-concerts-en-2024/")
				_ = appli.OpenURL(
					u)
			}),
			fyne.NewMenuItem("Documentation", func() {
				u, _ := url.Parse("https://www.nrj.be/article/23085/quels-newListArtists-ont-ete-les-plus-ecoutes-durant-l-annee-2022")
				_ = appli.OpenURL(u)
			}),
			fyne.NewMenuItemSeparator(), // Séparateur entre les éléments.
			fyne.NewMenuItem("Sponsor", func() {
				u, _ := url.Parse("https://nike.com/")
				_ = appli.OpenURL(u)
			}),
		))

	// Configure le menu principal de la fenêtre.
	window.SetMainMenu(mainMenu)

	// Initialisation des éléments pour afficher les informations sur l'artiste sélectionné.
	aTitle := widget.NewLabelWithStyle("Info Artist :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	aTitle.Move(fyne.NewPos(520, 0))

	// Définit un label pour afficher le nom de l'artiste sélectionné.
	aName := widget.NewLabel(" ")

	// Définit un label pour afficher les membres du groupe de l'artiste sélectionné.
	aMembers := widget.NewLabelWithStyle(" ", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})
	aMembers.Move(fyne.NewPos(520, 250))

	// Définit une image pour afficher la photo de l'artiste sélectionné.
	aImage := canvas.NewImageFromFile("")
	aImage.Resize(fyne.NewSize(150, 190))
	aImage.Move(fyne.NewPos(520, 50))

	// Définit un label pour afficher la date de création du groupe de l'artiste sélectionné.
	aCreationDate := widget.NewLabel(" ")

	// Définit un label pour afficher le titre du premier album de l'artiste sélectionné.
	aFirstAlbum := widget.NewLabel(" ")

	// Définit des labels pour les listes des concerts passés et futurs.
	aLabelPastConcert := widget.NewLabelWithStyle("Concerts Passés :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	aLabelPastConcert.Move(fyne.NewPos(520, 450))
	aLabelFuturConcert := widget.NewLabelWithStyle("Concerts Futurs :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	aLabelFuturConcert.Move(fyne.NewPos(750, 450))

	// Définit une liste pour afficher les concerts passés de l'artiste sélectionné.
	aPastConcerts := widget.NewList(
		func() int { return 1 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {},
	)
	aPastConcerts.Resize(fyne.NewSize(250, 200))
	aPastConcerts.Move(fyne.NewPos(520, 500))

	// Définit une liste pour afficher les concerts futurs de l'artiste sélectionné.
	aFuturConcerts := widget.NewList(
		func() int { return 1 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {},
	)
	aFuturConcerts.Resize(fyne.NewSize(250, 200))
	aFuturConcerts.Move(fyne.NewPos(750, 500))

	// Crée une liste des noms d'newListArtists pour affichage.
	for _, artist := range artists {
		listArtistName = append(listArtistName, artist.Name)
	}

	// Crée une liste de sélection des newListArtists.
	list := widget.NewList(
		func() int { return len(listArtistName) },
		func() fyne.CanvasObject { return widget.NewLabel("Liste des newListArtists") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(listArtistName[lii])
		},
	)
	list.Resize(fyne.NewSize(280, 500))
	list.Move(fyne.NewPos(0, 200))

	// Crée une zone de recherche pour filtrer les newListArtists.
	searchEntry := widget.NewEntry()
	searchButton := widget.NewButton("Rechercher", func() {
		// Nouvelle liste pour les résultats de la recherche
		filteredList := []string{}

		// Parcourir la liste d'origine et ajouter les éléments correspondants à la nouvelle liste
		for _, item := range listArtistName {
			if strings.Contains(strings.ToLower(item), strings.ToLower(searchEntry.Text)) {
				filteredList = append(filteredList, item)
			}
		}

		// Mettre à jour la liste avec les résultats de la recherche
		list.Length = func() int {
			return len(filteredList)
		}
		list.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		list.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(filteredList[index])
		}

		// définittion d'une nouvelle liste d'artiste (vide)
		var newListArtists []data_structure.Artist

		// parcours de l'ancienne liste d'artistes
		for _, a := range artists {
			// parcours de la liste contenant les artistes filtré
			for _, n := range filteredList {
				// test de la correspondance (est-ce que l'artiste séléctionné fait partie de la liste des filtrés)
				if a.Name == n {
					// ajout de l'artiste à la nouvelle liste si la condition est validé
					newListArtists = append(newListArtists, a)
				}
			}
		}

		// mise à jour de la liste
		artists = newListArtists

		list.Refresh()
	})
	clearButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		searchEntry.SetText("")

		// Mettre à jour la liste avec les résultats de la recherche
		list.Length = func() int {
			return len(listArtistName)
		}
		list.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		list.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(listArtistName[index])
		}

		// reset de la liste à l'état initial
		artists = utils.GetAllArtists()

		list.Refresh()
	})

	searchBar := container.NewVBox(
		searchEntry,
		searchButton,
		clearButton,
	)
	searchBar.Resize(fyne.NewSize(280, 400))

	// Contenu de l'application.
	content := container.NewWithoutLayout(
		searchBar,
		list,
	)

	separator := widget.NewSeparator()
	separator.Move(fyne.NewPos(500, 0))

	nameContainer := container.NewVBox(
		aName,
		aFirstAlbum,
	)
	nameContainer.Move(fyne.NewPos(750, 50))

	infoArtist := container.NewWithoutLayout(
		aTitle,
		aImage,
		nameContainer,
		aMembers,
		aPastConcerts,
		aFuturConcerts,
		aLabelPastConcert,
		aLabelFuturConcert,
	)
	infoArtist.Resize(fyne.NewSize(480, 800))
	infoArtist.Hide()

	list.OnSelected = func(id widget.ListItemID) {
		// Action lorsqu'un artiste est sélectionné dans la liste.

		// Récupère les informations de l'artiste sélectionné.
		artist := artists[id]

		// Met à jour les différents éléments d'affichage avec les informations de l'artiste.
		aName.Text = artist.Name
		aName.Refresh()

		// Construit la liste des membres du groupe de l'artiste.
		aMembersList := "Membre : \n"
		for _, member := range artist.Members {
			aMembersList += "- " + member + "\n"
		}
		aMembers.Text = aMembersList
		aMembers.Refresh()

		aCreationDate.Text = strconv.Itoa(artist.CreationDate)
		aImagePath := utils.DownloadImage(artist.Image)
		aImage.File = aImagePath
		aImage.Refresh()

		aFirstAlbum.Text = "Date première album : " + artist.FirstAlbum
		aFirstAlbum.Refresh()

		// Construit la liste des concerts passés de l'artiste.
		listConcertPast = nil
		for _, concert := range artists[id].PastConcert {
			listConcertPast = append(listConcertPast, concert.Location+" : "+concert.Dates)
		}
		aPastConcerts.Length = func() int {
			return len(listConcertPast)
		}
		aPastConcerts.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		aPastConcerts.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(listConcertPast[index])
		}
		aPastConcerts.Refresh()

		// Construit la liste des concerts futurs de l'artiste.
		listConcertFutur = nil
		for _, concert := range artists[id].FuturConcert {
			listConcertFutur = append(listConcertFutur, concert.Location+" : "+concert.Dates)
		}
		aFuturConcerts.Length = func() int {
			return len(listConcertFutur)
		}
		aFuturConcerts.CreateItem = func() fyne.CanvasObject {
			return widget.NewLabel("")
		}
		aFuturConcerts.UpdateItem = func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(listConcertFutur[index])
		}
		aFuturConcerts.Refresh()

		// Affiche les informations de l'artiste sélectionné.
		if infoArtist.Hidden {
			infoArtist.Show()
		}
	}

	// Configure le contenu de la fenêtre.
	window.SetContent(
		container.NewWithoutLayout(
			content,
			separator,
			infoArtist),
	)

	// Affiche et exécute l'application.
	window.ShowAndRun()
}