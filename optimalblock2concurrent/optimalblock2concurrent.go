package optimalblock2concurrent

import (
	"sync"
)

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
	infrastructureDistances := make([][]int, len(infrastructureKeys))
	var wg sync.WaitGroup
	wg.Add(len(infrastructureKeys))
	for infKeyIndex, infKey := range infrastructureKeys {
		go func(infKey string, infKeyIndex int) {
			infrastructureDistances[infKeyIndex] = getInfrastructureDistances(infKey, blocks)
			wg.Done()
		}(infKey, infKeyIndex)
	}
	wg.Wait()

	// count max dinstances for every infrastructure
	maxDistances := make([]int, len(blocks))
	var maxDist int
	for i := range blocks {
		maxDist = getMaxDistanceToInfrastructure(i, infrastructureDistances)
		maxDistances[i] = maxDist
	}

	return getBlockWithMinDistanceIndex(maxDistances)
}

func getInfrastructureDistances(infrastructureKey string, blocks []map[string]bool) []int {
	distances := make([]int, len(blocks))
	for i := range distances {
		distances[i] = -1
	}
	// finding all distances which is qeual to 0
	for blockKey, block := range blocks {
		if block[infrastructureKey] {
			distances[blockKey] = 0
		}
	}

	// finding all distances which are more then 0
	var dataChanged = true
	for i := 1; dataChanged; i++ {
		dataChanged = false
		for blockKey := 0; blockKey < len(distances); blockKey++ {
			if distances[blockKey] == i-1 {
				if blockKey > 0 && distances[blockKey-1] < 0 {
					distances[blockKey-1] = i
					dataChanged = true
				}
				if blockKey < len(distances)-1 && distances[blockKey+1] < 0 {
					distances[blockKey+1] = i
					dataChanged = true
				}
			}
		}
	}

	return distances
}

// Returns distance to farthest infrastructure from given block
func getMaxDistanceToInfrastructure(blockIndex int, infDistances [][]int) int {
	max := 0
	for infKey := range infDistances {
		if infDistances[infKey][blockIndex] > max {
			max = infDistances[infKey][blockIndex]
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
