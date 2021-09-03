package main

import (
	"fmt"
	"math"
)

func main() {
	Data := []float64{-20, -25.0, -27.0, -21.0, 10, 13.0, 19.0, 15.5, 20, 24.5, 34.2}
	result := make(map[float64][]float64)

	for i := 0; i < len(Data); i++ {
		value := math.Round(Data[i])
		valueShift := int(value) % 10
		result[value-float64(valueShift)] = append(result[value-float64(valueShift)], Data[i])
	}
	fmt.Println(result)
}
