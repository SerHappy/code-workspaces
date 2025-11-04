package workspaces

func BuildIndexByRelPath(ws []Workspace) map[string]Workspace {
	index := make(map[string]Workspace, len(ws))
	for _, ws := range ws {
		index[ws.RelDir] = ws
	}

	return index
}

func Keys(ws []Workspace) []string {
	keys := make([]string, 0, len(ws))
	for _, ws := range ws {
		keys = append(keys, ws.RelDir)
	}
	
	return keys
}
