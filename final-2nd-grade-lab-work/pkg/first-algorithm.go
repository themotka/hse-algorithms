package pkg

func BruteForce(rectangles []Rectangle, coordinates []DotCoordinates) []int {
	result := make([]int, len(coordinates))
	for i := 0; i < len(coordinates); i++ {
		cnt := 0
		for j := 0; j < len(rectangles); j++ {
			if coordinates[i].X >= rectangles[j].LeftLower.X && rectangles[j].RightUpper.X > coordinates[i].X &&
				rectangles[j].LeftLower.Y <= coordinates[i].Y && coordinates[i].Y < rectangles[j].RightUpper.Y {
				cnt += 1
			}
		}
		result[i] = cnt
	}
	return result
}
