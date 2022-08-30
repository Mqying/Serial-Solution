package util

import (
	"fmt"
	"strconv"
)

func FixPrecision(value float64) float64 {
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return result
}