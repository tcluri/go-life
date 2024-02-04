package gameoflife

func Evolution(currentCellsInGrid Grid) Grid {
	cellsGridLen := len(currentCellsInGrid)
	evolvedCellsInGrid := CreateEmptyGrid(cellsGridLen)
	// Calculate according to the evolution rules
	for row := range currentCellsInGrid {
		for col := range currentCellsInGrid[row] {
			currentCell := currentCellsInGrid[row][col]
			// Create surrounding cells for a particular cell
			numSurroundingCellsAlive := AliveSurroundingCells(currentCellsInGrid, row, col)
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
