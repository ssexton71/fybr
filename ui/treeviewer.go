package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/clbanning/mxj/v2"
)

type treeViewer struct {
	Content   *fyne.Container
	toolbar   *viewerToolbar
	treeView  *widget.Tree
	lblStatus *widget.Label
}

func (tv *treeViewer) initTreeViewer() *treeViewer {
	tv.treeView = widget.NewTreeWithStrings(nil)
	tv.lblStatus = widget.NewLabel("ready")
	tv.toolbar = NewViewerToolbar(
		func(data []byte) {
			fyne.Do(func() {
				m, err := mxj.NewMapXml(data)
				fmt.Printf("%v %v", m, err)

			})
		}, tv.lblStatus)
	tv.Content = container.NewBorder(tv.toolbar.container, tv.lblStatus, nil, nil, tv.treeView)
	return tv
}

func (tv *treeViewer) SetFocus(canvas fyne.Canvas) {
	canvas.Focus(tv.toolbar.pathInput)
}

func NewTreeViewer() *treeViewer {
	return (&treeViewer{}).initTreeViewer()
}
