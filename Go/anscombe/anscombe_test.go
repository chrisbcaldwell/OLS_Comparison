package main

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestRegression(t *testing.T) {
	// testSeries has regression slope = 3, intercept = 0
	testSeries := []stats.Coordinate{
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 4},
	}
	exSlope := 3.0
	exIntercept := 0.0
	tolerance := 0.00001
	slope, intercept, err := regression(testSeries)
	if err != nil {
		t.Error("Error performing test regression")
	}
	if math.Abs(slope-exSlope) > tolerance {
		t.Log("Expected slope:", exSlope)
		t.Log("Result slope:", slope)
		t.Error("Test regression slope does not match expected")
	}
	if math.Abs(intercept-exIntercept) > tolerance {
		t.Log("Expected intercept:", exIntercept)
		t.Log("Result intercept:", intercept)
		t.Error("Test regression intercept does not match expected")
	}
}
