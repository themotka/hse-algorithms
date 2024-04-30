package pkg

import (
	"sort"
)

func PersistentSegmentTree(rectangles []Rectangle, coordinates []DotCoordinates) []int {
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
	nodes := segmentTree(rectangles, zipY)
	ans := make([]int, len(coordinates))
	for i, cord := range coordinates {
		indX := search(zipX, cord.X)
		indY := search(zipY, cord.Y)
		if indX >= 0 && indY >= 0 {
			ans[i] = count(nodes[indX], indY)
		}
	}
	return ans
}

func binaryTree(lInd, rInd int) *TreeNode {
	if lInd >= rInd {
		return &TreeNode{lInd: lInd, rInd: rInd}
	}
	middle := (lInd + rInd) / 2
	left := binaryTree(lInd, middle)
	right := binaryTree(middle+1, rInd)
	return &TreeNode{lInd: lInd, rInd: rInd, l: left, r: right}
}

func segmentTree(rectangles []Rectangle, zipY []int) []*TreeNode {
	if len(rectangles) == 0 {
		return nil
	}
	borders := make([]RectBorder, len(rectangles)*2)
	nodes := make([]*TreeNode, len(borders)+1)
	ind := 0
	for _, rect := range rectangles {
		borders[ind] = RectBorder{
			x:      rect.LeftLower.X,
			begY:   search(zipY, rect.LeftLower.Y),
			endY:   search(zipY, rect.RightUpper.Y),
			border: 1,
		}
		ind++
		borders[ind] = RectBorder{
			x:      rect.RightUpper.X,
			begY:   search(zipY, rect.LeftLower.Y),
			endY:   search(zipY, rect.RightUpper.Y),
			border: -1,
		}
		ind++
	}
	sort.Slice(borders, func(i, j int) bool {
		return borders[i].x < borders[j].x
	})
	root := binaryTree(0, len(zipY)-1)
	endX := borders[0].x
	ind = 0
	for _, border := range borders {
		if endX != border.x {
			nodes[ind] = root
			ind++
			endX = border.x
		}
		root = insert(root, border.begY, border.endY-1, border.border)
	}
	nodes[ind] = root
	return nodes
}

func count(node *TreeNode, target int) int {
	if node == nil {
		return 0
	}
	mid := (node.lInd + node.rInd) / 2
	if target <= mid {
		return node.val + count(node.l, target)
	}
	return node.val + count(node.r, target)
}

func insert(node *TreeNode, beg, end, val int) *TreeNode {
	if beg <= node.lInd && node.rInd <= end {
		return &TreeNode{
			val:  node.val + val,
			l:    node.l,
			r:    node.r,
			lInd: node.lInd,
			rInd: node.rInd,
		}
	}
	if node.rInd < beg || end < node.lInd {
		return node
	}
	newNode := &TreeNode{
		val:  node.val,
		lInd: node.lInd,
		rInd: node.rInd,
		l:    node.l,
		r:    node.r,
	}
	newNode.l = insert(newNode.l, beg, end, val)
	newNode.r = insert(newNode.r, beg, end, val)
	return newNode
}

func search(list []int, x int) int {
	l, r := 0, len(list)
	for l <= r-1 {
		middle := (r + l) / 2
		if list[middle] > x {
			r = middle
		} else {
			l = middle + 1
		}
	}
	return l - 1
}
