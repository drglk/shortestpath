package algorithm

import (
	"errors"
	"reflect"
	"testing"
)

func TestShortestPath(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		start    Point
		end      Point
		expected [][]Point
		err      error
	}{
		// case 1: simple path
		{
			name: "simple path",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			start: Point{0, 0},
			end:   Point{2, 2},
			expected: [][]Point{
				{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, // Path 1
				{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}, // Path 2
			},
		},
		//case 2: no path
		{
			name: "no path",
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			start:    Point{0, 0},
			end:      Point{2, 2},
			expected: nil,
			err:      ErrPathNotFound,
		},
		//case 2: too many walls
		{
			name: "too many walls",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			start:    Point{0, 0},
			end:      Point{2, 2},
			expected: nil,
			err:      ErrStartPointIsWall,
		},
		// case 3: start equals end
		{
			name: "start equals end",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			start: Point{0, 0},
			end:   Point{0, 0},
			expected: [][]Point{
				{{0, 0}},
			},
		},

		// case 3: start in in a wall
		{
			name: "start is in a wall",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			start:    Point{1, 1},
			end:      Point{2, 2},
			expected: nil,
			err:      ErrStartPointIsWall,
		},
		// case 4: complex path
		{
			name: "complex path",
			grid: [][]int{
				{1, 2, 0},
				{2, 0, 1},
				{9, 1, 0},
			},
			start: Point{0, 0},
			end:   Point{2, 1},
			expected: [][]Point{
				{{0, 0}, {1, 0}, {2, 0}, {2, 1}},
			},
		},
		// case 4: many paths
		{
			name: "many paths",
			grid: [][]int{
				{9, 9, 9},
				{9, 9, 9},
				{9, 9, 9},
			},
			start: Point{0, 0},
			end:   Point{2, 2},
			expected: [][]Point{
				{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
				{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
				{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}},
				{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {2, 2}},
			},
		},
	}

	for _, tt := range tests {
		result, err := ShortestPath(tt.grid, tt.start, tt.end)

		isValid := false
		for _, expected := range tt.expected {
			if reflect.DeepEqual(result, expected) && errors.Is(err, tt.err) {
				isValid = true
				break
			}
		}

		if !isValid && len(result) == 0 && len(tt.expected) == 0 && errors.Is(err, tt.err) {
			isValid = true
		}

		if !isValid {
			t.Errorf("Test `%s` failed: expected %v, got %v", tt.name, tt.expected, result)
		}
	}
}

func TestIsValid(t *testing.T) {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	visited := [][]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}

	tests := []struct {
		point    Point
		expected bool
	}{
		{Point{0, 0}, true},
		{Point{1, 1}, false}, // Visited
		{Point{1, 0}, true},
		{Point{3, 3}, false}, // Out of bounds
		{Point{1, 2}, true},
		{Point{1, 3}, false}, // Out of bounds
	}

	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			result := needVisit(grid, visited, tt.point)
			if result != tt.expected {
				t.Errorf("Test %d failed: expected %v, got %v", i, tt.expected, result)
			}
		})
	}
}
