package optimalblock2

// This algorithm goes through the blocks, considering each type of infrastructure
// object separately. It first marks all zero distances (i.e., those blocks that
// contain infrastructure objects). Then, it marks the distances equal to 1 around
// the zero distances. After that, it marks the distances equal to 2 around the ones, and so on.
func GetOptimalBlock(blocks []map[string]bool) int {
	// get all infrastructure keys in a slice
	infrastructureKeys := make([]string, len(blocks[0]))
	i := 0
	for infKey := range blocks[0] {
		infrastructureKeys[i] = infKey
		i++
	}

	// count all distances for all infrastructure keys
	distances := make([]map[string]int, len(blocks))
	for k := range distances {
		distances[k] = make(map[string]int)
	}
	for _, infKey := range infrastructureKeys {
		countInfrastructureDistances(infKey, blocks, distances)
	}

	// count max dinstances for every infrastructure
	maxDistances := make([]int, len(distances))
	var maxDist int
	for i := range blocks {
		maxDist = getMaxDistanceToInfrastructure(distances[i])
		maxDistances[i] = maxDist
	}

	return getBlockWithMinDistanceIndex(maxDistances)
}

// Counts distances to objects of given infrastructure.
// Results are saving into a @distances parameter
func countInfrastructureDistances(infrastructureKey string, blocks []map[string]bool, distances []map[string]int) {
	// finding all distances which is qeual to 1
	for blockKey, block := range blocks {
		if block[infrastructureKey] {
			distances[blockKey][infrastructureKey] = 0
		}
	}

	// finding all distances which are more then 1
	var ok, dataChanged = false, true
	var dis int
	for i := 1; dataChanged; i++ {
		dataChanged = false
		for blockKey := 0; blockKey < len(distances); blockKey++ {
			if dis, ok = distances[blockKey][infrastructureKey]; ok && dis == i-1 {
				if blockKey > 0 {
					if _, ok = distances[blockKey-1][infrastructureKey]; !ok {
						distances[blockKey-1][infrastructureKey] = i
						dataChanged = true
					}
				}
				if blockKey < len(distances)-1 {
					if _, ok = distances[blockKey+1][infrastructureKey]; !ok {
						distances[blockKey+1][infrastructureKey] = i
						dataChanged = true
					}
				}
			}
		}
	}
}

// Returns distance to farthest infrastructure from given block
func getMaxDistanceToInfrastructure(blockDistances map[string]int) int {
	max := 0
	for _, infrastructureDistance := range blockDistances {
		if infrastructureDistance > max {
			max = infrastructureDistance
		}
	}
	return max
}

// Returns an index of min value from given slice
func getBlockWithMinDistanceIndex(maxDistances []int) int {
	min := len(maxDistances)
	var minIndex int
	for key, value := range maxDistances {
		if value < min {
			min = value
			minIndex = key
		}
	}
	return minIndex
}
