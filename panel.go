package main

import "fyne.io/fyne/widget"

type update func()

type panel struct {
	change   update
	selects  []*widget.Select
	settings []*widget.Form
	details  []*widget.Box
}
