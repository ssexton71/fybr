package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type textViewer struct {
	Content   *fyne.Container
	toolbar   *viewerToolbar
	textView  *widget.Entry
	lblStatus *widget.Label
}

func (tv *textViewer) initTextViewer() *textViewer {
	tv.textView = widget.NewMultiLineEntry()
	tv.lblStatus = widget.NewLabel("ready")
	tv.toolbar = NewViewerToolbar(
		func(data []byte) {
			fyne.Do(func() {
				if data == nil {
					tv.textView.SetText("")
				} else {
					tv.textView.SetText(string(data))
				}
			})
		}, tv.lblStatus)
	tv.Content = container.NewBorder(tv.toolbar.container, tv.lblStatus, nil, nil, tv.textView)
	return tv
}

func (tv *textViewer) SetFocus(canvas fyne.Canvas) {
	canvas.Focus(tv.toolbar.pathInput)
}

func NewTextViewer() *textViewer {
	return (&textViewer{}).initTextViewer()
}
