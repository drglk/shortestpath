package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"shortestpath/algorithm"
	"strconv"
	"strings"
)

func main() {
	grid, startX, startY, endX, endY, err := readInput(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	start := algorithm.Point{X: startX, Y: startY}
	end := algorithm.Point{X: endX, Y: endY}

	result, err := algorithm.ShortestPath(grid, start, end)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Path wasn't found")
		os.Exit(1)

	} else {
		for _, point := range result {
			fmt.Printf("%d %d\n", point.X, point.Y)
		}
		fmt.Println(".")
	}
}

// ReadInput reads the input data and returns the grid matrix, start and end coordinates if the input data is valid.
// Otherwise, it returns an error.
func readInput(r io.Reader) ([][]int, int, int, int, int, error) {
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		return nil, 0, 0, 0, 0, fmt.Errorf("not enough data for maze dimensions")
	}
	dimensions := strings.Fields(scanner.Text())
	if len(dimensions) != 2 {
		return nil, 0, 0, 0, 0, fmt.Errorf("maze dimensions must retain exactly two numbers")
	}

	rows, err1 := strconv.Atoi(dimensions[0])
	cols, err2 := strconv.Atoi(dimensions[1])
	if err1 != nil || err2 != nil || rows <= 0 || cols <= 0 {
		return nil, 0, 0, 0, 0, fmt.Errorf("the dimensions of the maze must be positive integers")
	}

	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			return nil, 0, 0, 0, 0, fmt.Errorf("not enough lines to define the structure of the maze")
		}
		rowValues := strings.Fields(scanner.Text())
		if len(rowValues) != cols {
			return nil, 0, 0, 0, 0, fmt.Errorf("line %d must have exactly %d numbers", i+1, cols)
		}

		grid[i] = make([]int, cols)
		for j, value := range rowValues {
			num, err := strconv.Atoi(value)
			if err != nil || num < 0 || num > 9 {
				return nil, 0, 0, 0, 0, fmt.Errorf("all elements of the maze must be numbers from 0 to 9 (error in line %d, column %d)", i+1, j+1)
			}
			grid[i][j] = num
		}
	}

	if !scanner.Scan() {
		return nil, 0, 0, 0, 0, fmt.Errorf("not enough data to determine start and end points")
	}
	coordinates := strings.Fields(scanner.Text())
	if len(coordinates) != 4 {
		return nil, 0, 0, 0, 0, fmt.Errorf("the start and end points must contain exactly four numbers")
	}

	startX, err1 := strconv.Atoi(coordinates[0])
	startY, err2 := strconv.Atoi(coordinates[1])
	endX, err3 := strconv.Atoi(coordinates[2])
	endY, err4 := strconv.Atoi(coordinates[3])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("coordinates must be integers")
	}

	if startX < 0 || startX >= rows || startY < 0 || startY >= cols {
		return nil, 0, 0, 0, 0, fmt.Errorf("the start point is outside the matrix")
	}
	if endX < 0 || endX >= rows || endY < 0 || endY >= cols {
		return nil, 0, 0, 0, 0, fmt.Errorf("the end point is outside the matrix")
	}

	if grid[startX][startY] == 0 {
		return nil, 0, 0, 0, 0, fmt.Errorf("the start point is on the wall")
	}
	if grid[endX][endY] == 0 {
		return nil, 0, 0, 0, 0, fmt.Errorf("the end point is on the wall")
	}

	return grid, startX, startY, endX, endY, nil
}
