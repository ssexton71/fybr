package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssexton71/fybr/util"
)

type viewerToolbar struct {
	container *fyne.Container
	pathInput *widget.Entry
	openBtn   *widget.Button
	lblStatus *widget.Label
	setData   func([]byte)
}

func (tb *viewerToolbar) initViewerToolbar(setData func([]byte), lblStatus *widget.Label) *viewerToolbar {
	tb.setData = setData
	tb.pathInput = widget.NewEntry()
	tb.pathInput.SetText("http://scooby:doo@10.0.0.168:8080/api/v1/servers")
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
	tb.lblStatus = lblStatus
	return tb
}

func (tb *viewerToolbar) afterOpen(data []byte, err error) {
	status := "ok"
	if err != nil {
		status = "error: " + err.Error()
	}
	if data != nil {
		status += fmt.Sprintf(" (%d bytes)", len(data))
	} else {
		status += " (no data)"
	}
	tb.setData(data)
	fyne.Do(func() {
		tb.lblStatus.SetText(status)
	})

}

func (tb *viewerToolbar) onOpen() {
	tb.lblStatus.SetText("reading...")
	go func() {
		path := util.Path{Path: tb.pathInput.Text,
			Progress: func(n int) {
				fyne.Do(func() {
					tb.lblStatus.SetText(fmt.Sprintf("reading (%d bytes)...", n))
				})
			}}
		data, err := path.ReadData()
		tb.afterOpen(data, err)
	}()
}

func NewViewerToolbar(setData func([]byte), lblStatus *widget.Label) *viewerToolbar {
	return (&viewerToolbar{}).initViewerToolbar(setData, lblStatus)
}
