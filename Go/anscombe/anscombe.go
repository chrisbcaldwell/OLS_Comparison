package main

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/montanaflynn/stats"
)

func main() {
	series1 := []stats.Coordinate{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}
	series2 := []stats.Coordinate{
		{X: 10, Y: 9.14},
		{X: 8, Y: 8.14},
		{X: 13, Y: 8.74},
		{X: 9, Y: 8.77},
		{X: 11, Y: 9.26},
		{X: 14, Y: 8.1},
		{X: 6, Y: 6.13},
		{X: 4, Y: 3.1},
		{X: 12, Y: 9.13},
		{X: 7, Y: 7.26},
		{X: 5, Y: 4.74},
	}
	series3 := []stats.Coordinate{
		{X: 10, Y: 7.46},
		{X: 8, Y: 6.77},
		{X: 13, Y: 12.74},
		{X: 9, Y: 7.11},
		{X: 11, Y: 7.81},
		{X: 14, Y: 8.84},
		{X: 6, Y: 6.08},
		{X: 4, Y: 5.39},
		{X: 12, Y: 8.15},
		{X: 7, Y: 6.42},
		{X: 5, Y: 5.73},
	}
	series4 := []stats.Coordinate{
		{X: 8, Y: 6.58},
		{X: 8, Y: 5.76},
		{X: 8, Y: 7.71},
		{X: 8, Y: 8.84},
		{X: 8, Y: 8.47},
		{X: 8, Y: 7.04},
		{X: 8, Y: 5.25},
		{X: 19, Y: 12.5},
		{X: 8, Y: 5.56},
		{X: 8, Y: 7.91},
		{X: 8, Y: 6.89},
	}
	start := time.Now()
	slope, intercept, _ := regression(series1)
	fmt.Println("Slope:", slope, "Intercept:", intercept)
	slope, intercept, _ = regression(series2)
	fmt.Println("Slope:", slope, "Intercept:", intercept)
	slope, intercept, _ = regression(series3)
	fmt.Println("Slope:", slope, "Intercept:", intercept)
	slope, intercept, _ = regression(series4)
	fmt.Println("Slope:", slope, "Intercept:", intercept)
	elapsed := time.Since(start)
	fmt.Println("Run time in Go (github.com/montanaflynn/stats package):")
	fmt.Println(elapsed)
}

// regression takes the best fit line coordinate pairs returned by stats.LinearRegression
// and returns the slope and intercept of the line
func regression(data stats.Series) (slope float64, intercept float64, err error) {
	r, err := stats.LinearRegression(data)
	if err != nil {
		return math.NaN(), math.NaN(), err
	}
	for i := 1; i < len(r); i++ {
		if r[0].X != r[i].X {
			slope = (r[i].Y - r[0].Y) / (r[i].X - r[0].X)
			intercept = r[0].Y - slope*r[0].X
			return slope, intercept, nil
		}

	}
	return math.NaN(), math.NaN(), errors.New("can not calculate slope if all x values are equal")
}
