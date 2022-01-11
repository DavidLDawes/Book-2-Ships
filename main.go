package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var (
	saveButton = widget.NewButton("Save", saveMe)
	loadButton = widget.NewButton("Load", loadMe)
)

var (
	settings *widget.Form
	details  *widget.Box
)

func buildDetails() {
	settings = widget.NewForm()
	details = widget.NewVBox()

	hull.buildHull()
	for _, nxt := range hull.hullPanel.settings {
		for _, nxtItem := range nxt.Items {
			settings.AppendItem(nxtItem)
		}
	}
	for _, nxt := range hull.hullPanel.details {
		details.Children = append(details.Children, nxt)
	}

	drives.init()
	drives.buildDrives()
	details.Children = append(details.Children, driveDetailsBox)

	for _, nxt := range drives.panel.settings {
		for _, nxtItem := range nxt.Items {
			settings.AppendItem(nxtItem)
		}
	}

	weapons.buildWeapons()
	for _, nxt := range weapons.panel.settings {
		for _, nxtItem := range nxt.Items {
			settings.AppendItem(nxtItem)
		}
	}

	berths.buildBerths()
	buildVehicles()
	// buildTotal()
}

func main() {
	a := app.New()

	hull.init()
	drives.init()
	buildDetails()
	weapons.weaponsInit()
	berths.berthsInit()
	// shipInit()
	vehiclesInit()

	w, mainPanel := a.NewWindow("Designer"), panel{
		change:   nothingAtAll,
		selects:  make([]*widget.Select, 0),
		settings: []*widget.Form{settings},
		details:  []*widget.Box{details},
	}

	setVehicleDetails()
	berthsSelectsInit()
	berths.adjustSlider()

	drives.startup()
	//	ui := widget.NewVBox(widget.NewHBox(widget.NewLabel("Drives"), shipSettings,shipDetails,
	//  	widget.NewLabel("Weapons"), weaponSettings, weaponDetails), widget.NewLabel("Berths and Crew"),
	//		widget.NewHBox(berthSettings, berthDetails, vehicleSettings, vehicleDetails),
	// ui := widget.NewHBox(widget.NewVBox(widget.NewLabel("Drives"), shipSettings, widget.NewLabel("Weapons"),
	//      weaponSettings, widget.NewLabel("Berths"), berthSettings),
	//      widget.NewVBox(widget.NewLabel("Vehicles"), vehicleSettings),
	//  	widget.NewVBox(shipDetails, weaponDetails, berthDetails, vehicleDetails))

	ui := widget.NewHBox(
		widget.NewVBox(widget.NewLabel("Drives"), mainPanel.settings[0], widget.NewLabel("Berths & Crew"),
			berthSettings, widget.NewLabel("Weapons"), weaponSettings),
		widget.NewVBox(widget.NewLabel("Vehicles"), vehicleSettings),
		widget.NewVBox(hullDetails, driveDetailsBox, weaponDetails, berthDetails, vehicleDetails))
	// widget.NewVBox(hullDetails, berthDetails)))

	w.SetContent(ui)

	w.ShowAndRun()
}

func saveMe() {
}

func loadMe() {
}
