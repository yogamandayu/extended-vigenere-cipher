package util

import (
	"math/rand"
)

func TwoDimensionArrayRandomInteger(seed, n int) [][]int {
	var arr [][]int
	rand.Seed(int64(seed))
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Perm(n))
	}
	return arr
}
