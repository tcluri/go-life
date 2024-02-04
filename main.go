package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tcluri/go-life/gameoflife"
)

func main() {
	// Use flag to set the length of grid
	gridLen := flag.Int("grid_size", 15, "Enter the length of the grid(square) of board for the game of life")
	flag.Parse()

	// Create an empty grid for the cells
	initialGrid := gameoflife.CreateEmptyGrid(*gridLen)
	// Populate the grid cells
	currentGeneration := gameoflife.PopulateCells(initialGrid)
	// Evolution!
	generationNum := 1
	for {
		fmt.Printf("Generation #%d\n", generationNum)
		aliveCells := gameoflife.CurrentAlive(currentGeneration)
		if aliveCells == 0 {
			fmt.Println("All cells have died")
			return
		} else {
			fmt.Printf("Alive: %d\n", aliveCells)
		}
		gameoflife.PrintCellsInGrid(currentGeneration)

		// Evolve the cell grid
		currentGeneration = gameoflife.Evolution(currentGeneration)
		time.Sleep(800 * time.Millisecond)
		fmt.Print("\033[H\033[2J") // Clear screen
		generationNum++
	}
}
