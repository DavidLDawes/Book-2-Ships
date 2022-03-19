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
}

var berths = berthInfo{}

var (
	lowLevel = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
		"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
	}

	room = widget.NewLabel("Staterooms")
	low  = widget.NewLabel("Low berths")
	eLow = widget.NewLabel("Emergency Low")
	cmd  = widget.NewLabel("Command crew")
	brdg = widget.NewLabel("Bridge crew")
	eng  = widget.NewLabel("Engineerings")
	gun  = widget.NewLabel("Gunners")
	stw  = widget.NewLabel("Stewards")
	plt  = widget.NewLabel("Pilots")

	berthDetailsBox    = widget.NewVBox(room, low, eLow, cmd, brdg, eng, gun, stw)
	stateroomSlider    = widget.NewSlider(0.0, float64(len(lowLevel)))
	lowBerthSelect     = widget.NewSelect(lowLevel, nothing)
	emergencyLowSelect = widget.NewSelect(lowLevel, nothing)
	berthsForm         = widget.NewForm(
		widget.NewFormItem("Staterooms", stateroomSlider),
		widget.NewFormItem("Low Berths", lowBerthSelect),
		widget.NewFormItem("Emergency Low Berths", emergencyLowSelect))
)

func (b berthInfo) init(form *widget.Form, box *widget.Box) {
	berths.staterooms = 4
	berths.pilots = 1
	berths.engineer = 1
	berths.stewards = 1
	berths.navigator = 1

	stateroomSlider.Value = 4.0
	stateroomSlider.OnChanged = b.stateroomChanged
	stateroomSlider.Show()
	b.adjustSlider()

	lowBerthSelect.PlaceHolder = noneString
	lowBerthSelect.Selected = noneString
	lowBerthSelect.OnChanged = b.lowBerthsChanged
	lowBerthSelect.Show()

	emergencyLowSelect.PlaceHolder = noneString
	emergencyLowSelect.Selected = noneString
	emergencyLowSelect.OnChanged = b.emergencyLowChanged
	emergencyLowSelect.Show()

	low.Hide()
	eLow.Hide()
	cmd.Hide()

	brdg.SetText("1x Navigator")
	brdg.Show()

	eng.SetText("1x Engineer")
	eng.Show()

	plt.SetText("1x Pilot")
	plt.Show()

	stw.SetText("1x Steward")
	stw.Show()

	form.AppendItem(widget.NewFormItem("Berths", berthsForm))

	box.Children = append(box.Children, berthDetailsBox)

	lowBerthSelect.SetSelected("0")
	emergencyLowSelect.SetSelected("0")
	b.stateroomChanged(4.0)
	b.lowBerthsChanged("0")
	b.emergencyLowChanged("0")
}

func (b berthInfo) stateroomChanged(rooms float64) {
	stateroomSlider.OnChanged = nothing64
	rooms = math.Floor(rooms + roundUp)
	if int(rooms) < b.getTotalCrew() {
		rooms = float64(b.getTotalCrew())
		stateroomSlider.Value = rooms
	}
	b.staterooms = int(rooms)
	b.buildStaterooms()
	b.buildCrew()
	summary.update()
	stateroomSlider.OnChanged = b.stateroomChanged
}

func (b berthInfo) lowBerthsChanged(value string) {
	low, err := strconv.Atoi(value)
	if err == nil {
		if low > -1 {
			b.lowBerths = low
			b.buildLowBerths()
			b.buildCrew()
			summary.update()
		}
	}
}

func (b berthInfo) emergencyLowChanged(value string) {
	elow, err := strconv.Atoi(value)
	if err == nil {
		if elow > -1 {
			b.emergencylow = elow
			if elow > 0 {
				b.buildEmergencyLow()
			}
			b.buildCrew()
			summary.update()
		}
	}
}

func (b berthInfo) buildStaterooms() {
	room.SetText(fmt.Sprintf("Staterooms: %d, tons: %d", b.staterooms, 4*b.staterooms))
	room.Refresh()
}

func (b berthInfo) buildLowBerths() {
	low.SetText(fmt.Sprintf("Low berths: %d, tons: %d", b.lowBerths, b.lowBerths/2))
	low.Refresh()
}

func (b berthInfo) buildEmergencyLow() {
	if b.emergencylow > 0 {
		eLow.SetText(fmt.Sprintf("Emergency low berths: %d, tons: %d", b.emergencylow, b.emergencylow))
		eLow.Refresh()
		eLow.Show()
	} else {
		eLow.Hide()
	}
}

func (b berthInfo) buildCrew() {
	b.pilots = 1 + vehicles.count()
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
	cmd.SetText(cmdCrew)
	cmd.Refresh()

	brdgCrew := fmt.Sprintf("%d Pilot, ", b.pilots)
	if b.navigator > 0 {
		brdgCrew += fmt.Sprintf("%d Nav, ", b.navigator)
	}
	if b.medic > 0 {
		brdgCrew += fmt.Sprintf("%d Medic, ", b.medic)
	}
	brdg.SetText(brdgCrew)
	brdg.Refresh()

	b.refreshEngineeringCrew()

	if b.security > 0 {
		if b.gunners > 0 {
			gun.SetText(fmt.Sprintf("%d Gunners, %d Security", b.gunners, b.security))
		} else {
			gun.SetText(fmt.Sprintf("%d Security", b.security))
		}
	} else {
		if b.gunners > 0 {
			gun.SetText(fmt.Sprintf("%d Gunners", b.gunners))
		} else {
			gun.SetText("No Gunners, No Security")
		}
	}
	gun.Refresh()

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
		stw.SetText(fmt.Sprintf("%d Stewards, %d Support", b.stewards, b.support))
	} else {
		stw.SetText(fmt.Sprintf("%d Stewards", b.stewards))
	}
	stw.Refresh()
	summary.update()
}

func (b berthInfo) buildBerths() {
	b.buildStaterooms()
	b.buildLowBerths()
	b.buildCrew()
}

func (b berthInfo) setEngineers() {
	if drives.tons() < 150 {
		b.engineer = int(drives.tons()+99) / 100.0
	} else {
		b.engineer = int(drives.j.tons+99)/100.0 + int(drives.m.tons+99)/100.0 + int(drives.p.tons+99)/100.0
	}
}

func (b berthInfo) refreshEngineeringCrew() {
	if b.service > 0 {
		eng.SetText(fmt.Sprintf("%dx Engineers, %dx Service", b.engineer, b.service))
	} else {
		eng.SetText(fmt.Sprintf("%dx Engineers", b.engineer))
	}
	eng.Refresh()
}

func (b berthInfo) refreshPilots() {
	b.pilots = 1 + vehicles.count()
	plt.SetText(fmt.Sprintf("%dx Pilots", b.pilots))
	plt.Refresh()
}

func (b berthInfo) adjustSlider() {
	maxStaterooms := float64(summary.cargo) / 4.0
	minStaterooms := b.getTotalCrew()
	stateroomSlider.Min = float64(minStaterooms)
	stateroomSlider.Max = float64(maxStaterooms)
}

func (b berthInfo) getTotalCrew() int {
	b.refreshPilots()
	return b.engineer + b.pilots + b.gunners + b.medic + b.stewards + b.navigator + b.exec + b.command +
		b.computer + b.comms + b.security + b.support + b.service
}

func (b berthInfo) getTotalRobots() int {
	return b.roboGunners + b.roboSecurity + b.roboService + b.roboStewards + b.roboSupport
}

func (b berthInfo) tons() float32 {
	return 4.0*float32(b.staterooms) + float32(b.lowBerths)/2 + float32(b.emergencylow)
}

func (b berthInfo) setStewards() {
	b.stewards = 0
	b.roboStewards = 0
	b.stewards = (6 + b.getTotalCrew()) / 7
}

func (b berthInfo) mCr() float32 {
	return float32(berths.staterooms)*0.5 + float32(berths.emergencylow)*0.1 + float32(berths.lowBerths)*0.05
}
