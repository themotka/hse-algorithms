package main

import (
	"sort"
)

func MapAlgorithm(rectangles []Rectangle, coordinates []DotCoordinates) []int {
	result := make([]int, len(coordinates))
	abscissa := make(map[int]struct{})
	ordinate := make(map[int]struct{})
	for _, rectangle := range rectangles {
		abscissa[rectangle.leftLower.x] = struct{}{}
		ordinate[rectangle.leftLower.y] = struct{}{}

		abscissa[rectangle.rightUpper.x] = struct{}{}
		ordinate[rectangle.rightUpper.y] = struct{}{}
	}
	zipX := make([]int, 0, len(abscissa))
	zipY := make([]int, 0, len(ordinate))
	for x := range abscissa {
		zipX = append(zipX, x)
	}
	for y := range ordinate {
		zipY = append(zipY, y)
	}
	sort.Ints(zipX)
	sort.Ints(zipY)
	mapCoords := make([][]int, len(zipX))
	for i := range mapCoords {
		mapCoords[i] = make([]int, len(zipY))
	}
	for _, rectangle := range rectangles {
		startX := bound(zipX, rectangle.leftLower.x)
		finX := bound(zipY, rectangle.rightUpper.x)

		startY := bound(zipY, rectangle.leftLower.y)
		finY := bound(zipY, rectangle.rightUpper.y)

		for x := startX; x < finX; x++ {
			for y := startY; y < finY; y++ {
				mapCoords[x][y]++
			}
		}
	}
	for i, point := range coordinates {
		x := bound(zipX, point.x)
		if x == len(zipX) {
			result[i] = 0
			continue
		}
		if point.x < zipX[x] {
			x--
			if x < 0 {
				result[i] = 0
				continue
			}
		}

		y := bound(zipY, point.y)
		if y == len(zipY) {
			result[i] = 0
			continue
		}
		if point.y < zipY[y] {
			y--
			if y < 0 {
				result[i] = 0
				continue
			}
		}

		result[i] = mapCoords[x][y]
	}

	return result
}

func bound(slice []int, val int) int {
	return sort.Search(len(slice), func(i int) bool { return slice[i] >= val })
}
