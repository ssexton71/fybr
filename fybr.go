package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/ssexton71/fybr/ui"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("FyBr")

	//	tv := ui.NewTextViewer()
	tv := ui.NewTreeViewer()
	myWindow.SetContent(tv.Content)
	myWindow.Resize(fyne.Size{Width: 800, Height: 600})
	myWindow.CenterOnScreen()
	tv.SetFocus(myWindow.Canvas())
	myWindow.ShowAndRun()
}
