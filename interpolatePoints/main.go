package main

import (
	"fmt"
)

// #    145┃                   o
// #    140┃
// #    135┃                c
// #    130┃
// #    125┃             b
// #    120┃
// # Value ┃          o
// #    110┃       a
// #    105┃    o
// #    100┃ o
// #       ┗━━━━━━━━━━━━━━━━━━━━
// #         0  5 10 15 20 25 30
// #         Time

// # for (x1, y1), (x2, y2) with x1 < x < x2
// # slope = (y2 - y1) / (x2 - x1)
// # y = y1 + slope * (x - x1)

type point struct {
	timestamp int
	value     int
}

var inputPoints = []point{{0, 100}, {5, 105}, {15, 115}, {30, 145}}

func interpolateTwoPoints(p1 point, p2 point, step int) []point {

	slope := (p2.value - p1.value) / (p2.timestamp - p1.timestamp)
	yIntercept := p2.value - slope*p2.timestamp

	var twoPointsInterpolate []point
	for x := p1.timestamp; x < p2.timestamp; x += step {
		y := slope*x + yIntercept
		twoPointsInterpolate = append(twoPointsInterpolate, point{x, y})
	}

	return twoPointsInterpolate
}

func interpolate(input []point, step int) []point {

	var interpolatedList []point
	var length = len(input)

	if length == 0 || length == 1 {
		return input
	}

	for i := 0; i < length-1; i++ {
		interpolatedList = append(interpolatedList, interpolateTwoPoints(input[i], input[i+1], step)...)
	}
	interpolatedList = append(interpolatedList, input[length-1])

	return interpolatedList
}

func main() {
	var interpolatedList = interpolate(inputPoints, 5)

	fmt.Printf("%v\n", inputPoints)
	fmt.Printf("%v\n", interpolatedList)
	fmt.Printf("%v\n", []point{{0, 100}, {5, 105}, {10, 110}, {15, 115}, {20, 125}, {25, 135}, {30, 145}})
}
