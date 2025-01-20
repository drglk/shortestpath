package algorithm

import (
	"container/heap"
	"errors"
)

var (
	ErrPathNotFound     = errors.New("path wasn't found")
	ErrStartPointIsWall = errors.New("start point is a wall")
)

// NeedVisis checks if a point is a wall, is within the grid matrix, and has been visited
func needVisit(grid [][]int, visited [][]bool, point Point) bool {
	return point.X >= 0 && point.X < len(grid) && point.Y >= 0 && point.Y < len(grid[0]) && grid[point.X][point.Y] != 0 && !visited[point.X][point.Y]
}

// ShortestPath finds the shortest path in a maze if such a path exists.
func ShortestPath(grid [][]int, start, end Point) ([]Point, error) {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	if grid[start.X][start.Y] == 0 {
		return nil, ErrStartPointIsWall
	}
	heap.Push(pq, Node{Point: start, Cost: grid[start.X][start.Y], Path: []Point{start}}) // If the starting point equal end point or is in a wall, the path will consist only of start point.

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		cur := node.Point

		if cur == end {
			return node.Path, nil // Founded path
		}

		visited[cur.X][cur.Y] = true

		for _, dir := range directions {
			newPoint := Point{X: cur.X + dir.X, Y: cur.Y + dir.Y}
			if needVisit(grid, visited, newPoint) {
				newPath := append([]Point{}, node.Path...)
				newPath = append(newPath, newPoint)
				heap.Push(pq, Node{Point: newPoint, Cost: node.Cost + grid[newPoint.X][newPoint.Y], Path: newPath})
			}
		}
	}

	return nil, ErrPathNotFound
}
