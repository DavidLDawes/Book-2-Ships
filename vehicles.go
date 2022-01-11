package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type vehicleDef struct {
	vehicleName string
	vehicleTons int
	vehicleCost float32
}

var vehicleDefs = [...]vehicleDef{
	vehicleDef{"Launch", 20, 14.0},
	vehicleDef{"Boat", 30, 16.0},
	vehicleDef{"Pinnace", 40, 20.0},
	vehicleDef{"Shuttle", 95, 33.0},
	vehicleDef{"Modular Cutter", 80, 20.0},
	vehicleDef{"Light Fighter", 10, 18.0},
	vehicleDef{"Medium Fighter", 23, 40.0},
	vehicleDef{"Heavy Fighter", 55, 96.0},
}


type vehicleCounts struct {
/*	atvWheel    int
	atvTrack    int
	airRaft     int
	speeder     int
	gCarrier    int */
	launch      int
	shipsBoat   int
	pinnace     int
	cutter      int
	slowBoat    int
	slowPinnace int
	shuttle     int
	ltFighter   int
	medFighter  int
	hvyFighter  int
}

var vehicles = vehicleCounts{
/*	atvWheel:    0,
	atvTrack:    0,
	airRaft:     0,
	speeder:     0,
	gCarrier:    0, */
	launch:      0,
	shipsBoat:   0,
	pinnace:     0,
	cutter:      0,
	slowBoat:    0,
	slowPinnace: 0,
	shuttle:     0,
	ltFighter:   0,
	medFighter:  0,
	hvyFighter:  0,
}

var (
	detailSurfaceVehicles *widget.Label = widget.NewLabel("")
	detailUtilityVehicles *widget.Label = widget.NewLabel("")
	detailHighEndVehicles *widget.Label = widget.NewLabel("")
)

var (
	vehicleDetails  *widget.Box = widget.NewVBox()
	vehicleSettings *widget.Form
)

var (
/*	atvWheelSelect    *widget.Select
	atvTrackSelect    *widget.Select
	airRaftSelect     *widget.Select
	speederSelect     *widget.Select
	gCarrierSelect    *widget.Select */
	launchSelect      *widget.Select
	shipsBoatSelect   *widget.Select
	pinnaceSelect     *widget.Select
	cutterSelect      *widget.Select
	slowBoatSelect    *widget.Select
	slowPinnaceSelect *widget.Select
	shuttleSelect     *widget.Select
	ltFigherSelect    *widget.Select
	medFigherSelect   *widget.Select
	hvyFigherSelect   *widget.Select
)

var ignorevehicles = false

func vehiclesInit() {
/*	atvWheelSelect = widget.NewSelect(weaponLevel, atvWheelChanged)
	atvTrackSelect = widget.NewSelect(weaponLevel, atvTrackChanged)
	airRaftSelect = widget.NewSelect(weaponLevel, airRaftChanged)
	speederSelect = widget.NewSelect(weaponLevel, speederChanged)
	gCarrierSelect = widget.NewSelect(weaponLevel, gCarrierChanged) */
	launchSelect = widget.NewSelect(weaponLevel, launchChanged)
	shipsBoatSelect = widget.NewSelect(weaponLevel, shipsBoatChanged)
	pinnaceSelect = widget.NewSelect(weaponLevel, pinnaceChanged)
	cutterSelect = widget.NewSelect(weaponLevel, cutterChanged)
	slowBoatSelect = widget.NewSelect(weaponLevel, slowBoatChanged)
	slowPinnaceSelect = widget.NewSelect(weaponLevel, slowPinnaceChanged)
	shuttleSelect = widget.NewSelect(weaponLevel, shuttleChanged)
	ltFigherSelect = widget.NewSelect(weaponLevel, ltFigherChanged)
	medFigherSelect = widget.NewSelect(weaponLevel, medFighterChanged)
	hvyFigherSelect = widget.NewSelect(weaponLevel, hvyFighterChanged)

	vehicleSettings = widget.NewForm(
/*		widget.NewFormItem("ATV, Wheeled", atvWheelSelect),
		widget.NewFormItem("ATV, Tracked", atvTrackSelect),
		widget.NewFormItem("Air/Raft", airRaftSelect),
		widget.NewFormItem("Speeder", speederSelect),
		widget.NewFormItem("GCarrier", gCarrierSelect), */
		widget.NewFormItem("Launch", launchSelect),
		widget.NewFormItem("Ship's Boat", shipsBoatSelect),
		widget.NewFormItem("Pinnace", pinnaceSelect),
		widget.NewFormItem("Cutter", cutterSelect),
		widget.NewFormItem("Slow Boat", slowBoatSelect),
		widget.NewFormItem("Slow Pinnace", slowPinnaceSelect),
		widget.NewFormItem("Shuttle", shuttleSelect),
		widget.NewFormItem("Light Fighter", ltFigherSelect),
		widget.NewFormItem("Medium Fighter", medFigherSelect),
		widget.NewFormItem("Heavy Fighter", hvyFigherSelect),
	)

	weapons.weaponsSelectInit()
}

/*
func atvWheelChanged(value string) {
	if !ignorevehicles {
		atvw, err := strconv.Atoi(value)
		if err == nil {
			vehicles.atvWheel = atvw
			ignorevehicles = true
			buildSurface()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func atvTrackChanged(value string) {
	if !ignorevehicles {
		atvt, err := strconv.Atoi(value)
		if err == nil {
			vehicles.atvTrack = atvt
			ignorevehicles = true
			buildSurface()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func airRaftChanged(value string) {
	if !ignorevehicles {
		air, err := strconv.Atoi(value)
		if err == nil {
			vehicles.airRaft = air
			ignorevehicles = true
			buildSurface()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func speederChanged(value string) {
	if !ignorevehicles {
		spdr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.speeder = spdr
			ignorevehicles = true
			buildSurface()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func gCarrierChanged(value string) {
	if !ignorevehicles {
		gc, err := strconv.Atoi(value)
		if err == nil {
			vehicles.gCarrier = gc
			ignorevehicles = true
			buildSurface()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}
*/

func launchChanged(value string) {
	if !ignorevehicles {
		launch, err := strconv.Atoi(value)
		if err == nil {
			vehicles.launch = launch
			ignorevehicles = true
			buildUtility()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func shipsBoatChanged(value string) {
	if !ignorevehicles {
		sboat, err := strconv.Atoi(value)
		if err == nil {
			vehicles.shipsBoat = sboat
			ignorevehicles = true
			buildUtility()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func pinnaceChanged(value string) {
	if !ignorevehicles {
		pinnace, err := strconv.Atoi(value)
		if err == nil {
			vehicles.pinnace = pinnace
			ignorevehicles = true
			buildUtility()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func cutterChanged(value string) {
	if !ignorevehicles {
		cutter, err := strconv.Atoi(value)
		if err == nil {
			vehicles.cutter = cutter
			ignorevehicles = true
			buildUtility()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func slowBoatChanged(value string) {
	if !ignorevehicles {
		sloboat, err := strconv.Atoi(value)
		if err == nil {
			vehicles.slowBoat = sloboat
			ignorevehicles = true
			buildUtility()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func slowPinnaceChanged(value string) {
	if !ignorevehicles {
		slopin, err := strconv.Atoi(value)
		if err == nil {
			vehicles.slowPinnace = slopin
			ignorevehicles = true
			buildUtility()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func shuttleChanged(value string) {
	if !ignorevehicles {
		shuttle, err := strconv.Atoi(value)
		if err == nil {
			vehicles.shuttle = shuttle
			ignorevehicles = true
			buildHighEnd()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func ltFigherChanged(value string) {
	if !ignorevehicles {
		lftr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.ltFighter = lftr
			ignorevehicles = true
			buildHighEnd()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func medFighterChanged(value string) {
	if !ignorevehicles {
		mftr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.medFighter = mftr
			ignorevehicles = true
			buildHighEnd()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

func hvyFighterChanged(value string) {
	if !ignorevehicles {
		hftr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.hvyFighter = hftr
			ignorevehicles = true
			buildHighEnd()
			setVehicleDetails()
			berths.buildCrew()
			// buildTotal()
			ignorevehicles = false
		}
	}
}

/*
func buildSurface() {
	surface := getSurfaceVehicles()
	detailSurfaceVehicles.SetText(surface)
	detailSurfaceVehicles.Refresh()
}
*/

func buildUtility() {
	utility := getUtilityVehicles()
	detailUtilityVehicles.SetText(utility)
	detailUtilityVehicles.Refresh()
}

func buildHighEnd() {
	highEnd := getHighEndVehicles()
	detailHighEndVehicles.SetText(highEnd)
	detailHighEndVehicles.Refresh()
}

func buildVehicles() {
	//buildSurface()
	buildUtility()
	buildHighEnd()
	setVehicleDetails()
	// buildTotal()
}

func countVehicles() int {
	// result := vehicles.atvWheel + vehicles.atvTrack + vehicles.airRaft + vehicles.speeder + vehicles.gCarrier +
	//     vehicles.launch + vehicles.shipsBoat + vehicles.pinnace + vehicles.cutter + vehicles.slowBoat +
	//     vehicles.slowPinnace + +vehicles.shuttle + vehicles.ltFighter + vehicles.medFighter + vehicles.hvyFighter
	result := vehicles.launch + vehicles.shipsBoat + vehicles.pinnace + vehicles.cutter + vehicles.slowBoat +
		vehicles.slowPinnace + +vehicles.shuttle + vehicles.ltFighter + vehicles.medFighter + vehicles.hvyFighter

	return result
}

func vehicleTonsUsed() int {
	// result := vehicles.atvWheel*10 + vehicles.atvTrack*10 + vehicles.airRaft*4 + vehicles.speeder*6 +
	//     vehicles.gCarrier*8 + vehicles.launch*20 + vehicles.shipsBoat*30 + vehicles.pinnace*40 +
	//     vehicles.cutter*80 + vehicles.slowBoat*30 + vehicles.slowPinnace*40 + vehicles.shuttle*95 +
	//     vehicles.ltFighter*10 + vehicles.medFighter*30 + vehicles.hvyFighter*50
	result := vehicles.launch*20 + vehicles.shipsBoat*30 + vehicles.pinnace*40 + vehicles.cutter*80 +
		vehicles.slowBoat*30 + vehicles.slowPinnace*40 + vehicles.shuttle*95 + vehicles.ltFighter*10 +
		vehicles.medFighter*30 + vehicles.hvyFighter*50

	return result
}

func setVehicleDetails() {
	// Start again
	vehicleDetails.Children = make([]fyne.CanvasObject, 0)
	// Check each and add if needed
	if len(detailSurfaceVehicles.Text) > 0 {
		vehicleDetails.Children = append(vehicleDetails.Children, detailSurfaceVehicles)
	}
	if len(detailUtilityVehicles.Text) > 0 {
		vehicleDetails.Children = append(vehicleDetails.Children, detailUtilityVehicles)
	}
	if len(detailHighEndVehicles.Text) > 0 {
		vehicleDetails.Children = append(vehicleDetails.Children, detailHighEndVehicles)
	}
	vehicleDetails.Refresh()
}

/*
func getSurfaceVehicles() string {
	surface := ""
	if vehicles.atvWheel > 0 {
		surface += fmt.Sprintf("%d ATV Wheeled, ", vehicles.atvWheel)
	}
	if vehicles.atvTrack > 0 {
		surface += fmt.Sprintf("%d ATV tracked, ", vehicles.atvTrack)
	}
	if vehicles.airRaft > 0 {
		surface += fmt.Sprintf("%d Air/Raft, ", vehicles.airRaft)
	}
	if vehicles.speeder > 0 {
		surface += fmt.Sprintf("%d Speeder, ", vehicles.speeder)
	}
	if vehicles.gCarrier > 0 {
		surface += fmt.Sprintf("%d GCarrier, ", vehicles.gCarrier)
	}
	return surface
}
*/

func getUtilityVehicles() string {
	utility := ""
	if vehicles.launch > 0 {
		utility += fmt.Sprintf("%d Launch, ", vehicles.launch)
	}
	if vehicles.shipsBoat > 0 {
		utility += fmt.Sprintf("%d Ship's Boat, ", vehicles.shipsBoat)
	}
	if vehicles.pinnace > 0 {
		utility += fmt.Sprintf("%d Pinnace, ", vehicles.pinnace)
	}
	if vehicles.cutter > 0 {
		utility += fmt.Sprintf("%d Cutter, ", vehicles.cutter)
	}
	if vehicles.slowBoat > 0 {
		utility += fmt.Sprintf("%d Slow Boat, ", vehicles.slowBoat)
	}
	if vehicles.slowPinnace > 0 {
		utility += fmt.Sprintf("%d SLow Pinnace, ", vehicles.slowPinnace)
	}
	return utility
}

func getHighEndVehicles() string {
	highEnd := ""
	if vehicles.shuttle > 0 {
		highEnd += fmt.Sprintf("%d Shuttle, ", vehicles.shuttle)
	}
	if vehicles.ltFighter > 0 {
		highEnd += fmt.Sprintf("%d Light Fighter, ", vehicles.ltFighter)
	}
	if vehicles.medFighter > 0 {
		highEnd += fmt.Sprintf("%d Medium Fighter, ", vehicles.medFighter)
	}
	if vehicles.hvyFighter > 0 {
		highEnd += fmt.Sprintf("%d Heavy Fighter, ", vehicles.hvyFighter)
	}
	return highEnd
}
