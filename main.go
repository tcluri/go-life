package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Grid [][]bool

func main() {
	// Use flag to set the length of grid
	gridLen := flag.Int("grid_size", 15, "Enter the length of the grid(square) of board for the game of life")
	flag.Parse()

	// Create an empty grid for the cells
	initialGrid := createEmptyGrid(*gridLen)
	// Populate the grid cells
	currentGeneration := populateCells(initialGrid)
	// Evolution!
	generationNum := 1
	for {
		fmt.Printf("Generation #%d\n", generationNum)
		aliveCells := currentAlive(currentGeneration)
		if aliveCells == 0 {
			fmt.Println("All cells have died")
			return
		} else {
			fmt.Printf("Alive: %d\n", aliveCells)
		}
		printCellsInGrid(currentGeneration)

		// Evolve the cell grid
		currentGeneration = evolution(currentGeneration)
		time.Sleep(800 * time.Millisecond)
		fmt.Print("\033[H\033[2J") // Clear screen
		generationNum++
	}
}

func evolution(currentCellsInGrid Grid) Grid {
	cellsGridLen := len(currentCellsInGrid)
	evolvedCellsInGrid := createEmptyGrid(cellsGridLen)
	// Calculate according to the evolution rules
	for row := range currentCellsInGrid {
		for col := range currentCellsInGrid[row] {
			currentCell := currentCellsInGrid[row][col]
			// Create surrounding cells for a particular cell
			numSurroundingCellsAlive := aliveSurroundingCells(currentCellsInGrid, row, col)
			// Cell birth/death logic
			if !currentCell && numSurroundingCellsAlive == 3 {
				evolvedCellsInGrid[row][col] = true
			} else if currentCell && (numSurroundingCellsAlive == 2 || numSurroundingCellsAlive == 3) {
				evolvedCellsInGrid[row][col] = true
			} else {
				evolvedCellsInGrid[row][col] = false
			}
		}
	}
	return evolvedCellsInGrid
}

func aliveSurroundingCells(cellGrid Grid, row int, col int) int {
	gridLen := len(cellGrid)
	// All directions surrounding a single cell - grid is overlapping in all directions
	northEast := cellGrid[findIndexInGrid(row-1, gridLen)][findIndexInGrid(col+1, gridLen)]
	north := cellGrid[findIndexInGrid(row-1, gridLen)][col]
	northWest := cellGrid[findIndexInGrid(row-1, gridLen)][findIndexInGrid(col-1, gridLen)]
	east := cellGrid[row][findIndexInGrid(col+1, gridLen)]
	west := cellGrid[row][findIndexInGrid(col-1, gridLen)]
	southEast := cellGrid[findIndexInGrid(row+1, gridLen)][findIndexInGrid(col+1, gridLen)]
	south := cellGrid[findIndexInGrid(row+1, gridLen)][col]
	southWest := cellGrid[findIndexInGrid(row+1, gridLen)][findIndexInGrid(col-1, gridLen)]

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

func createEmptyGrid(N int) Grid {
	grid := make([][]bool, N)
	for row := range grid {
		grid[row] = make([]bool, N)
	}
	return grid
}

func populateCells(grid Grid) Grid {
	for row := range grid {
		for col := range grid[row] {
			grid[row][col] = rand.Intn(2) == 1
		}
	}
	return grid
}

func findIndexInGrid(index int, gridLen int) int {
	if index < 0 {
		return gridLen - 1
	} else if index >= gridLen {
		return 0
	} else {
		return index
	}
}

func printCellsInGrid(grid Grid) {
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

func currentAlive(grid Grid) int {
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
