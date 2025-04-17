package ui

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/clbanning/mxj/v2"
	"github.com/ssexton71/fybr/util"
)

type treeViewer struct {
	Content   *fyne.Container
	toolbar   *viewerToolbar
	treeView  *widget.Tree
	treeModel map[string]any
	lblStatus *widget.Label
}

func (tv *treeViewer) ChildUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	uids := []string{}
	if id == "" {
		uids = slices.AppendSeq(uids, maps.Keys(tv.treeModel))
	} else {
		val := util.WalkNodes(tv.treeModel, strings.Split(id, ".")...)
		m, ok := util.AsNode(val)
		if !ok {
			return []string{}
		}
		for k := range maps.Keys(m) {
			uids = append(uids, id+"."+k)
		}
	}
	slices.Sort(uids)
	return uids
}

func (tv *treeViewer) IsBranch(id widget.TreeNodeID) bool {
	return len(tv.ChildUIDs(id)) > 0
}

func (tv *treeViewer) CreateNode(branch bool) fyne.CanvasObject {
	if branch {
		return widget.NewLabel("Branch template")
	}
	return widget.NewLabel("Leaf template")
}

func (tv *treeViewer) UpdateNode(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	text := id
	if branch {
		text += " (branch)"
	} else {
		text += ": "
		text += util.WalkNodes(tv.treeModel, strings.Split(id, ".")...).(string)

	}
	o.(*widget.Label).SetText(text)
}

func (tv *treeViewer) initTreeViewer() *treeViewer {
	tv.treeView = widget.NewTree(tv.ChildUIDs,
		tv.IsBranch, tv.CreateNode, tv.UpdateNode)
	tv.treeModel = map[string]any{
		"a": "b",
		"c": "d",
	}
	tv.treeModel["e"] = map[string]any{
		"f": "g",
	}

	tv.lblStatus = widget.NewLabel("ready")
	tv.toolbar = NewViewerToolbar(
		func(data []byte) {
			m, err := mxj.NewMapXml(data)
			fmt.Printf("%v %v", m, err)
			tv.treeModel = m
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
