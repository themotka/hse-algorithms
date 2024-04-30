package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"themotka/hse-algorithms/final-2nd-grade-lab-work/pkg"
	"time"
)

func main() {
	dots := generateDots(10000)
	file, err := os.Create("results/result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"rectangles", "BruteForce", "MapAlgorithm", "PersistentSegmentTree"})
	for i := 2; i < 5; i++ {
		quantity := int(math.Pow(2, float64(i)))
		testRec := generateRectangles(quantity)
		start0 := time.Now()
		res0 := pkg.BruteForce(testRec, dots)
		duration0 := time.Since(start0)
		start1 := time.Now()
		res1 := pkg.MapAlgorithm(testRec, dots)
		duration1 := time.Since(start1)
		start2 := time.Now()
		res2 := pkg.PersistentSegmentTree(testRec, dots)
		duration2 := time.Since(start2)
		if !slices.Equal(res0, res1) || !slices.Equal(res0, res2) {
			fmt.Print("damn")
		}
		fmt.Printf("%d,%d,%d\n", duration0.Nanoseconds(), duration1.Nanoseconds(), duration2.Nanoseconds())
		writer.Write([]string{strconv.Itoa(quantity), strconv.FormatInt(duration0.Nanoseconds(), 10), strconv.FormatInt(duration1.Nanoseconds(), 10), strconv.FormatInt(duration2.Nanoseconds(), 10)})
	}
}

func generateDots(n int) []pkg.DotCoordinates {
	coordinates := make([]pkg.DotCoordinates, n)
	for i := 0; i < n; i++ {
		x := int(math.Pow(float64(1000*i), 31)) % (20 * n)
		y := int(math.Pow(float64(1100*i), 31)) % (20 * n)
		coordinates = append(coordinates, pkg.DotCoordinates{
			X: x,
			Y: y,
		})
	}
	return coordinates
}

func generateRectangles(n int) []pkg.Rectangle {
	rectangles := make([]pkg.Rectangle, n)
	fir := pkg.DotCoordinates{}
	sec := pkg.DotCoordinates{}
	for i := 0; i < n; i++ {
		fir.X = 10 * i
		fir.Y = 10 * i
		sec.X = 10 * (2*n - 1)
		sec.Y = 10 * (2*n - 1)
		rectangles = append(rectangles, pkg.Rectangle{
			LeftLower:  fir,
			RightUpper: sec,
		})
	}
	return rectangles
}
