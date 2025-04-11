package ui

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type textViewerToolbar struct {
	container *fyne.Container
	pathInput *widget.Entry
	openBtn   *widget.Button
	textInput *widget.Entry
}

func (tb *textViewerToolbar) initToolbar(textInput *widget.Entry) *textViewerToolbar {
	tb.pathInput = widget.NewEntry()
	tb.pathInput.OnChanged = func(s string) {
		if s > "" {
			tb.openBtn.Enable()
		} else {
			tb.openBtn.Disable()
		}
	}
	tb.openBtn = widget.NewButton("Open", func() { tb.onOpen() })
	tb.openBtn.Disable()
	tb.container = container.NewBorder(nil, nil, nil, tb.openBtn, tb.pathInput)
	tb.textInput = textInput
	return tb
}

func (tb *textViewerToolbar) onOpen() {
	data, err := os.ReadFile(tb.pathInput.Text)
	str := string(data)
	fmt.Printf("%v\n%v\n", err, str)
	// FIXME: this doesn't populate/change the contents of the textInput
	tb.textInput.Text = str
}

func NewTextViewerToolbar(textInput *widget.Entry) *textViewerToolbar {
	return (&textViewerToolbar{}).initToolbar(textInput)
}

type textViewer struct {
	Content   *fyne.Container
	toolbar   *textViewerToolbar
	textInput *widget.Entry
}

func (tv *textViewer) initTextViewer() *textViewer {
	tv.textInput = widget.NewMultiLineEntry()
	tv.toolbar = NewTextViewerToolbar(tv.textInput)
	tv.Content = container.NewBorder(tv.toolbar.container, nil, nil, nil, tv.textInput)
	return tv
}

func NewTextViewer() *textViewer {
	return (&textViewer{}).initTextViewer()
}
