package main

import "fmt"

var (
	m  *Map
	_P *PathFindManager
)

func main() {
	presetMap := []string{
		". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . X . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . X . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . X X X X X X X X X X X X X X X X X X X X X . . . . . . . . . . . . . . . . . . .",
		". . . . . . X . X . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . X X . . . . . . . . . . . . x . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . X . . . . . . . . . . . . X . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . X X X X X X X X X X X X X X X X . X . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . X . . . . . X . . X . . . X . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . X . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . X . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .",
	}
	NewMap(presetMap)
	NewPathFindManager()

	start := m.GetTile(0, 9)
	end := m.GetTile(12, 7)

	AStarSearch(start, end)
	if len(_P.ThePath) > 0 {
		fmt.Println("find the path")
		m.PrintMap(_P, start, end)
	}

	ThetaStarSearch(start, end)
	if len(_P.ThePath) > 0 {
		fmt.Println("theta find the path")
		m.PrintMap(_P, start, end)
	}

	LazyThetaStarSearch(start, end)
	if len(_P.ThePath) > 0 {
		fmt.Println("lazy theta find the path")
		m.PrintMap(_P, start, end)
	}
}
