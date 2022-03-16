package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	settings = widget.NewForm()
	details = widget.NewVBox()

	a := app.New()

	w := a.NewWindow("Designer")

	hull.init(settings, details)
	drives.init(settings, details)
	weapons.init(settings, details)
	vehicles.init(settings, details)

	berthsSelectsInit()
	berths.adjustSlider()

	drives.startup()

	ui := widget.NewHBox(settings, details)

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
