package main

import (
	"errors"
	"reflect"
	"shortestpath/algorithm"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedGrid  [][]int
		expectedStart algorithm.Point
		expectedEnd   algorithm.Point
		expectedErr   error
	}{
		{
			name: "valid input",
			input: `3 3
			1 2 0
			2 0 1
			9 1 0
			0 0 2 1`,
			expectedGrid: [][]int{
				{1, 2, 0},
				{2, 0, 1},
				{9, 1, 0},
			},
			expectedStart: algorithm.Point{X: 0, Y: 0},
			expectedEnd:   algorithm.Point{X: 2, Y: 1},
			expectedErr:   nil,
		},
		{
			name: "invalid maze dimensions",
			input: `3
			1 2 0
			2 0 1
			9 1 0
			0 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("maze dimensions must retain exactly two numbers"),
		},
		{
			name: "negative grid dimensions",
			input: `-3 3
			1 2 0
			2 0 1
			9 1 0
			0 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the dimensions of the maze must be positive integers"),
		},
		{
			name: "invalid grid dimensions(input < lines)",
			input: `2 3
			1 2 0
			2 0 1
			9 1 0
			0 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the start and end points must contain exactly four numbers"),
		},
		{
			name: "invalid grid dimensions(input < cols)",
			input: `2 3
			1 2
			2 0
			9 1
			0 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("line 1 must have exactly 3 numbers"),
		},
		{
			name: "invalid grid dimensions(input > cols)",
			input: `2 3
			1 2 3 4
			2 0 3 4 
			9 1 3 4
			0 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("line 1 must have exactly 3 numbers"),
		},
		{
			name: "not enough grid lines",
			input: `3 3
			1 2 0
			2 0 1
			0 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("line 3 must have exactly 3 numbers"),
		},
		{
			name: "start point outside the matrix",
			input: `3 3
			1 2 0
			2 0 1
			9 1 0
			-1 0 2 1`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the start point is outside the matrix"),
		},
		{
			name: "end point outside the matrix",
			input: `3 3
  			1 2 0
  			2 0 1
			9 1 0
 			0 0 2 3`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the end point is outside the matrix"),
		},
		{
			name: "invalid data input",
			input: `a a
			a a a
			a a a a`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the dimensions of the maze must be positive integers"),
		},
		{
			name:          "invalid data input",
			input:         `a a`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the dimensions of the maze must be positive integers"),
		},
		{
			name:          "empty input",
			input:         ``,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("not enough data for maze dimensions"),
		},
		{
			name: "invalid matrix element",
			input: `3 3
			1 1 1
			0 a 1
			1 1 1
			0 0 2 0`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("all elements of the maze must be numbers from 0 to 9 (error in line 2, column 2)"),
		},
		{
			name: "invalid(not integer) start coordinates",
			input: `3 3
			1 1 1
			0 0 1
			1 1 1
			0 a 2 0`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("coordinates must be integers"),
		},
		{
			name: "invalid(not integer) end coordinates",
			input: `3 3
			1 1 1
			0 0 1
			1 1 1
			0 0 2 a`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("coordinates must be integers"),
		},
		{
			name: "invalid amount of coordinates",
			input: `3 3
			1 1 1
			0 0 1
			1 1 1
			0 0 2`,
			expectedGrid:  nil,
			expectedStart: algorithm.Point{},
			expectedEnd:   algorithm.Point{},
			expectedErr:   errors.New("the start and end points must contain exactly four numbers"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.input)
			grid, startX, startY, endX, endY, err := readInput(reader)

			if test.expectedErr != nil {
				if err == nil || err.Error() != test.expectedErr.Error() {
					t.Fatalf("Expected error: %v, got: %v", test.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(grid, test.expectedGrid) {
				t.Errorf("Grid mismatch. Expected %v, got %v", test.expectedGrid, grid)
			}

			if startX != test.expectedStart.X || startY != test.expectedStart.Y {
				t.Errorf("Start point mismatch. Expected %v, got (%d, %d)", test.expectedStart, startX, startY)
			}
			if endX != test.expectedEnd.X || endY != test.expectedEnd.Y {
				t.Errorf("End point mismatch. Expected %v, got (%d, %d)", test.expectedEnd, endX, endY)
			}
		})
	}
}
