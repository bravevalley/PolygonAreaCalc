package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// TODO:  Fix stuff

// Store the cordinates
type cordinates struct {
	x int
	y int
}
var (

	// Regexp to capture the int
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
)

func computeArea(input string)  {
	// [[(110,0) 110 0] [(108,0) 108 0] [(77,1) 77 1] [(72,2) 72 2]]
	var points = []cordinates{}

	var area float64

	for _, data := range r.FindAllStringSubmatch(input, -1) {

		xAxis, _ := strconv.Atoi(data[1])
		yAxis, _ := strconv.Atoi(data[2])

		// Structure I want [{x, y}, {x, y}, {x, y}]
		points = append(points, cordinates{x:xAxis, y:yAxis,})	
	}
	
	for i := 0; i < len(points); i++ {
		a, b := points[i], points[(i + 1)%len(points)]
		area += float64(a.x * b.y) - float64(a.y * b.x)
	}
	
	
	fmt.Println(math.Abs(area) / 2.0)
	

}

// (110,0),(108,0),(77,1),(72,2),(68,2),(58,5)

func main() {

	// Import the data
	input := "(4,10),(12,8),(10,3),(2,2),(7,5)"

	computeArea(input)

	// Extract the data

	// Trap the data into a file structure for easier computation

	// Compute the Area

	// Return the area

}
