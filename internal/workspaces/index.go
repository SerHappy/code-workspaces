package workspaces

func BuildIndexByRelPath(wsList []Workspace) map[string]Workspace {
	index := make(map[string]Workspace, len(wsList))
	for _, ws := range wsList {
		index[ws.RelDir] = ws
	}

	return index
}

func Keys(wsList []Workspace) []string {
	keys := make([]string, 0, len(wsList))
	for _, ws := range wsList {
		keys = append(keys, ws.RelDir)
	}

	return keys
}
