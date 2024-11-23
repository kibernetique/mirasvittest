package optimalblock1concurrent

import "sync"

type blockDistanceMap map[string]int

// This algorithm concurrently processes an array of blocks, calculating the distance
// from each block to the nearest infrastructure objects. It generates
// a separate array that stores the maximum distances to the nearest
// infrastructure objects, which is later used to identify the optimal block.
// To find the nearest object, the algorithm generates a "wave" around the
// block under consideration, and when the wave reaches the closest
// infrastructure object, the distance is recorded.
func GetOptimalBlock(blocks []map[string]bool) int {
	maxDistances := make([]int, len(blocks))
	var maxDist int
	var wg sync.WaitGroup
	wg.Add(len(blocks))
	for i := range blocks {
		go func(i int) {
			defer wg.Done()
			maxDist = getMaxDistanceToInfrastructure(getBlockDistancesToAllInfrastuctures(i, blocks))
			maxDistances[i] = maxDist
		}(i)
	}
	wg.Wait()
	return getBlockWithMinDistanceIndex(maxDistances)
}

// Returns distances to all infrastructure object
func getBlockDistancesToAllInfrastuctures(index int, blocks []map[string]bool) blockDistanceMap {
	blockDistance := blockDistanceMap{}
	for infrastructureKey := range blocks[index] {
		blockDistance[infrastructureKey] = getBlockToInfrastructureDistance(infrastructureKey, index, blocks)
	}
	return blockDistance
}

// Returns distance from given block to given nearest infrastructure
//   - key - key name of infrastructure
//   - index - index of given block
func getBlockToInfrastructureDistance(key string, index int, blocks []map[string]bool) int {
	maxIndex := max(index, len(blocks)-index)
	for i := 0; i < maxIndex; i++ {
		if index-i >= 0 && blocks[index-i][key] || index+i < len(blocks) && blocks[index+i][key] {
			return i
		}
	}
	return len(blocks) + 1
}

// Returns distance to farthest infrastructure from given block
func getMaxDistanceToInfrastructure(block blockDistanceMap) int {
	max := 0
	for _, infrastrucrureDistance := range block {
		if infrastrucrureDistance > max {
			max = infrastrucrureDistance
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
