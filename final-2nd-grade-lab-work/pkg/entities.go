package pkg

type DotCoordinates struct {
	X, Y int
}

type Rectangle struct {
	LeftLower, RightUpper DotCoordinates
}

type RectBorder struct {
	x, begY, endY, border int
}

type TreeNode struct {
	val, lInd, rInd int
	l, r            *TreeNode
}
