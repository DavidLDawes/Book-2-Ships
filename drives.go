package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

// drive structure defines what we need to know about a given drive (A-Z) selection.
type drive struct {
	code string // Engine, basically size from A (smallest) to Z (biggest)
	tons int    // tonnage used by the engine
	cost int    // cost of the engine, in MCr
	perf int    // Actual performance, like -2 or -3. So a Jump drive J-2 has 2 for perf.
}

// drivedetails gathers drives, fuel, and the output UI stuff needed externally
// Most of the methods in this module are off of this driveDetails struct.
type driveDetails struct {
	j    drive
	m    drive
	p    drive
	fuel float32
}

var drives = driveDetails{
	j:    drive{"B", 20, 15, 2},
	m:    drive{"B", 3, 8, 2},
	p:    drive{"B", 7, 16, 2},
	fuel: 20,
}

// Use this to convert from hullIndex (0-9, A-Z) to the effect (the performance, M-2 or P-5 etc.) that an engine
// has on that hull.
type effectPerHullByEngine struct {
	hullIndex int
	effect    int
}

// A "row" of the effects table
// for the name (engine name, so it is A through Z) in this record we record the available hulls that can use this
// engine along with the resulting effect (i.e J-4 or P-6, gets the performance integer for any of the 3 engine types)
// The array of these is used to figure out if a given engine selection works in a given hull, and if so what the
// poerformance ends up being for that combination,.
type engineEffect struct {
	name    string
	effects []effectPerHullByEngine
}

var engineEffects = [24]engineEffect{
	{TrvIndex[0], []effectPerHullByEngine{
		{0, 2}, {1, 2}, {2, 1},
	}},
	{TrvIndex[1], []effectPerHullByEngine{
		{0, 4},
		{1, 4},
		{2, 2},
		{3, 1},
		{4, 1},
	}},
	{TrvIndex[2], []effectPerHullByEngine{
		{0, 6},
		{1, 6},
		{2, 3},
		{3, 1},
		{4, 1},
		{5, 1},
	}},
	{TrvIndex[3], []effectPerHullByEngine{
		{2, 4},
		{3, 2},
		{4, 2},
		{5, 1},
		{6, 1},
		{7, 1},
		{8, 1},
	}},
	{TrvIndex[4], []effectPerHullByEngine{
		{2, 5},
		{3, 3},
		{4, 2},
		{5, 2},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 1},
		{10, 1},
	}},
	{TrvIndex[5], []effectPerHullByEngine{
		{2, 6},
		{3, 4},
		{4, 3},
		{5, 2},
		{6, 2},
		{7, 1},
		{8, 1},
		{9, 1},
		{10, 1},
		{11, 1},
		{12, 1},
	}},
	{TrvIndex[6], []effectPerHullByEngine{
		{3, 4},
		{4, 3},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 1},
		{10, 1},
		{11, 1},
		{12, 1},
		{13, 1},
		{14, 1},
	}},
	{TrvIndex[7], []effectPerHullByEngine{
		{3, 5},
		{4, 4},
		{5, 3},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 2},
		{10, 2},
		{11, 1},
		{12, 1},
		{13, 1},
		{14, 1},
		{15, 1},
		{16, 1},
	}},
	{TrvIndex[8], []effectPerHullByEngine{
		{3, 6},
		{4, 4},
		{5, 3},
		{6, 3},
		{7, 2},
		{8, 2},
		{9, 2},
		{10, 2},
		{11, 2},
		{12, 2},
		{13, 2},
		{14, 1},
		{15, 1},
		{16, 1},
		{17, 1},
		{18, 1},
	}},
	{TrvIndex[9], []effectPerHullByEngine{
		{4, 5},
		{5, 4},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 2},
		{10, 2},
		{11, 2},
		{12, 2},
		{13, 2},
		{14, 2},
		{15, 2},
		{16, 1},
		{17, 1},
		{18, 1},
		{19, 1},
		{20, 1},
	}},
	{TrvIndex[10], []effectPerHullByEngine{
		{4, 5},
		{5, 4},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 3},
		{10, 3},
		{11, 3},
		{12, 2},
		{13, 2},
		{14, 2},
		{15, 2},
		{16, 2},
		{17, 2},
		{18, 1},
		{19, 1},
		{20, 1},
	}},
	{TrvIndex[11], []effectPerHullByEngine{
		{4, 6},
		{5, 4},
		{6, 4},
		{7, 3},
		{8, 3},
		{9, 3},
		{10, 3},
		{11, 3},
		{12, 3},
		{13, 3},
		{14, 2},
		{15, 2},
		{16, 2},
		{17, 2},
		{18, 1},
		{19, 1},
		{20, 1},
	}},
	{TrvIndex[12], []effectPerHullByEngine{
		{4, 6},
		{5, 5},
		{6, 4},
		{7, 4},
		{8, 4},
		{9, 3},
		{10, 3},
		{11, 3},
		{12, 3},
		{13, 3},
		{14, 3},
		{15, 3},
		{16, 2},
		{17, 2},
		{18, 2},
		{19, 2},
		{20, 2},
	}},
	{TrvIndex[13], []effectPerHullByEngine{
		{5, 5},
		{6, 4},
		{7, 4},
		{8, 4},
		{9, 4},
		{10, 4},
		{11, 3},
		{12, 3},
		{13, 3},
		{14, 3},
		{15, 3},
		{16, 3},
		{17, 3},
		{18, 2},
		{19, 2},
		{20, 2},
	}},
	{TrvIndex[14], []effectPerHullByEngine{
		{5, 6},
		{6, 5},
		{7, 4},
		{8, 4},
		{9, 4},
		{10, 4},
		{11, 4},
		{12, 4},
		{13, 4},
		{14, 3},
		{15, 3},
		{16, 3},
		{17, 3},
		{18, 3},
		{19, 3},
		{20, 2},
	}},
	{TrvIndex[15], []effectPerHullByEngine{
		{5, 6},
		{6, 5},
		{7, 5},
		{8, 5},
		{9, 4},
		{10, 4},
		{11, 4},
		{12, 4},
		{13, 4},
		{14, 4},
		{15, 4},
		{16, 3},
		{17, 3},
		{18, 3},
		{19, 3},
		{20, 3},
	}},
	{TrvIndex[16], []effectPerHullByEngine{
		{5, 6},
		{6, 5},
		{7, 5},
		{8, 5},
		{9, 5},
		{10, 5},
		{11, 5},
		{12, 4},
		{13, 4},
		{14, 4},
		{15, 4},
		{16, 4},
		{17, 4},
		{18, 3},
		{19, 3},
		{20, 3},
	}},
	{TrvIndex[17], []effectPerHullByEngine{
		{6, 6},
		{7, 5},
		{8, 5},
		{9, 5},
		{10, 5},
		{11, 5},
		{12, 5},
		{13, 5},
		{14, 4},
		{15, 4},
		{16, 4},
		{17, 4},
		{18, 4},
		{19, 4},
		{20, 3},
	}},
	{TrvIndex[18], []effectPerHullByEngine{
		{6, 6},
		{7, 6},
		{8, 5},
		{9, 5},
		{10, 5},
		{11, 5},
		{12, 5},
		{13, 5},
		{14, 4},
		{15, 4},
		{16, 4},
		{17, 4},
		{18, 4},
		{19, 4},
		{20, 4},
	}},
	{TrvIndex[19], []effectPerHullByEngine{
		{6, 6},
		{7, 6},
		{8, 6},
		{9, 5},
		{10, 5},
		{11, 5},
		{12, 5},
		{13, 5},
		{14, 5},
		{15, 5},
		{16, 4},
		{17, 4},
		{18, 4},
		{19, 4},
		{20, 4},
	}},
	{TrvIndex[20], []effectPerHullByEngine{
		{7, 6},
		{8, 6},
		{9, 6},
		{10, 5},
		{11, 5},
		{12, 5},
		{13, 5},
		{14, 5},
		{15, 5},
		{16, 4},
		{17, 4},
		{18, 4},
		{19, 4},
		{20, 4},
	}},
	{TrvIndex[21], []effectPerHullByEngine{
		{7, 6},
		{8, 6},
		{9, 6},
		{10, 6},
		{11, 6},
		{12, 5},
		{13, 5},
		{14, 5},
		{15, 5},
		{16, 5},
		{17, 5},
		{18, 4},
		{19, 4},
		{20, 4},
	}},
	{TrvIndex[22], []effectPerHullByEngine{
		{7, 6},
		{8, 6},
		{9, 6},
		{10, 6},
		{11, 6},
		{12, 5},
		{13, 5},
		{14, 5},
		{15, 5},
		{16, 5},
		{17, 5},
		{18, 4},
		{19, 4},
		{20, 4},
	}},
	{TrvIndex[23], []effectPerHullByEngine{
		{7, 6},
		{8, 6},
		{9, 6},
		{10, 6},
		{11, 6},
		{12, 6},
		{13, 6},
		{14, 5},
		{15, 5},
		{16, 5},
		{17, 5},
		{18, 5},
		{19, 5},
		{20, 4},
	}},
}

var minEnginePerHull = [13]int{
	1, 1, 1, 2, 3, 4, 5, 9, 15, 20, 21, 22, 23,
}

type engineDetail struct {
	name  string
	jTons int
	jCost int
	mTons int
	mCost int
	pTons int
	pCost int
}

func (d driveDetails) getIndexFromDrive(dString string) int {
	for resultInt, dMatch := range TrvIndex {
		if dMatch == dString {
			return resultInt
		}
	}
	return -1
}

func (d driveDetails) getDriveFromIndex(index int) string {
	return TrvIndex[index]
}

var engineDetails = []engineDetail{
	{TrvIndex[0], 10, 10, 1, 4, 4, 8},
	{TrvIndex[1], 15, 20, 3, 8, 7, 16},
	{TrvIndex[2], 20, 30, 5, 12, 10, 24},
	{TrvIndex[3], 25, 40, 7, 16, 13, 32},
	{TrvIndex[4], 30, 50, 9, 20, 16, 40},
	{TrvIndex[5], 35, 60, 11, 24, 19, 48},
	{TrvIndex[6], 40, 70, 13, 28, 22, 56},
	{TrvIndex[7], 45, 80, 17, 32, 25, 64},
	{TrvIndex[8], 50, 90, 19, 36, 28, 72},
	{TrvIndex[9], 55, 100, 21, 40, 31, 80},
	{TrvIndex[0], 60, 110, 23, 44, 34, 88},
	{TrvIndex[10], 65, 120, 25, 48, 37, 96},
	{TrvIndex[11], 70, 130, 27, 52, 40, 104},
	{TrvIndex[12], 75, 140, 29, 56, 43, 112},
	{TrvIndex[13], 80, 150, 31, 60, 46, 120},
	{TrvIndex[14], 85, 160, 33, 64, 49, 128},
	{TrvIndex[15], 90, 170, 35, 68, 52, 136},
	{TrvIndex[16], 95, 180, 37, 72, 55, 144},
	{TrvIndex[17], 100, 190, 39, 76, 58, 152},
	{TrvIndex[18], 105, 200, 41, 80, 61, 160},
	{TrvIndex[19], 110, 210, 43, 84, 64, 168},
	{TrvIndex[20], 115, 220, 47, 88, 67, 176},
	{TrvIndex[21], 120, 230, 49, 92, 70, 184},
	{TrvIndex[22], 125, 240, 51, 96, 73, 192},
	{TrvIndex[23], 130, 250, 53, 100, 76, 200},
}

var (
	defaultDrive    = defaultDriveCode
	defaultIndex    = 1
	detailJump      = widget.NewLabel("Jump")
	detailJumpFuel  = widget.NewLabel("Fuel")
	detailManeuver  = widget.NewLabel("Maneuver")
	detailPower     = widget.NewLabel("Power")
	driveDetailsBox = widget.NewVBox(
		widget.NewLabel("Drives"), detailJump, detailJumpFuel, detailManeuver, detailPower,
	)

	jumpSelect     = widget.NewSelect(TrvIndex, nothing)
	maneuverSelect = widget.NewSelect(TrvIndex, nothing)
	powerSelect    = widget.NewSelect(TrvIndex, nothing)
	driveForm      = widget.NewForm(
		widget.NewFormItem("Jump", jumpSelect),
		widget.NewFormItem("Maneuver", maneuverSelect),
		widget.NewFormItem("Power", powerSelect),
	)
)

func (d *driveDetails) init(form *widget.Form, box *widget.Box) {
	jumpSelect.PlaceHolder = defaultDrive
	jumpSelect.OnChanged = d.jumpChanged

	maneuverSelect.PlaceHolder = defaultDrive
	maneuverSelect.OnChanged = d.maneuverChanged

	powerSelect.PlaceHolder = defaultDrive
	powerSelect.OnChanged = d.powerChanged

	form.AppendItem(widget.NewFormItem("Drives", driveForm))

	box.Children = append(box.Children, driveDetailsBox)
	d.jumpChanged(defaultDrive)
	d.maneuverChanged(defaultDrive)
	d.powerChanged(defaultDrive)

	jumpSelect.OnChanged = d.jumpChanged
	maneuverSelect.OnChanged = d.maneuverChanged
	powerSelect.OnChanged = d.powerChanged
}

func (d *driveDetails) checkDrive(engineCode string, drv drive, checkPower bool) (good bool, effect int, newDrive string) {
	if checkPower {
		// If we are checking power (i.e. Jump orManeuver) then power == max, cap it there
		if engineCode > d.p.code {
			engineCode = d.p.code
			jumpSelect.SetSelected(fmt.Sprintf("%s", engineCode))
		}
	}

	hIndex := getIndexFromHull(hull.code)
	dIndex := d.getIndexFromDrive(engineCode)
	good = false
	if hIndex > -1 && dIndex > -1 {
		for _, fx := range engineEffects[dIndex].effects {
			if fx.hullIndex == hIndex {
				good = true
				effect = fx.effect
				newDrive = engineCode

				break
			}
		}
	}

	return
}

func (d *driveDetails) minDrives(hullCode string) (j int, m int, p int) {
	hIndex := getIndexFromHull(hullCode)
	if hIndex > 12 {
		hIndex = 12
	}
	j = minEnginePerHull[hIndex]
	m = j
	p = j

	return
}

func (d *driveDetails) jumpChanged(value string) {
	jumpSelect.OnChanged = nothing
	valid, effect, fixedDrive := d.checkDrive(value, d.j, true)
	if valid {
		d.j.cost = engineDetails[d.getIndexFromDrive(value)].jCost
		d.j.tons = engineDetails[d.getIndexFromDrive(value)].jTons
		d.j.code = fixedDrive
		d.j.perf = effect
		// StarShip.computer = computer[jump]

		d.buildJump()
		d.buildFuel()
		berths.buildCrew()
		// buildTotal()
	}
	jumpSelect.SetSelected(fmt.Sprintf("%s", d.j.code))
	jumpSelect.Refresh()
	jumpSelect.OnChanged = d.jumpChanged
}

func (d *driveDetails) maneuverChanged(value string) {
	maneuverSelect.OnChanged = nothing
	valid, effect, fixedDrive := d.checkDrive(value, d.m, true)
	if valid {
		d.m.code = fixedDrive
		d.m.cost = engineDetails[d.getIndexFromDrive(value)].mCost
		d.m.tons = engineDetails[d.getIndexFromDrive(value)].mTons
		d.m.perf = effect

		d.buildManeuver()
		berths.buildCrew()
		// buildTotal()
	}
	maneuverSelect.SetSelected(fmt.Sprintf("%s", d.m.code))
	maneuverSelect.Refresh()
	maneuverSelect.OnChanged = d.maneuverChanged
}

func (d *driveDetails) powerChanged(value string) {
	powerSelect.OnChanged = nothing
	valid, effect, fixedDrive := d.checkDrive(value, d.p, false)

	if valid {
		d.p.code = fixedDrive
		d.p.cost = engineDetails[d.getIndexFromDrive(value)].pCost
		d.p.tons = engineDetails[d.getIndexFromDrive(value)].pTons
		d.p.perf = effect

		d.buildPower()
		berths.buildCrew()
		// buildTotal()
	}
	powerSelect.SetSelected(fmt.Sprintf("%s", d.p.code))
	powerSelect.Refresh()
	powerSelect.OnChanged = d.powerChanged
}

func (d *driveDetails) buildJump() {
	detailJump.SetText(fmt.Sprintf("Jump Drive %s: J-%d, tons: %d, cost: %d MCr",
		d.j.code, d.j.perf, d.j.tons, d.j.cost))
	detailJump.Refresh()
	//	detailComputer.SetText(fmt.Sprintf("computer %d: %d tons",
	//    StarShip.jump, int(armor()*float32(computer[StarShip.jump-1])+.9999)))
	//	detailComputer.Refresh()
	berths.setEngineers()
	berths.refreshEngineeringCrew()
}

func (d *driveDetails) buildManeuver() {
	detailManeuver.SetText(fmt.Sprintf("Maneuver Drive %s: M-%d, tons: %d, cost: %d MCr",
		d.m.code, d.m.perf, d.m.tons, d.m.cost))
	detailManeuver.Refresh()
	berths.setEngineers()
	berths.refreshEngineeringCrew()
}

func (d *driveDetails) buildPower() {
	detailPower.SetText(fmt.Sprintf("Power Plant %s: P-%d, tons: %d, cost %d",
		d.p.code, d.p.perf, d.p.tons, d.p.cost))
	detailPower.Refresh()
	berths.setEngineers()
	berths.refreshEngineeringCrew()
}

func (d *driveDetails) buildFuel() {
	d.fuel = float32(hull.tons) * float32(d.p.perf) / 10.0
	detailJumpFuel.SetText(fmt.Sprintf("Jump fuel: %.1f", d.fuel))
	detailJumpFuel.Refresh()
}

func (d *driveDetails) tons() int {
	return d.j.tons + d.m.tons + d.p.tons + int(d.fuel+roundUp) // Rounded up
}

func (d *driveDetails) mCr() int {
	return d.j.cost + d.m.cost + d.p.cost
}

func (d *driveDetails) tonsX3() (int, int, int) {
	return d.j.tons, d.m.tons, d.p.tons
}

func (d *driveDetails) perfX3() (int, int, int) {
	return d.j.perf, d.m.perf, d.p.perf
}

func (d *driveDetails) costX3() (int, int, int) {
	return d.j.cost, d.m.cost, d.p.cost
}
