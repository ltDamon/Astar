package main

import (
	"fmt"
	"strings"
)

type Tile struct {
	x    int
	y    int
	view string
}

// 保存地图的基本信息
type Map struct {
	Tiles [][]*Tile
	maxX  int
	maxY  int
}

func (m *Map) GetTile(x, y int) *Tile {
	if x >= m.maxX || y >= m.maxY || x < 0 || y < 0 {
		return nil
	}
	return m.Tiles[x][y]
}

//是否可穿越
func (m *Map) IsTraversable(x, y int) bool {
	if x >= m.maxX || y >= m.maxY || x < 0 || y < 0 {
		return false
	}
	tile := m.Tiles[x][y]
	if tile.view == "X" {
		return false
	}
	return true
}

func (m *Map) GetNeighborTraversableTiles(curPoint *Tile) (neighbors []*Tile) {
	if x, y := curPoint.x, curPoint.y-1; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}
	if x, y := curPoint.x, curPoint.y+1; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}
	if x, y := curPoint.x+1, curPoint.y; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}
	if x, y := curPoint.x-1, curPoint.y; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}

	if x, y := curPoint.x+1, curPoint.y-1; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}

	if x, y := curPoint.x+1, curPoint.y+1; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}

	if x, y := curPoint.x-1, curPoint.y+1; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}
	if x, y := curPoint.x-1, curPoint.y-1; x >= 0 && x < m.maxX && y >= 0 && y < m.maxY {
		if m.IsTraversable(x, y) {
			neighbors = append(neighbors, m.Tiles[x][y])
		}
	}
	return neighbors
}

func (m *Map) PrintMap(r *PathFindManager, start, end *Tile) {
	fmt.Println("map's border:", m.maxX, m.maxY)
	for y := m.maxY - 1; y >= 0; y-- {
		for x := 0; x <= m.maxX-1; x++ {
			if r != nil {
				if x == start.x && y == start.y {
					fmt.Print("S")
					goto NEXT
				}
				if x == end.x && y == end.y {
					fmt.Print("E")
					goto NEXT
				}
				for i := 0; i < len(r.ThePath); i++ {
					if r.ThePath[i].Tile.x == x && r.ThePath[i].Tile.y == y {
						fmt.Print("R")
						goto NEXT
					}
				}
			}
			fmt.Print(m.Tiles[x][y].view)
		NEXT:
		}
		fmt.Println()
	}
}

func NewMap(charMap []string) {
	m = new(Map)
	maxX := len(strings.Split(charMap[0], " "))
	maxY := len(charMap)
	m.Tiles = make([][]*Tile, maxX)
	for x := 0; x < maxX; x++ {
		m.Tiles[x] = make([]*Tile, maxY)
		for y := 0; y < maxY; y++ {
			cols := strings.Split(charMap[maxY-1-y], " ")
			m.Tiles[x][y] = &Tile{x, y, cols[x]}
		}
	}

	m.maxX = maxX
	m.maxY = maxY

	return
}

func NewPathFindManager() {
	_P = &PathFindManager{
		OpenList: make([]*PathNode, 0),
		Nodes:    make([][]*PathNode, m.maxX),
		ThePath:  make([]*PathNode, 0),
	}

	for i := 0; i < m.maxX; i++ {
		_P.Nodes[i] = make([]*PathNode, m.maxY)
		for j := 0; j < m.maxY; j++ {
			_P.Nodes[i][j] = &PathNode{}
		}
	}
}
