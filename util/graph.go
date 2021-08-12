package util

func ReverseTopologySort(graph map[string][]string, start string) []string {
	visited := make(map[string]bool)
	return recWalk(graph, visited, start)
}

// find parallel struct -> A depends on B and B depends on A
func recWalk(grpah map[string][]string, visited map[string]bool, start string) []string {
	result := make([]string, 0, len(grpah))

	visited[start] = true
	for _, node := range grpah[start] {
		if !visited[node] {
			result = append(result, recWalk(grpah, visited, node)...)
		}
	}

	result = append(result, start)
	return result
}
