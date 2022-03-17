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
	w := a.NewWindow("Designer")

	hull.init(settings, details)
	drives.init(settings, details)
	weapons.init(settings, details)
	vehicles.init(vehicleSettings, details)
	berths.init(settings, details)

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

func nothingAtAll() {
}
