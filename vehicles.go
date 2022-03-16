package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type vehicleDef struct {
	vehicleName string
	vehicleTons int
	vehicleCost float32
}

var vehicleDefs = [...]vehicleDef{
	{"Wheeled ATV", 10, 0.03},
	{"Tracked ATV", 10, 0.03},
	{"Air Raft", 4, 0.6},
	{"Speeder", 6, 1.0},
	{"Launch", 20, 14.0},
	{"Boat", 30, 16.0},
	{"Pinnace", 40, 20.0},
	{"Modular Cutter", 50, 28.0},
	{"Slow Boat", 30, 15.0},
	{"Slow Pinnace", 40, 10.0},
	{"Shuttle", 95, 33.0},
	{"Light Fighter", 10, 18.0},
	{"Medium Fighter", 23, 40.0},
	{"Heavy Fighter", 55, 96.0},
}

type vehicleDetails struct {
	atvWheel    int
	atvTrack    int
	airRaft     int
	speeder     int
	gCar        int
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

var (
	vehicles = vehicleDetails{
		atvWheel:    0,
		atvTrack:    0,
		airRaft:     0,
		speeder:     0,
		gCar:        0,
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

	atvW              = widget.NewLabel("Wheeled ATV: 0")
	atvT              = widget.NewLabel("Tracked ATV: 0")
	airRaft           = widget.NewLabel("Air Raft: 0")
	speeder           = widget.NewLabel("Speeder: 0")
	gCar              = widget.NewLabel("Grav Car: 0")
	launch            = widget.NewLabel("Launch: 0")
	boat              = widget.NewLabel("Ship's Boat: 0")
	pinnace           = widget.NewLabel("Pinnace: 0")
	cutter            = widget.NewLabel("Modular CVutter: 0")
	slowBoat          = widget.NewLabel("Slow Boat: 0")
	slowPinnace       = widget.NewLabel("Slow Pinnace: 0")
	shuttle           = widget.NewLabel("Shuttle: 0")
	ltFighter         = widget.NewLabel("Light Fighter: 0")
	medFighter        = widget.NewLabel("Medium Fighter: 0")
	hvyFighter        = widget.NewLabel("Heavy Fighter: 0")
	vehicleDetailsBox = widget.NewVBox(atvW, atvT, speeder, gCar,
		launch, boat, pinnace, cutter, slowBoat, slowPinnace, shuttle,
		ltFighter, medFighter, hvyFighter)

	atvWheelSelect    = widget.NewSelect(vehicleCount, nothing)
	atvTrackSelect    = widget.NewSelect(vehicleCount, nothing)
	airRaftSelect     = widget.NewSelect(vehicleCount, nothing)
	speederSelect     = widget.NewSelect(vehicleCount, nothing)
	gCarrierSelect    = widget.NewSelect(vehicleCount, nothing)
	launchSelect      = widget.NewSelect(vehicleCount, nothing)
	shipsBoatSelect   = widget.NewSelect(vehicleCount, nothing)
	pinnaceSelect     = widget.NewSelect(vehicleCount, nothing)
	cutterSelect      = widget.NewSelect(vehicleCount, nothing)
	slowBoatSelect    = widget.NewSelect(vehicleCount, nothing)
	slowPinnaceSelect = widget.NewSelect(vehicleCount, nothing)
	shuttleSelect     = widget.NewSelect(vehicleCount, nothing)
	ltFigherSelect    = widget.NewSelect(vehicleCount, nothing)
	medFigherSelect   = widget.NewSelect(vehicleCount, nothing)
	hvyFigherSelect   = widget.NewSelect(vehicleCount, nothing)

	atvWheelItem    = widget.NewFormItem("Wheeled ATV", atvWheelSelect)
	atvTrackItem    = widget.NewFormItem("Wheeled ATV", atvTrackSelect)
	airRaftItem     = widget.NewFormItem("Wheeled ATV", airRaftSelect)
	speederItem     = widget.NewFormItem("Wheeled ATV", speederSelect)
	gCarrierItem    = widget.NewFormItem("Wheeled ATV", gCarrierSelect)
	launchItem      = widget.NewFormItem("Wheeled ATV", launchSelect)
	shipsBoatItem   = widget.NewFormItem("Wheeled ATV", shipsBoatSelect)
	pinnaceItem     = widget.NewFormItem("Wheeled ATV", pinnaceSelect)
	cutterItem      = widget.NewFormItem("Wheeled ATV", cutterSelect)
	slowBoatItem    = widget.NewFormItem("Wheeled ATV", slowBoatSelect)
	slowPinnaceItem = widget.NewFormItem("Wheeled ATV", slowPinnaceSelect)
	shuttleItem     = widget.NewFormItem("Wheeled ATV", shuttleSelect)
	ltFigherItem    = widget.NewFormItem("Wheeled ATV", ltFigherSelect)
	medFigherItem   = widget.NewFormItem("Wheeled ATV", medFigherSelect)
	hvyFigherItem   = widget.NewFormItem("Wheeled ATV", hvyFigherSelect)

	vehicleForm = widget.NewForm(
		atvWheelItem, atvTrackItem, airRaftItem, speederItem, gCarrierItem,
		launchItem, shipsBoatItem, pinnaceItem, cutterItem, slowBoatItem, slowPinnaceItem, shuttleItem,
		ltFigherItem, medFigherItem, hvyFigherItem,
	)
)

func (v *vehicleDetails) init(form *widget.Form, box *widget.Box) {
	atvWheelSelect.SetSelected("0")
	atvWheelSelect.OnChanged = v.atvWheelChanged

	atvTrackSelect.SetSelected("0")
	atvTrackSelect.OnChanged = v.atvTrackChanged

	airRaftSelect.SetSelected("0")
	airRaftSelect.OnChanged = v.airRaftChanged

	speederSelect.SetSelected("0")
	speederSelect.OnChanged = v.speederChanged

	gCarrierSelect.SetSelected("0")
	gCarrierSelect.OnChanged = v.gCarrierChanged

	launchSelect.SetSelected("0")
	launchSelect.OnChanged = v.launchChanged

	shipsBoatSelect.SetSelected("0")
	shipsBoatSelect.OnChanged = v.shipsBoatChanged

	pinnaceSelect.SetSelected("0")
	pinnaceSelect.OnChanged = v.pinnaceChanged

	cutterSelect.SetSelected("0")
	cutterSelect.OnChanged = v.cutterChanged

	slowBoatSelect.SetSelected("0")
	slowBoatSelect.OnChanged = v.slowBoatChanged

	slowPinnaceSelect.SetSelected("0")
	slowPinnaceSelect.OnChanged = v.slowPinnaceChanged

	shuttleSelect.SetSelected("0")
	shuttleSelect.OnChanged = v.shuttleChanged

	ltFigherSelect.SetSelected("0")
	ltFigherSelect.OnChanged = v.ltFigherChanged

	medFigherSelect.SetSelected("0")
	medFigherSelect.OnChanged = v.medFighterChanged

	hvyFigherSelect.SetSelected("0")
	hvyFigherSelect.OnChanged = v.hvyFighterChanged

	box.Children = append(box.Children, vehicleDetailsBox)

	atvW.Hide()
	atvT.Hide()
	airRaft.Hide()
	speeder.Hide()
	gCar.Hide()
	launch.Hide()
	boat.Hide()
	pinnace.Hide()
	cutter.Hide()
	slowBoat.Hide()
	slowPinnace.Hide()
	shuttle.Hide()
	ltFighter.Hide()
	medFighter.Hide()
	hvyFighter.Hide()
}

func (v *vehicleDetails) atvWheelChanged(value string) {
	atvw, err := strconv.Atoi(value)
	if err == nil {
		vehicles.atvWheel = atvw
		if atvw > 0 {
			atvW.SetText(fmt.Sprintf("Wheeled ATV: %s", value))
			atvW.Show()
		} else {
			atvW.SetText("")
			atvW.Hide()
		}
	}
}

func (v *vehicleDetails) atvTrackChanged(value string) {
	atvt, err := strconv.Atoi(value)
	if err == nil {
		vehicles.atvTrack = atvt
		if atvt > 0 {
			atvT.SetText(fmt.Sprintf("Tracked ATV: %s", value))
			atvT.Show()
		} else {
			atvT.SetText("")
			atvT.Hide()
		}
	}
}

func (v *vehicleDetails) airRaftChanged(value string) {
	air, err := strconv.Atoi(value)
	if err == nil {
		vehicles.airRaft = air
		if air > 0 {
			airRaft.SetText(fmt.Sprintf("Tracked ATV: %s", value))
			airRaft.Show()
		} else {
			airRaft.SetText("")
			airRaft.Hide()
		}
	}
}

func (v *vehicleDetails) speederChanged(value string) {
	spdr, err := strconv.Atoi(value)
	if err == nil {
		vehicles.speeder = spdr
		if spdr > 0 {
			speeder.SetText(fmt.Sprintf("Speeder: %s", value))
			speeder.Show()
		} else {
			speeder.SetText("")
			speeder.Hide()
		}
	}
}

func (v *vehicleDetails) gCarrierChanged(value string) {
	gc, err := strconv.Atoi(value)
	if err == nil {
		vehicles.gCar = gc
		if gc > 0 {
			gCar.SetText(value)
			gCar.Show()
		} else {
			gCar.SetText("0")
			gCar.Hide()
		}
	}
}

func (v *vehicleDetails) launchChanged(value string) {
	launches, err := strconv.Atoi(value)
	if err == nil {
		vehicles.launch = launches
		if launches > 0 {
			launch.SetText(value)
			launch.Show()
		} else {
			launch.SetText("0")
			launch.Hide()
		}
	}
}

func (v *vehicleDetails) shipsBoatChanged(value string) {
	sboat, err := strconv.Atoi(value)
	if err == nil {
		vehicles.shipsBoat = sboat
		if sboat > 0 {
			slowBoat.SetText(value)
			slowBoat.Show()
		} else {
			slowBoat.SetText("0")
			slowBoat.Hide()
		}
	}
}

func (v *vehicleDetails) pinnaceChanged(value string) {
	pinnaceCount, err := strconv.Atoi(value)
	if err == nil {
		vehicles.pinnace = pinnaceCount
		if pinnaceCount > 0 {
			pinnace.SetText(value)
			pinnace.Show()
		} else {
			pinnace.SetText("0")
			pinnace.Hide()
		}
	}
}

func (v *vehicleDetails) cutterChanged(value string) {
	cutterCount, err := strconv.Atoi(value)
	if err == nil {
		vehicles.cutter = cutterCount
		if cutterCount > 0 {
			cutter.SetText(value)
			cutter.Show()
		} else {
			cutter.SetText("0")
			cutter.Hide()
		}
	}
}

func (v *vehicleDetails) slowBoatChanged(value string) {
	sloboat, err := strconv.Atoi(value)
	if err == nil {
		vehicles.slowBoat = sloboat
		if sloboat > 0 {
			slowBoat.SetText(value)
			slowBoat.Show()
		} else {
			slowBoat.SetText("0")
			slowBoat.Hide()
		}
	}
}

func (v *vehicleDetails) slowPinnaceChanged(value string) {
	slopin, err := strconv.Atoi(value)
	if err == nil {
		vehicles.slowPinnace = slopin
		if slopin > 0 {
			slowPinnace.SetText(value)
			slowPinnace.Show()
		} else {
			slowPinnace.SetText("0")
			slowPinnace.Hide()
		}
	}
}

func (v *vehicleDetails) shuttleChanged(value string) {
	shuttleCount, err := strconv.Atoi(value)
	if err == nil {
		vehicles.shuttle = shuttleCount
		if shuttleCount > 0 {
			shuttle.SetText(value)
			shuttle.Show()
		} else {
			shuttle.SetText("0")
			shuttle.Hide()
		}
	}
}

func (v *vehicleDetails) ltFigherChanged(value string) {
	lftr, err := strconv.Atoi(value)
	if err == nil {
		vehicles.ltFighter = lftr
		if lftr > 0 {
			ltFighter.SetText(value)
			ltFighter.Show()
		} else {
			ltFighter.SetText("0")
			ltFighter.Hide()
		}
	}
}

func (v *vehicleDetails) medFighterChanged(value string) {
	mftr, err := strconv.Atoi(value)
	if err == nil {
		vehicles.medFighter = mftr
		if mftr > 0 {
			medFighter.SetText(value)
			medFighter.Show()
		} else {
			medFighter.SetText("0")
			medFighter.Hide()
		}
	}
}

func (v *vehicleDetails) hvyFighterChanged(value string) {
	hftr, err := strconv.Atoi(value)
	if err == nil {
		vehicles.hvyFighter = hftr
		if hftr > 0 {
			hvyFighter.SetText(value)
			hvyFighter.Show()
		} else {
			hvyFighter.SetText("0")
			hvyFighter.Hide()
		}
	}
}

func countVehicles() int {
	result := vehicles.atvWheel + vehicles.atvTrack + vehicles.airRaft + vehicles.speeder + vehicles.gCar +
		vehicles.launch + vehicles.shipsBoat + vehicles.pinnace + vehicles.cutter + vehicles.slowBoat +
		vehicles.slowPinnace + +vehicles.shuttle + vehicles.ltFighter + vehicles.medFighter + vehicles.hvyFighter

	return result
}

func vehicleTonsUsed() int {
	result := vehicles.atvWheel*10 + vehicles.atvTrack*10 + vehicles.airRaft*4 + vehicles.speeder*6 +
		vehicles.gCar*8 + vehicles.launch*20 + vehicles.shipsBoat*30 + vehicles.pinnace*40 +
		vehicles.cutter*80 + vehicles.slowBoat*30 + vehicles.slowPinnace*40 + vehicles.shuttle*95 +
		vehicles.ltFighter*10 + vehicles.medFighter*30 + vehicles.hvyFighter*50

	return result
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
