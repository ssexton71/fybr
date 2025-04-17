package util

func AsNode(val any) (node map[string]any, ok bool) {
	node, ok = val.(map[string]any)
	return
}

func IsNode(val any) (ok bool) {
	_, ok = AsNode(val)
	return
}

// TODO: use https://github.com/AsaiYusuke/jsonpath instead?  it handles arrays too

func WalkNodes(node map[string]any, path ...string) any {
	for _, p := range path {
		val, ok := node[p]
		if !ok {
			return nil
		}
		node, ok = AsNode(val)
		if !ok {
			return val
		}
	}
	return node
}
