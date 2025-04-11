package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	input.MultiLine = true

	content := container.NewBorder(nil, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}), nil, nil, input)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.Size{Width: 320, Height: 240})
	myWindow.ShowAndRun()
}
