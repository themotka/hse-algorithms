package pkg

import (
	"sort"
)

func MapAlgorithm(rectangles []Rectangle, coordinates []DotCoordinates) []int {
	result := make([]int, len(coordinates))
	abscissa := make(map[int]struct{})
	ordinate := make(map[int]struct{})
	for _, rectangle := range rectangles {
		abscissa[rectangle.LeftLower.X] = struct{}{}
		ordinate[rectangle.LeftLower.Y] = struct{}{}

		abscissa[rectangle.RightUpper.X] = struct{}{}
		ordinate[rectangle.RightUpper.Y] = struct{}{}
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
		startX := bound(zipX, rectangle.LeftLower.X)
		finX := bound(zipY, rectangle.RightUpper.X)
		startY := bound(zipY, rectangle.LeftLower.Y)
		finY := bound(zipY, rectangle.RightUpper.Y)

		for x := startX; x < finX; x++ {
			for y := startY; y < finY; y++ {
				mapCoords[x][y]++
			}
		}
	}
	for i, point := range coordinates {
		x := bound(zipX, point.X)
		if x == len(zipX) {
			result[i] = 0
			continue
		}
		if point.X < zipX[x] {
			x--
			if x < 0 {
				result[i] = 0
				continue
			}
		}
		y := bound(zipY, point.Y)
		if y == len(zipY) {
			result[i] = 0
			continue
		}
		if point.Y < zipY[y] {
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
