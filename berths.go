package main

import (
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne/widget"
)

type berthInfo struct {
	staterooms   int
	lowBerths    int
	emergencylow int
	pilots       int
	engineer     int
	stewards     int
	roboStewards int
	navigator    int
	medic        int
	gunners      int
	roboGunners  int
	exec         int
	command      int
	computer     int
	comms        int
	support      int
	roboSupport  int
	security     int
	roboSecurity int
	service      int
	roboService  int
	detailLabels []*widget.Label
}

var berths = berthInfo{
	staterooms:   4,
	lowBerths:    0,
	emergencylow: 1,
	pilots:       1,
	engineer:     1,
	stewards:     1,
	roboStewards: 0,
	navigator:    1,
	medic:        0,
	gunners:      0,
	roboGunners:  0,
	exec:         0,
	command:      0,
	computer:     0,
	comms:        0,
	support:      0,
	roboSupport:  0,
	security:     0,
	roboSecurity: 0,
	service:      0,
	roboService:  0,
	detailLabels: []*widget.Label{
		widget.NewLabel("Staterooms"), widget.NewLabel("Low berths"),
		widget.NewLabel("Emergency Low"), widget.NewLabel("Command crew"),
		widget.NewLabel("Bridge crew"), widget.NewLabel("Engineerings"),
		widget.NewLabel("Gunners"), widget.NewLabel("Stewards"),
	},
}

var (
	stateroomSlider    *widget.Slider = widget.NewSlider(4.0, 28.0)
	lowBerthSelect     *widget.Select
	emergencyLowSelect *widget.Select
	berthSettings      *widget.Form
	berthDetails       *widget.Box
	ignoreBerthChanges = false
)

func (b berthInfo) berthsInit() {
	stateroomSlider.Value = 4.0
	stateroomSlider.OnChanged = b.stateroomChanged

	lowLevel := make([]string, 401)
	for i := 0; i < 401; i++ {
		lowLevel[i] = strconv.Itoa(i)
	}

	lowBerthSelect = widget.NewSelect(lowLevel, b.lowBerthsChanged)
	emergencyLowSelect = widget.NewSelect(lowLevel, b.emergencyLowChanged)
	stateroomSlider = widget.NewSlider(4.0, 12.0)

	berthSettings = widget.NewForm(
		widget.NewFormItem("Staterooms", stateroomSlider),
		widget.NewFormItem("Low Berths", lowBerthSelect),
		widget.NewFormItem("Emergency Low Berths", emergencyLowSelect),
	)

	b.adjustSlider()

	berthDetails = widget.NewVBox()
	for _, detail := range b.detailLabels {
		berthDetails.Append(detail)
	}
	/*
		detailStaterooms,
		detailLowBerths,
		detailEmergencyLow,
		detailCommandCrew,
		detailBridgeCrew,
		detailEngCrew,
		detailGunCrew,
		detailStewardCrew,
	*/
}

func berthsSelectsInit() {
	lowBerthSelect.SetSelected("0")
	emergencyLowSelect.SetSelected("0")
}

func (b berthInfo) stateroomChanged(rooms float64) {
	rooms = math.Floor(rooms + .999999)
	if int(rooms) < b.getTotalCrew() {
		rooms = float64(b.getTotalCrew())
		ignoreBerthChanges = true
		stateroomSlider.Value = rooms
		ignoreBerthChanges = false
	}
	b.staterooms = int(rooms)
	b.buildStaterooms()
	b.buildCrew()
	//	buildTotal()
}

func (b berthInfo) lowBerthsChanged(value string) {
	if !ignoreBerthChanges {
		low, err := strconv.Atoi(value)
		if err == nil {
			if low > -1 {
				b.lowBerths = low
				b.buildLowBerths()
				b.buildCrew()
				//				buildTotal()
			}
		}
	}
}

func (b berthInfo) emergencyLowChanged(value string) {
	if !ignoreBerthChanges {
		elow, err := strconv.Atoi(value)
		if err == nil {
			if elow > -1 {
				b.emergencylow = elow
				b.buildEmergencyLow()
				b.buildCrew()
				//				buildTotal()
			}
		}
	}
}

func (b berthInfo) buildStaterooms() {
	ignoreBerthChanges = true
	b.detailLabels[0].SetText(fmt.Sprintf("Staterooms: %d, tons: %d", b.staterooms, 4*b.staterooms))
	ignoreBerthChanges = false
	b.detailLabels[0].Refresh()
}

func (b berthInfo) buildLowBerths() {
	ignoreBerthChanges = true
	b.detailLabels[1].SetText(fmt.Sprintf("Low berths: %d, tons: %d", b.lowBerths, b.lowBerths/2))
	ignoreBerthChanges = false
	b.detailLabels[1].Refresh()
}

func (b berthInfo) buildEmergencyLow() {
	ignoreBerthChanges = true
	b.detailLabels[2].SetText(fmt.Sprintf("Emergency low berths: %d, tons: %d", b.emergencylow, b.emergencylow))
	ignoreBerthChanges = false
	b.detailLabels[2].Refresh()
}

func (b berthInfo) buildCrew() {
	b.pilots = 1
	b.setEngineers()
	b.gunners = weapons.countWeapons()
	b.service = int(hull.tons/1000) * 2
	if hull.tons > 1000 {
		b.command = 1
		b.exec = 1
		b.computer = 1
		b.comms = 1
		b.navigator = 2
		b.medic = 1
		support := 4
		b.security = hull.tons / 333
		if hull.tons > 20000 {
			support = hull.tons / 2000
			if support < 4 {
				support = 4
			}
			b.support = support
		}
	} else {
		b.command = 0
		b.exec = 0
		b.computer = 0
		b.comms = 0
		b.navigator = 1
		b.medic = 0
		b.support = 0
		b.security = 0
	}
	b.setStewards()
	cmdCrew := ""
	if b.command > 0 {
		cmdCrew = "1 Commander, "
	}

	if b.exec > 0 {
		cmdCrew += fmt.Sprintf("%d Exec, ", b.exec)
	}

	if b.computer > 0 {
		cmdCrew += fmt.Sprintf("%d Computer, ", b.computer)
	}

	if b.comms > 0 {
		cmdCrew += fmt.Sprintf("%d Comms, ", b.comms)
	}
	b.detailLabels[3].SetText(cmdCrew)
	b.detailLabels[3].Refresh()

	brdgCrew := fmt.Sprintf("%d Pilot, ", b.pilots)
	if b.navigator > 0 {
		brdgCrew += fmt.Sprintf("%d Nav, ", b.navigator)
	}
	if b.medic > 0 {
		brdgCrew += fmt.Sprintf("%d Medic, ", b.medic)
	}
	b.detailLabels[4].SetText(brdgCrew)
	b.detailLabels[4].Refresh()

	b.refreshEngineeringCrew()

	if b.security > 0 {
		if b.gunners > 0 {
			b.detailLabels[6].SetText(fmt.Sprintf("%d Gunners, %d Security", b.gunners, b.security))
		} else {
			b.detailLabels[6].SetText(fmt.Sprintf("%d Security", b.security))
		}
	} else {
		if b.gunners > 0 {
			b.detailLabels[6].SetText(fmt.Sprintf("%d Gunners", b.gunners))
		} else {
			b.detailLabels[6].SetText("No Gunners, No Security")
		}
	}
	b.detailLabels[6].Refresh()

	if b.getTotalCrew() > 120 {
		b.medic = (119 + b.staterooms) / 120
	}
	b.setStewards()
	if b.getTotalCrew() > 120 {
		b.medic = (119 + b.staterooms) / 120
	}
	b.setStewards()

	if b.staterooms < b.getTotalCrew() {
		b.staterooms = b.getTotalCrew()
		b.buildStaterooms()
		stateroomSlider.Value = float64(b.staterooms)
	}

	if b.support > 0 {
		b.detailLabels[7].SetText(fmt.Sprintf("%d Stewards, %d Support", b.stewards, b.support))
	} else {
		b.detailLabels[7].SetText(fmt.Sprintf("%d Stewards", b.stewards))
	}
	b.detailLabels[7].Refresh()
}

func (b berthInfo) buildBerths() {
	b.buildStaterooms()
	b.buildLowBerths()
	b.buildCrew()
}

func (b berthInfo) setEngineers() {
	b.engineer = int((drives.j.tons + drives.m.tons + drives.p.tons) / 100.0)
}

func (b berthInfo) refreshEngineeringCrew() {
	if b.service > 0 {
		b.detailLabels[5].SetText(fmt.Sprintf("%d Engineers, %d Service", b.engineer, b.service))
	} else {
		b.detailLabels[5].SetText(fmt.Sprintf("%d Engineers", b.engineer))
	}
	b.detailLabels[5].Refresh()
}

func (b berthInfo) refreshPilots() {
	b.pilots = 1 + countVehicles()
	b.detailLabels[4].Refresh()
}

func (b berthInfo) adjustSlider() {
	maxStaterooms := float64(b.remainingTons() / 4.0)
	minStaterooms := b.getTotalCrew()
	stateroomSlider.Min = float64(minStaterooms)
	stateroomSlider.Max = float64(maxStaterooms)
}

func (b berthInfo) getTotalCrew() int {
	b.refreshPilots()
	return b.engineer + b.pilots + b.gunners + b.medic + b.stewards + b.navigator + b.exec + b.command + b.computer + b.comms + b.security + b.support + b.service
}

func (b berthInfo) getTotalRobots() int {
	return b.roboGunners + b.roboSecurity + b.roboService + b.roboStewards + b.roboSupport
}

func (b berthInfo) berthsTonsUsed() int {
	return 4*b.staterooms + (b.lowBerths+1)/2
}

func (b berthInfo) setStewards() {
	b.stewards = 0
	b.roboStewards = 0
	b.stewards = (6 + b.getTotalCrew()) / 7
}

func (b berthInfo) remainingTons() (tonsRemaining int) {
	tonsRemaining = hull.tons - weapons.weaponsTonsUsed() - drives.drivesTonsUsed() - b.berthsTonsUsed()
	return
}
