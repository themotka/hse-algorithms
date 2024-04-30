package main

func BruteForce(rectangles []Rectangle, coordinates []DotCoordinates) []int {
	result := make([]int, len(coordinates))
	for i := 0; i < len(coordinates); i++ {
		cnt := 0
		for j := 0; j < len(rectangles); j++ {
			if coordinates[i].x >= rectangles[j].leftLower.x && rectangles[j].rightUpper.x > coordinates[i].x &&
				rectangles[j].leftLower.y <= coordinates[i].y && coordinates[i].y < rectangles[j].rightUpper.y {
				cnt += 1
			}
		}
		result[i] = cnt
	}
	return result
}
