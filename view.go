package main

import "fyne.io/fyne/widget"

var (
	saveButton = widget.NewButton("Save", saveMe)
	loadButton = widget.NewButton("Load", loadMe)
)

var (
	settings *widget.Form
	details  *widget.Box
)
