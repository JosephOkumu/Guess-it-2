package calcs

import (
	"math"
)

func PredictRange(inputData []float64)(int, int) {
	start := len(inputData) - 4

	if start < 0 {
		start = 0
	}

	numSlice := inputData[start:]
	m, b := LinearRegression(numSlice)
	newIndex := float64(len(numSlice))

	ymxc := m*newIndex + b
	vari := Variance(numSlice)

	stDev := StdDev(vari)

	min := ymxc - stDev*1.8
	max := ymxc + stDev*1.8

	return int(math.Round(min)), int(math.Round(max))	
}