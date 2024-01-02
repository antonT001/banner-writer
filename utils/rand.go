package utils

import (
	"math/rand"
	"time"
)

var count = rand.New(rand.NewSource(time.Now().Unix()))

// RandInt returns, as an int, a non-negative number in the interval [n,k). It panics if n <= k.
func RandInt(n, k int) int {
	return count.Intn(k-n) + n
}
