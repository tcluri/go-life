package gameoflife

import (
	"fmt"
	"math/rand"
	"strings"
)

func CreateEmptyGrid(N int) Grid {
	grid := make([][]bool, N)
	for row := range grid {
		grid[row] = make([]bool, N)
	}
	return grid
}

func PopulateCells(grid Grid) Grid {
	for row := range grid {
		for col := range grid[row] {
			grid[row][col] = rand.Intn(2) == 1
		}
	}
	return grid
}

func FindIndexInGrid(index int, gridLen int) int {
	if index < 0 {
		return gridLen - 1
	} else if index >= gridLen {
		return 0
	} else {
		return index
	}
}

func AliveSurroundingCells(cellGrid Grid, row int, col int) int {
	gridLen := len(cellGrid)
	// All directions surrounding a single cell - grid is overlapping in all directions
	northEast := cellGrid[FindIndexInGrid(row-1, gridLen)][FindIndexInGrid(col+1, gridLen)]
	north := cellGrid[FindIndexInGrid(row-1, gridLen)][col]
	northWest := cellGrid[FindIndexInGrid(row-1, gridLen)][FindIndexInGrid(col-1, gridLen)]
	east := cellGrid[row][FindIndexInGrid(col+1, gridLen)]
	west := cellGrid[row][FindIndexInGrid(col-1, gridLen)]
	southEast := cellGrid[FindIndexInGrid(row+1, gridLen)][FindIndexInGrid(col+1, gridLen)]
	south := cellGrid[FindIndexInGrid(row+1, gridLen)][col]
	southWest := cellGrid[FindIndexInGrid(row+1, gridLen)][FindIndexInGrid(col-1, gridLen)]

	surroundingEightCells := []bool{northEast, north, northWest, east, west, southEast, south, southWest}

	// Cells alive
	numAliveCells := 0
	for _, cell := range surroundingEightCells {
		if cell {
			numAliveCells += 1
		}
	}
	return numAliveCells
}

func PrintCellsInGrid(grid Grid) {
	var sb strings.Builder
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				sb.WriteRune('\u2588')
			} else {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func CurrentAlive(grid Grid) int {
	alive := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] {
				alive += 1
			}
		}
	}
	return alive
}
