package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	inch         float64 = 1.0
	half                 = 1 / 2.0
	quarter              = 1 / 4.0
	eigth                = 1 / 8.0
	sixteenth            = 1 / 16.0
	thirtySecond         = 1 / 32.0
	extraSpace           = sixteenth
)

var sizes = []float64{inch, half, quarter, eigth, sixteenth, thirtySecond}
var names = []string{"1", "1/2", "1/4", "1/8", "1/16", "1/32"}

func main() {
	input := os.Args[1]
	cutSize, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input")
	}

	cutSize += extraSpace

	totals := getResults(sizes, cutSize)

	printResults(totals)
}

func calculateRings(remaingingSize, ringSize float64) (float64, int) {
	numberOfRings := int(remaingingSize / ringSize)
	rs := math.Mod(remaingingSize, ringSize)

	return rs, numberOfRings
}

func getResults(sizes []float64, cutSize float64) []int {
	totals := make([]int, len(sizes))

	for i, s := range sizes {
		cutSize, totals[i] = calculateRings(cutSize, s)
	}

	if cutSize > 0 {
		totals[len(totals)-1]++
	}

	return totals
}

func printResults(totals []int) {
	for i, t := range totals {
		fmt.Printf("%s:\t%d\n", names[i], t)
	}
}
