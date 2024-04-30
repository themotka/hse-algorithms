package main

type DotCoordinates struct {
	x, y int
}

type Rectangle struct {
	leftLower, rightUpper DotCoordinates
}

type RectBorder struct {
	x, begY, endY, border int
}

type TreeNode struct {
	val, lInd, rInd int
	l, r            *TreeNode
}
