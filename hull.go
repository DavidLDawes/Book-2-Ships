package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

type hullProperties struct {
	code    string
	tons    int
	price   int
	maxHP   int
	armored bool
}

var hull = hullProperties{
	code:    defaultHullCode,
	tons:    defaultTons,
	price:   defaultHullPrice,
	maxHP:   defaultMaxhardpts,
	armored: false,
}

var (
	hullSelect          *widget.Select
	detailHull          *widget.Label = widget.NewLabel("Hull")
	detailMaxHardPoints *widget.Label = widget.NewLabel("Hard Points")
	hullDetails         *widget.Box   = widget.NewVBox(detailHull, detailMaxHardPoints)
)

func (h *hullProperties) init(form *widget.Form, box *widget.Box) {
	hullSelect = widget.NewSelect(hullSelectionCode, nothing)
	hullSelect.Selected = hullSelectionCode[2]
	form.AppendItem(widget.NewFormItem("hull", hullSelect))

	box.Children = append(box.Children, hullDetails)

	h.updateHull()
	h.updateHardPoints()
	hullSelect.OnChanged = h.hullChanged
}

func (h *hullProperties) armorChanged(armored bool) {
	hull.armored = armored
}

func (h *hullProperties) hullChanged(value string) {
	index := getIndexFromHull(value)
	if index > -1 {
		hull.code = value
		hull.tons = hullSelections[index].tons
		hull.price = hullSelections[index].price
		j, m, p := drives.minDrives(value)
		drives.jumpChanged(TrvIndex[j])
		drives.maneuverChanged(TrvIndex[m])
		drives.powerChanged(TrvIndex[p])
	}
	h.updateHull()
}

func (h *hullProperties) updateHull() {
	detailHull.SetText(fmt.Sprintf("Hull %s tons: %d, cost: %d MCr", hull.code, hull.tons, hull.price))
	detailHull.Refresh()
}

func (h *hullProperties) updateHardPoints() {
	index := getIndexFromHull(hull.code)
	if index > -1 {
		hull.tons = hullSelections[index].tons
		hull.maxHP = hull.tons / 100
		detailMaxHardPoints.SetText(fmt.Sprintf("Maximum Hardpoints: %d", hull.maxHP))
		detailMaxHardPoints.Refresh()
	}
}

func (h *hullProperties) getTons() int {
	return hull.tons
}

func (h *hullProperties) getMCr() int {
	return hull.price
}
