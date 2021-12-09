package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

type hullProperties struct {
	code      string
	tons      int
	price     int
	maxHP     int
	armored   bool
	hullPanel panel
}

var hull = hullProperties{
	code:    defaultHullCode,
	tons:    defaultTons,
	price:   defaultHullPrice,
	maxHP:   defaultMaxhardpts,
	armored: false,
	hullPanel: panel{
		change:   nothingAtAll,
		selects:  make([]*widget.Select, 0),
		settings: []*widget.Form{},
		details:  []*widget.Box{},
	},
}

var (
	hullSelect    *widget.Select
	armoredSelect *widget.Check
)

var (
	detailHull          *widget.Label = widget.NewLabel("Hull")
	detailMaxHardPoints *widget.Label = widget.NewLabel("Hard Points")
)

var hullDetails *widget.Box

func (h hullProperties) init() {
	hullSelect = widget.NewSelect(hullSelectionCode, nothing)
	hullSelect.Selected = hullSelectionCode[2]

	armoredSelect = widget.NewCheck("Armored bulkheads", hull.armorChanged)
	armoredSelect.Checked = false

	hull.hullPanel.settings = []*widget.Form{widget.NewForm(
		widget.NewFormItem("hull", hullSelect),
		widget.NewFormItem("Armor", armoredSelect),
	)}

	hull.hullPanel.details = []*widget.Box{widget.NewVBox(detailHull)}
	hullDetails = hull.hullPanel.details[0]

	h.buildHull()
	drives.buildDrives()
	h.buildHardPoints()
	hullSelect.OnChanged = h.hullChanged
}

func (h hullProperties) armorChanged(armored bool) {
	hull.armored = armored
	if armored {
		h.tonsChanged(int(0.999+10*float32(hull.tons)*1.1) / 10)
	} else {
		h.tonsChanged(hull.tons)
	}
}

func (h hullProperties) hullChanged(value string) {
	index := getIndexFromHull(value)
	if index > -1 {
		hull.code = value
		hull.tons = hullSelections[index].tons
	}
	h.buildHull()
}

func (h hullProperties) tonsChanged(value int) {
	h.tons = value
	h.buildHull()
	drives.buildDrives()
	h.buildHardPoints()
}

func (h hullProperties) buildHull() {
	detailHull.SetText(fmt.Sprintf("Hull %s tons: %d, cost: %d", hull.code, hull.tons, hull.price))
	detailHull.Refresh()
}

func (h hullProperties) buildHardPoints() {
	index := getIndexFromHull(hull.code)
	if index > -1 {
		hull.tons = hullSelections[index].tons
		hull.maxHP = hull.tons / 100
		detailMaxHardPoints.SetText(fmt.Sprintf("Maximum Hardpoints: %d", hull.maxHP))
		detailMaxHardPoints.Refresh()
	}
}

func (h hullProperties) getTonnage() {
	if hull.armored {
		h.tonsChanged(int(0.999+10*float32(hull.tons)*1.1) / 10)
	} else {
		h.tonsChanged(hull.tons)
	}
}
