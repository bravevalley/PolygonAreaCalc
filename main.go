package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Store the cordinates
type cordinates struct {
	x int
	y int
}

var (

	// Regexp to capture the int
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)

	// Wait Group
	wg = sync.WaitGroup{}
)

func computeArea(input <-chan string) {

	// Import from the data channel
	for v := range input {

		// [[(110,0) 110 0] [(108,0) 108 0] [(77,1) 77 1] [(72,2) 72 2]]
		var points = []cordinates{}

		var area float64

		// Loop over the slice of string returned by the RegExp
		for _, data := range r.FindAllStringSubmatch(v, -1) {

			// Convert the string to interger
			xAxis, _ := strconv.Atoi(data[1])
			yAxis, _ := strconv.Atoi(data[2])

			// Structure I want [{x, y}, {x, y}, {x, y}]
			points = append(points, cordinates{x: xAxis, y: yAxis})
		}

		// Loop through the number elements from the last process
		for i := 0; i < len(points); i++ {

			// Shoelace algorithm
			// [a    b] \
			// [c    d]  > len() = 3 line66
			// [e    f] /
			// ((ad) + (cf) + (eb)) - ((bc) + (de) + (fa))
			// Absolute value / 2

			// points[(i + 1)%len(points)] - this is to pass in the first element in the slice
			// (2 + 1) % len()
			// 3 % 3 = 0
			// points[0] = Fiest element
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}

		// Print result
		fmt.Println(math.Abs(area) / 2.0)

	}
	wg.Done()

}

func main() {

	start := time.Now()

	wg.Add(1)
	dataChannel := make(chan string, 100)
	for i := 1; i < runtime.NumCPU(); i++ {
		go computeArea(dataChannel)

	}
	// Import the data
	filePath, _ := filepath.Abs("./polygon.txt")

	// (110,0),(108,0),(77,1),(72,2),(68,2),(58,5)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Unable to open file")
	}

	// Convert slice of byte to string
	inputStream := string(file)

	// Please the string into lines
	data := strings.Split(inputStream, "\n")

	// Feed each to the datachannel which would be queued on the buffer chan
	for _, v := range data {
		dataChannel <- v
	}
	close(dataChannel)

	wg.Wait()

	end := time.Since(start)
	fmt.Println(end)

}
