package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssexton71/fybr/util"
)

type textViewerToolbar struct {
	container *fyne.Container
	pathInput *widget.Entry
	openBtn   *widget.Button
	textInput *widget.Entry
	lblStatus *widget.Label
}

func (tb *textViewerToolbar) initToolbar(textInput *widget.Entry, lblStatus *widget.Label) *textViewerToolbar {
	tb.pathInput = widget.NewEntry()
	tb.pathInput.OnChanged = func(s string) {
		if s > "" {
			tb.openBtn.Enable()
		} else {
			tb.openBtn.Disable()
		}
	}
	tb.pathInput.OnSubmitted = func(s string) { tb.onOpen() }
	tb.openBtn = widget.NewButton("Open", func() { tb.onOpen() })
	tb.openBtn.Disable()
	tb.container = container.NewBorder(nil, nil, nil, tb.openBtn, tb.pathInput)
	tb.textInput = textInput
	tb.lblStatus = lblStatus
	return tb
}

func (tb *textViewerToolbar) afterOpen(data []byte, err error) {
	status := "ok"
	if err != nil {
		status = "error: " + err.Error()
	}
	if data != nil {
		status += fmt.Sprintf(" (%d bytes)", len(data))
		tb.textInput.SetText(string(data))
	} else {
		status += " (no data)"
		tb.textInput.SetText("")
	}
	tb.lblStatus.SetText(status)
}

func (tb *textViewerToolbar) onOpen() {
	tb.lblStatus.SetText("reading...")
	go func() {
		path := util.Path{Path: tb.pathInput.Text,
			Progress: func(n int) {
				fyne.Do(func() {
					tb.lblStatus.SetText(fmt.Sprintf("reading (%d bytes)...", n))
				})
			}}
		data, err := path.ReadData()
		fyne.Do(func() {
			tb.afterOpen(data, err)
		})
	}()
}

func NewTextViewerToolbar(textInput *widget.Entry, lblStatus *widget.Label) *textViewerToolbar {
	return (&textViewerToolbar{}).initToolbar(textInput, lblStatus)
}

type textViewer struct {
	Content   *fyne.Container
	toolbar   *textViewerToolbar
	textInput *widget.Entry
	lblStatus *widget.Label
}

func (tv *textViewer) initTextViewer() *textViewer {
	tv.textInput = widget.NewMultiLineEntry()
	tv.lblStatus = widget.NewLabel("ready")
	tv.toolbar = NewTextViewerToolbar(tv.textInput, tv.lblStatus)
	tv.Content = container.NewBorder(tv.toolbar.container, tv.lblStatus, nil, nil, tv.textInput)
	return tv
}

func (tv *textViewer) SetFocus(canvas fyne.Canvas) {
	canvas.Focus(tv.toolbar.pathInput)
}

func NewTextViewer() *textViewer {
	return (&textViewer{}).initTextViewer()
}
