package main

import (
	"fmt"
	"runtime"
	"time"

	"mirasvittest/optimalblock1"
	"mirasvittest/optimalblock1concurrent"
	"mirasvittest/optimalblock1opt"
	"mirasvittest/optimalblock2"
	"mirasvittest/optimalblock2concurrent"
	"mirasvittest/optimalblock2opt"
)

type testConditions struct {
	data          []map[string]bool
	description   string
	iterationsNum int
}
type testFunction struct {
	function    func([]map[string]bool) int
	description string
}

func main() {
	testsConditions := []testConditions{
		{
			data:          testBlocks,
			description:   "With small array of blocks",
			iterationsNum: 100000,
		},
		{
			data:          testBlocksLong,
			description:   "With long array of blocks",
			iterationsNum: 1000,
		},
		{
			data:          testBlocksWithZeroDistance,
			description:   "With long array of blocks with zero distance block",
			iterationsNum: 1000,
		},
	}

	testFunctions := []testFunction{
		{
			function:    optimalblock1.GetOptimalBlock,
			description: "optimalblock1",
		},
		{
			function:    optimalblock1opt.GetOptimalBlock,
			description: "optimalblock1opt",
		},
		{
			function:    optimalblock1concurrent.GetOptimalBlock,
			description: "optimalblock1concurrent",
		},
		{
			function:    optimalblock2.GetOptimalBlock,
			description: "optimalblock2",
		},
		{
			function:    optimalblock2opt.GetOptimalBlock,
			description: "optimalblock2opt",
		},
		{
			function:    optimalblock2concurrent.GetOptimalBlock,
			description: "optimalblock2concurrent",
		},
	}

	fmt.Printf("runtime.GOMAXPROCS = %v\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU:%v\n", runtime.NumCPU())

	var optimalBlock int
	var start time.Time
	var elapsed time.Duration

	for _, testCond := range testsConditions {
		fmt.Printf("\nTest details: %s\n", testCond.description)
		fmt.Printf("Number of iterations: %d, Number of blocks: %d\n", testCond.iterationsNum, len(testCond.data))

		for _, testFunc := range testFunctions {
			start = time.Now()
			for i := 0; i < testCond.iterationsNum; i++ {
				optimalBlock = testFunc.function(testCond.data)
			}
			elapsed = time.Since(start)
			fmt.Printf("%s took %s\t\t\tReturned value: %d\n", testFunc.description, elapsed, optimalBlock)
		}
	}
}
