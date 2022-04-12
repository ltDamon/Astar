package main

import (
	"container/heap"
	"fmt"
	"math"
)

const (
	NO_LIST     = 0
	OPEN_LIST   = 1
	CLOSED_LIST = 2
	EPSILON     = 0.00001
)

type OpenList []*PathNode

func (l OpenList) Len() int { return len(l) }
func (l OpenList) Less(i, j int) bool {
	return l[i].gVal+l[i].hVal < l[j].gVal+l[j].hVal
}
func (l OpenList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l *OpenList) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*l = append(*l, x.(*PathNode))
}

func (l *OpenList) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

//========================================================================================
type PathNode struct {
	Tile      *Tile
	parent    *PathNode
	gVal      float64
	hVal      float64
	List      int
	SearchCnt int64
}

type PathFindManager struct {
	OpenList   OpenList
	Nodes      [][]*PathNode
	ThePath    []*PathNode
	_SearchCnt int64 //当前寻路次数
}

func (s *PathFindManager) GetNodes(tile *Tile) *PathNode {
	return s.Nodes[tile.x][tile.y]
}

func (s *PathFindManager) ClearOpenList() {
	s.OpenList = make([]*PathNode, 0)
	heap.Init(&s.OpenList)
}

func (s *PathFindManager) CleanRoad() {
	s.ThePath = make([]*PathNode, 0)
}

func (s *PathFindManager) AddToOpen(cur *Tile) {
	node := s.GetNodes(cur)
	node.List = OPEN_LIST
	heap.Push(&s.OpenList, node)
}

//========================================================================================
//欧几里得距离
func EuclideanDistance(s, s1 *Tile) float64 {
	deltaX := math.Abs(float64(s1.x - s.x))
	deltaY := math.Abs(float64(s1.y - s.y))
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func NewPathNode(s *Tile, end *Tile) *PathNode {
	node := _P.GetNodes(s)
	if node.SearchCnt < _P._SearchCnt {
		node.Tile = s
		node.SearchCnt = _P._SearchCnt
		node.hVal = EuclideanDistance(s, end)
		node.gVal = math.MaxFloat64
		node.List = NO_LIST
	}
	return node
}

//========================================================================================

func AStarSearch(start, end *Tile) {
	_P._SearchCnt++
	_P.ClearOpenList()
	_P.CleanRoad()
	startNode := NewPathNode(start, end)
	endNode := NewPathNode(end, end)
	startNode.gVal = 0
	_P.AddToOpen(start)
	for len(_P.OpenList) != 0 {
		a := heap.Pop(&_P.OpenList)
		curr := a.(*PathNode)
		curr.List = CLOSED_LIST
		if endNode.gVal < curr.gVal+curr.hVal+EPSILON {
			fmt.Println("hahahaha")
			break
		}
		neighbors := m.GetNeighborTraversableTiles(curr.Tile)
		for _, tile := range neighbors {
			succ := NewPathNode(tile, end)
			if succ.List != CLOSED_LIST {
				//fmt.Println(*curr.Tile, "neibor", *tile)
				newGValue := curr.gVal + EuclideanDistance(curr.Tile, succ.Tile)
				//fmt.Println("curr:", *curr.Tile, "succ:", *succ.Tile, "new:", newGValue, "old:", succ.gVal)
				if newGValue+EPSILON < succ.gVal {
					succ.gVal = newGValue
					//fmt.Println("curr:", *curr.Tile, "succ:", *succ.Tile, "fvalue:", succ.gVal+succ.hVal)
					succ.parent = curr
					_P.AddToOpen(succ.Tile)
				}
			}
		}
	}
	if endNode.gVal < math.MaxFloat64 {
		curNode := endNode
		for curNode != startNode {
			_P.ThePath = append(_P.ThePath, curNode)
			curNode = curNode.parent
		}
	}
}

func ThetaStarSearch(start, end *Tile) {
	_P._SearchCnt++
	_P.ClearOpenList()
	_P.CleanRoad()
	startNode := NewPathNode(start, end)
	endNode := NewPathNode(end, end)
	startNode.gVal = 0

	//Set 'start's parent as itself. When 'start' is expanded for the first time, the grandparent of 'start's successor will be 'start' as well.
	startNode.parent = startNode
	_P.AddToOpen(start)
	for len(_P.OpenList) != 0 {
		a := heap.Pop(&_P.OpenList)
		curr := a.(*PathNode)
		curr.List = CLOSED_LIST
		if endNode.gVal < curr.gVal+curr.hVal+EPSILON {
			fmt.Println("hahahaha")
			break
		}
		neighbors := m.GetNeighborTraversableTiles(curr.Tile)
		for _, tile := range neighbors {
			succ := NewPathNode(tile, end)
			if succ.List != CLOSED_LIST {
				var newGValue float64
				var newParent *PathNode
				if LineOfSight(curr.parent.Tile, succ.Tile) {
					newParent = curr.parent
				} else {
					newParent = curr
				}
				newGValue = newParent.gVal + EuclideanDistance(newParent.Tile, succ.Tile)

				if newGValue+EPSILON < succ.gVal {
					succ.gVal = newGValue
					succ.parent = newParent
					_P.AddToOpen(succ.Tile)
				}
			}
		}
	}

	if endNode.gVal < math.MaxFloat64 {
		curNode := endNode
		for curNode != startNode {
			_P.ThePath = append(_P.ThePath, curNode)
			curNode = curNode.parent
		}
	}
}

func LazyThetaStarSearch(start, end *Tile) {
	_P._SearchCnt++
	_P.ClearOpenList()
	_P.CleanRoad()
	startNode := NewPathNode(start, end)
	endNode := NewPathNode(end, end)
	startNode.gVal = 0

	//Set 'start's parent as itself. When 'start' is expanded for the first time, the grandparent of 'start's successor will be 'start' as well.
	startNode.parent = startNode
	_P.AddToOpen(start)
	for len(_P.OpenList) != 0 {
		a := heap.Pop(&_P.OpenList)
		curr := a.(*PathNode)
		curr.List = CLOSED_LIST
		if endNode.gVal < curr.gVal+curr.hVal+EPSILON {
			fmt.Println("hahahaha")
			break
		}
		ValidateParent(curr, endNode)
		newParent := curr.parent
		neighbors := m.GetNeighborTraversableTiles(curr.Tile)
		for _, tile := range neighbors {
			succ := NewPathNode(tile, end)
			if succ.List != CLOSED_LIST {
				newGValue := newParent.gVal + EuclideanDistance(newParent.Tile, succ.Tile)
				if newGValue+EPSILON < succ.gVal {
					succ.gVal = newGValue
					succ.parent = newParent
					_P.AddToOpen(succ.Tile)
				}
			}
		}
	}

	if endNode.gVal < math.MaxFloat64 {
		ValidateParent(endNode, endNode)
		curNode := endNode
		for curNode != startNode {
			_P.ThePath = append(_P.ThePath, curNode)
			curNode = curNode.parent
		}
	}

}

func ValidateParent(s, end *PathNode) {
	if !LineOfSight(s.parent.Tile, s.Tile) {
		s.gVal = math.MaxFloat64
		neighbors := m.GetNeighborTraversableTiles(s.Tile)
		for _, tile := range neighbors {
			newParent := NewPathNode(tile, end.Tile)
			if newParent.List == CLOSED_LIST {
				newGValue := newParent.gVal + EuclideanDistance(newParent.Tile, s.Tile)
				if newGValue < s.gVal {
					s.gVal = newGValue
					s.parent = newParent
				}
			}
		}
	}
}

func LineOfSight(l1, l2 *Tile) bool {
	x1 := l1.x
	y1 := l1.y

	x2 := l2.x
	y2 := l2.y

	dy := y2 - y1
	dx := x2 - x1

	var f, sy, sx, x_offset, y_offset int
	if dy < 0 {
		dy = -dy
		sy = -1
		y_offset = 0
	} else {
		sy = 1
		y_offset = 1
	}

	if dx < 0 {
		dx = -dx
		sx = -1
		x_offset = 0
	} else {
		sx = 1
		x_offset = 1
	}

	if dx >= dy { // Move along the x axis and increment/decrement y when f >= dx.
		for x1 != x2 {
			f = f + dy
			if f >= dx { // We are changing rows, we might need to check two cells this iteration.
				if !m.IsTraversable(x1+x_offset, y1+y_offset) {
					return false
				}

				y1 = y1 + sy
				f = f - dx
			}

			if f != 0 { // If f == 0, then we are crossing the row at a corner point and we don't need to check both cells.
				if !m.IsTraversable(x1+x_offset, y1+y_offset) {
					return false
				}
			}

			if dy == 0 { // If we are moving along a horizontal line, either the north or the south cell should be unblocked.

				if !m.IsTraversable(x1+x_offset, y1) && !m.IsTraversable(x1+x_offset, y1+1) {
					return false
				}
			}

			x1 += sx
		}
	} else { //if (dx < dy). Move along the y axis and increment/decrement x when f >= dy.
		for y1 != y2 {
			f = f + dx
			if f >= dy {
				if !m.IsTraversable(x1+x_offset, y1+y_offset) {
					return false
				}

				x1 = x1 + sx
				f = f - dy
			}

			if f != 0 {
				if !m.IsTraversable(x1+x_offset, y1+y_offset) {
					return false
				}
			}

			if dx == 0 {
				if !m.IsTraversable(x1, y1+y_offset) && !m.IsTraversable(x1+1, y1+y_offset) {
					return false
				}
			}

			y1 += sy
		}
	}
	return true
}
