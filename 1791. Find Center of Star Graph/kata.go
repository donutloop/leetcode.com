package _791__Find_Center_of_Star_Graph

func findCenter(edges [][]int) int {

	edgesCounter := make(map[int]int)

	for _, nodePair := range edges {
		edgesCounter[nodePair[0]]++
		edgesCounter[nodePair[1]]++

		if edgesCounter[nodePair[1]] == len(edges) {
			return nodePair[1]
		}

		if edgesCounter[nodePair[0]] == len(edges) {
			return nodePair[0]
		}
	}

	return -1
}
