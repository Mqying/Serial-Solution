package util

import (
	"math"
	"math/rand"
)

func Round(f float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Trunc((float64(f) + 0.5 / pow10N) * pow10N) / pow10N

	return r
}

//Generate N digits number by 10^(n)
func GenerateInt(x int) int {
	randNum := rand.Intn((9 * x) - 1) + x
	return randNum
}
