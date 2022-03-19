package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

// drive structure defines what we need to know about a given drive (A-Z) selection.
type totals struct {
	tonnage float32 // tons used by all selected items
	cost    float32 // cost of all selected items in MCr
	cargo   float32 // Cargo capacity: tons remaining after using up tons listed above
}

var (
	summary = totals{
		tonnage: 83.0,
		cost:    98.0,
		cargo:   117.0,
	}
	detailSummary     = widget.NewLabel("Summary")
	summaryDetailsBox = widget.NewVBox(detailSummary)
)

func (t *totals) init(form *widget.Form, box *widget.Box) {
	box.Children = append(box.Children, summaryDetailsBox)
}

func (t *totals) update() {
	t.tonnage = float32(drives.tons()) + berths.tons() + float32(vehicles.tons()) + float32(weapons.tons())
	t.cost = float32(hull.price) +
		float32(drives.mCr()) + float32(berths.mCr()) + float32(vehicles.mCr()) + float32(weapons.cost())
	t.cargo = float32(hull.tons) - t.tonnage
	detailSummary.SetText(fmt.Sprintf(
		"Total Tons used: %.2f, cost: %.2fMcr, Cargo Space %.2f tons",
		t.tonnage, t.cost, t.cargo,
	))
	detailSummary.Refresh()
}
