package main

import (
	"mirasvittest/optimalblock1"
	"mirasvittest/optimalblock1concurrent"
	"mirasvittest/optimalblock1opt"
	"mirasvittest/optimalblock2"
	"mirasvittest/optimalblock2concurrent"
	"mirasvittest/optimalblock2opt"
	"testing"
)

func BenchmarkOptimalBlock1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1.GetOptimalBlock(testBlocks)
	}
}

func BenchmarkOptimalBlock1Opt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1opt.GetOptimalBlock(testBlocks)
	}
}

func BenchmarkOptimalBlock1Concurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1concurrent.GetOptimalBlock(testBlocks)
	}
}

func BenchmarkOptimalBlock2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2.GetOptimalBlock(testBlocks)
	}
}

func BenchmarkOptimalBlock2Opt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2opt.GetOptimalBlock(testBlocks)
	}
}

func BenchmarkOptimalBlock2Concurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2concurrent.GetOptimalBlock(testBlocks)
	}
}

func BenchmarkOptimalBlock1LongArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1.GetOptimalBlock(testBlocksLong)
	}
}

func BenchmarkOptimalBlock1LongArrayOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1opt.GetOptimalBlock(testBlocksLong)
	}
}

func BenchmarkOptimalBlock1LongArrayConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1concurrent.GetOptimalBlock(testBlocksLong)
	}
}

func BenchmarkOptimalBlock2LongArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2.GetOptimalBlock(testBlocksLong)
	}
}

func BenchmarkOptimalBlock2LongArrayOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2opt.GetOptimalBlock(testBlocksLong)
	}
}

func BenchmarkOptimalBlock2LongArrayConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2concurrent.GetOptimalBlock(testBlocksLong)
	}
}

func BenchmarkOptimalBlock1WithZeroDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1.GetOptimalBlock(testBlocksWithZeroDistance)
	}
}

func BenchmarkOptimalBlock1WithZeroDistanceOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1opt.GetOptimalBlock(testBlocksWithZeroDistance)
	}
}

func BenchmarkOptimalBlock1WithZeroDistanceConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock1concurrent.GetOptimalBlock(testBlocksWithZeroDistance)
	}
}

func BenchmarkOptimalBlock2WithZeroDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2.GetOptimalBlock(testBlocksWithZeroDistance)
	}
}

func BenchmarkOptimalBlock2WithZeroDistanceOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2opt.GetOptimalBlock(testBlocksWithZeroDistance)
	}
}

func BenchmarkOptimalBlock2WithZeroDistanceConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = optimalblock2concurrent.GetOptimalBlock(testBlocksWithZeroDistance)
	}
}
