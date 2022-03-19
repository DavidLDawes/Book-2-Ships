package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	settings = widget.NewForm()
	vehicleSettings = widget.NewForm()
	details = widget.NewVBox()

	a := app.New()
	w := a.NewWindow("Taveller Old School Book 2 Star Ship Designer")

	hull.init(settings, details)
	drives.init(settings, details)
	weapons.init(settings, details)
	vehicles.init(vehicleSettings, details)
	berths.init(settings, details)
	summary.init(settings, details)
	summary.update()

	ui := widget.NewHBox(settings, vehicleSettings, details)
	w.SetContent(ui)
	w.ShowAndRun()
}

func saveMe() {
}

func loadMe() {
}

func nothing(value string) {
}

func nothing64(value float64) {
}

func nothingAtAll() {
}
