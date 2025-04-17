package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"maps"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/AsaiYusuke/jsonpath"
	xj "github.com/basgys/goxml2json"
)

type treeViewer struct {
	Content   *fyne.Container
	toolbar   *viewerToolbar
	treeView  *widget.Tree
	treeModel any
	lblStatus *widget.Label
}

func (tv *treeViewer) ChildUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	uids := []string{}
	if id == "" {
		id = "$"
	}
	items, _ := jsonpath.Retrieve(id, tv.treeModel)
	for _, item := range items {
		m, ok := item.(map[string]any)
		if ok {
			for k := range maps.Keys(m) {
				uids = append(uids, fmt.Sprintf("%s.%s", id, k))
			}
		}
		a, ok := item.([]any)
		if ok {
			for idx := range a {
				uids = append(uids, fmt.Sprintf("%s[%d]", id, idx))
			}
		}
	}
	slices.Sort(uids)
	return uids
}

func (tv *treeViewer) IsBranch(id widget.TreeNodeID) bool {
	return len(tv.ChildUIDs(id)) > 0
}

func (tv *treeViewer) CreateNode(branch bool) fyne.CanvasObject {
	return widget.NewLabel("")
}

func (tv *treeViewer) UpdateNode(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	text := id
	if branch {
		text += " (branch)"
	} else {
		text += ": "
		items, _ := jsonpath.Retrieve(id, tv.treeModel)
		for _, item := range items {
			str, ok := item.(string)
			if ok {
				text += str
			}
		}
	}
	o.(*widget.Label).SetText(text)
}

func (tv *treeViewer) initTreeViewer() *treeViewer {
	tv.treeView = widget.NewTree(tv.ChildUIDs,
		tv.IsBranch, tv.CreateNode, tv.UpdateNode)
	top := map[string]any{}
	tv.treeModel = top

	tv.lblStatus = widget.NewLabel("ready")
	tv.toolbar = NewViewerToolbar(
		func(data []byte) {
			if len(data) > 0 && data[0] == '<' {
				buf, _ := xj.Convert(bytes.NewReader(data))
				data = buf.Bytes()
			}
			json.Unmarshal(data, &tv.treeModel)
			fyne.Do(func() {
				tv.treeView.Refresh()
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
