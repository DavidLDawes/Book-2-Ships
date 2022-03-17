package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type vehicleDef struct {
	vehicleTons int
	vehicleCost float32
}

var vehicleMap = map[string]vehicleDef{
	wATVType:          {10, 0.03},
	tATVType:          {10, 0.03},
	airRaftType:       {4, 0.6},
	speederType:       {6, 1.0},
	shipsLaunchType:   {20, 14.0},
	shipsBoatType:     {30, 16.0},
	pinnaceType:       {40, 20.0},
	cutterType:        {50, 28.0},
	slowBoatType:      {30, 15.0},
	slowPinnaceType:   {40, 10.0},
	shuttleType:       {95, 33.0},
	lightFighterType:  {10, 18.0},
	mediumFighterType: {23, 40.0},
	heavyFighterType:  {55, 96.0},
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

	atvW              = widget.NewLabel("")
	atvT              = widget.NewLabel("")
	airRaft           = widget.NewLabel("")
	speeder           = widget.NewLabel("")
	gCar              = widget.NewLabel("")
	launch            = widget.NewLabel("")
	boat              = widget.NewLabel("")
	pinnace           = widget.NewLabel("")
	cutter            = widget.NewLabel("")
	slowBoat          = widget.NewLabel("")
	slowPinnace       = widget.NewLabel("")
	shuttle           = widget.NewLabel("")
	ltFighter         = widget.NewLabel("")
	medFighter        = widget.NewLabel("")
	hvyFighter        = widget.NewLabel("")
	vehicleDetailsBox = widget.NewVBox(atvW, atvT, airRaft, speeder, gCar,
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

	atvWheelItem    = widget.NewFormItem(wATVType, atvWheelSelect)
	atvTrackItem    = widget.NewFormItem(tATVType, atvTrackSelect)
	airRaftItem     = widget.NewFormItem(airRaftType, airRaftSelect)
	speederItem     = widget.NewFormItem(speederType, speederSelect)
	gCarrierItem    = widget.NewFormItem(gCarrierType, gCarrierSelect)
	launchItem      = widget.NewFormItem(shipsLaunchType, launchSelect)
	shipsBoatItem   = widget.NewFormItem(shipsBoatType, shipsBoatSelect)
	pinnaceItem     = widget.NewFormItem(pinnaceType, pinnaceSelect)
	cutterItem      = widget.NewFormItem(cutterType, cutterSelect)
	slowBoatItem    = widget.NewFormItem(slowBoatType, slowBoatSelect)
	slowPinnaceItem = widget.NewFormItem(slowPinnaceType, slowPinnaceSelect)
	shuttleItem     = widget.NewFormItem(shuttleType, shuttleSelect)
	ltFigherItem    = widget.NewFormItem(lightFighterType, ltFigherSelect)
	medFigherItem   = widget.NewFormItem(mediumFighterType, medFigherSelect)
	hvyFigherItem   = widget.NewFormItem(heavyFighterType, hvyFigherSelect)

	vehicleForm = widget.NewForm(
		atvWheelItem, atvTrackItem, airRaftItem, speederItem, gCarrierItem,
		launchItem, shipsBoatItem, pinnaceItem, cutterItem, slowBoatItem, slowPinnaceItem, shuttleItem,
		ltFigherItem, medFigherItem, hvyFigherItem,
	)
)

func (v *vehicleDetails) init(form *widget.Form, box *widget.Box) {
	atvWheelSelect.SetSelected("0")
	atvWheelSelect.OnChanged = v.atvWheelChanged
	atvWheelSelect.Show()

	atvTrackSelect.SetSelected("0")
	atvTrackSelect.OnChanged = v.atvTrackChanged
	atvTrackSelect.Show()

	airRaftSelect.SetSelected("0")
	airRaftSelect.OnChanged = v.airRaftChanged
	airRaftSelect.Show()

	speederSelect.SetSelected("0")
	speederSelect.OnChanged = v.speederChanged
	speederSelect.Show()

	gCarrierSelect.SetSelected("0")
	gCarrierSelect.OnChanged = v.gCarrierChanged
	gCarrierSelect.Show()

	launchSelect.SetSelected("0")
	launchSelect.OnChanged = v.launchChanged
	launchSelect.Show()

	shipsBoatSelect.SetSelected("0")
	shipsBoatSelect.OnChanged = v.shipsBoatChanged
	shipsBoatSelect.Show()

	pinnaceSelect.SetSelected("0")
	pinnaceSelect.OnChanged = v.pinnaceChanged
	pinnaceSelect.Show()

	cutterSelect.SetSelected("0")
	cutterSelect.OnChanged = v.cutterChanged
	cutterSelect.Show()

	slowBoatSelect.SetSelected("0")
	slowBoatSelect.OnChanged = v.slowBoatChanged
	slowBoatSelect.Show()

	slowPinnaceSelect.SetSelected("0")
	slowPinnaceSelect.OnChanged = v.slowPinnaceChanged
	slowPinnaceSelect.Show()

	shuttleSelect.SetSelected("0")
	shuttleSelect.OnChanged = v.shuttleChanged
	shuttleSelect.Show()

	ltFigherSelect.SetSelected("0")
	ltFigherSelect.OnChanged = v.ltFigherChanged
	ltFigherSelect.Show()

	medFigherSelect.SetSelected("0")
	medFigherSelect.OnChanged = v.medFighterChanged
	medFigherSelect.Show()

	hvyFigherSelect.SetSelected("0")
	hvyFigherSelect.OnChanged = v.hvyFighterChanged
	hvyFigherSelect.Show()

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

	form.AppendItem(widget.NewFormItem(vehicleTypes, vehicleForm))
}

func (v *vehicleDetails) atvWheelChanged(value string) {
	vehicleChanged(atvW, wATVType, value, &vehicles.atvWheel)
}

func (v *vehicleDetails) atvTrackChanged(value string) {
	vehicleChanged(atvT, tATVType, value, &vehicles.atvTrack)
}

func (v *vehicleDetails) airRaftChanged(value string) {
	vehicleChanged(airRaft, airRaftType, value, &vehicles.airRaft)
}

func (v *vehicleDetails) speederChanged(value string) {
	vehicleChanged(speeder, speederType, value, &vehicles.speeder)
}

func (v *vehicleDetails) gCarrierChanged(value string) {
	vehicleChanged(gCar, gCarrierType, value, &vehicles.gCar)
}

func (v *vehicleDetails) launchChanged(value string) {
	vehicleChanged(launch, shipsLaunchType, value, &vehicles.launch)
}

func (v *vehicleDetails) shipsBoatChanged(value string) {
	vehicleChanged(boat, shipsBoatType, value, &vehicles.shipsBoat)
}

func (v *vehicleDetails) pinnaceChanged(value string) {
	vehicleChanged(pinnace, pinnaceType, value, &vehicles.pinnace)
}

func (v *vehicleDetails) cutterChanged(value string) {
	vehicleChanged(cutter, cutterType, value, &vehicles.cutter)
}

func (v *vehicleDetails) slowBoatChanged(value string) {
	vehicleChanged(slowBoat, slowBoatType, value, &vehicles.slowBoat)
}

func (v *vehicleDetails) slowPinnaceChanged(value string) {
	vehicleChanged(slowPinnace, slowPinnaceType, value, &vehicles.slowPinnace)
}

func (v *vehicleDetails) shuttleChanged(value string) {
	vehicleChanged(shuttle, shuttleType, value, &vehicles.shuttle)
}

func (v *vehicleDetails) ltFigherChanged(value string) {
	vehicleChanged(ltFighter, lightFighterType, value, &vehicles.ltFighter)
}

func (v *vehicleDetails) medFighterChanged(value string) {
	vehicleChanged(medFighter, mediumFighterType, value, &vehicles.medFighter)
}

func (v *vehicleDetails) hvyFighterChanged(value string) {
	vehicleChanged(hvyFighter, heavyFighterType, value, &vehicles.hvyFighter)
}

func (v *vehicleDetails) count() int {
	result := vehicles.atvWheel + vehicles.atvTrack + vehicles.airRaft + vehicles.speeder + vehicles.gCar +
		vehicles.launch + vehicles.shipsBoat + vehicles.pinnace + vehicles.cutter + vehicles.slowBoat +
		vehicles.slowPinnace + +vehicles.shuttle + vehicles.ltFighter + vehicles.medFighter + vehicles.hvyFighter

	return result
}

func (v *vehicleDetails) tons() int {
	result := vehicles.atvWheel*vehicleMap[wATVType].vehicleTons +
		vehicles.atvTrack*vehicleMap[tATVType].vehicleTons +
		vehicles.airRaft*vehicleMap[airRaftType].vehicleTons +
		vehicles.speeder*vehicleMap[speederType].vehicleTons +
		vehicles.gCar*vehicleMap[gCarrierType].vehicleTons +
		vehicles.launch*vehicleMap[shipsLaunchType].vehicleTons +
		vehicles.shipsBoat*vehicleMap[shipsBoatType].vehicleTons +
		vehicles.pinnace*vehicleMap[pinnaceType].vehicleTons +
		vehicles.cutter*vehicleMap[cutterType].vehicleTons +
		vehicles.slowBoat*vehicleMap[slowBoatType].vehicleTons +
		vehicles.slowPinnace*vehicleMap[slowPinnaceType].vehicleTons +
		vehicles.shuttle*vehicleMap[shuttleType].vehicleTons +
		vehicles.ltFighter*vehicleMap[lightFighterType].vehicleTons +
		vehicles.medFighter*vehicleMap[mediumFighterType].vehicleTons +
		vehicles.hvyFighter*vehicleMap[heavyFighterType].vehicleTons

	return result
}

func (v *vehicleDetails) mCr() float32 {
	result := float32(vehicles.atvWheel)*vehicleMap[wATVType].vehicleCost +
		float32(vehicles.atvTrack)*vehicleMap[tATVType].vehicleCost +
		float32(vehicles.airRaft)*vehicleMap[airRaftType].vehicleCost +
		float32(vehicles.speeder)*vehicleMap[speederType].vehicleCost +
		float32(vehicles.gCar)*vehicleMap[gCarrierType].vehicleCost +
		float32(vehicles.launch)*vehicleMap[shipsLaunchType].vehicleCost +
		float32(vehicles.shipsBoat)*vehicleMap[shipsBoatType].vehicleCost +
		float32(vehicles.pinnace)*vehicleMap[pinnaceType].vehicleCost +
		float32(vehicles.cutter)*vehicleMap[cutterType].vehicleCost +
		float32(vehicles.slowBoat)*vehicleMap[slowBoatType].vehicleCost +
		float32(vehicles.slowPinnace)*vehicleMap[slowPinnaceType].vehicleCost +
		float32(vehicles.shuttle)*vehicleMap[shuttleType].vehicleCost +
		float32(vehicles.ltFighter)*vehicleMap[lightFighterType].vehicleCost +
		float32(vehicles.medFighter)*vehicleMap[mediumFighterType].vehicleCost +
		float32(vehicles.hvyFighter)*vehicleMap[heavyFighterType].vehicleCost

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

*/

func vehicleChanged(detail *widget.Label, description string, value string, setting *int) {
	vehicleCount, err := strconv.Atoi(value)

	if err == nil {
		if vehicleCount < 0 {
			vehicleCount = 0
		}

		*setting = vehicleCount

		if vehicleCount > 0 {
			detail.SetText(fmt.Sprintf(description+" x %s; %d tons, %.2f MCr",
				value,
				vehicleCount*vehicleMap[description].vehicleTons,
				float32(vehicleCount)*vehicleMap[description].vehicleCost))
			detail.Show()
		} else {
			detail.SetText("")
			detail.Hide()
		}
	}
}
