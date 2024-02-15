package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

type  struct {

}

func main() {
	//Créer Appli
	Appli := app.New()
	//Nouvelle Fenêtre
	Windows := Appli.NewWindow("Application")

	Windows.Resize(fyne.NewSize(400, 400))
	//Créer une url
	url, _ := url.Parse("https://www.youtube.com/watch?v=mTde30G6NLg&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=6")
	//Hyperlink Widget
	Lien := widget.NewHyperlink("Visit", url)
	//Creer un bouton avec sa fonction
	Quitter := widget.NewButton("Quitter", func() {
		Appli.Quit()
	})
	Box := container.NewVBox(
		Lien,
		Checkbox,
		Quitter)


	HomePage := func(){
		container.NewVBox(
			widget.NewLabel("HomePage"),
			widget.NewButton("AboutPage",func(){
				Appli.Quit()
			}))
	}

	AboutPage :=

	//Cocher
	Checkbox := widget.NewCheck("êtes-vous gay", func(resultat bool) {
		if resultat == true{
			Windows.SetContent(Box2)	}
	})

	Windows.SetContent(Box)
	//Impérative a la fin du main
	Windows.ShowAndRun()
}
