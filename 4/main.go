package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	puzzle := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
	var array []float64
	for _, row := range puzzle {
		for _, character := range row {
			if character == '@' {
				array = append(array, 1.0)
			} else {
				array = append(array, 0.0)
			}
		}
	}
	A := mat.NewDense(10, 10, array)
	fmt.Println(mat.Formatted(A))

	output := mat.DenseCopyOf(A)

	// loop through elements, add surroundings for each roll of paper
	for row := range 10 {
		for col := range 10 {
			element := A.At(row, col)
			if element == 1.0 {
				up := getUp(A, row, col)
				down := getDown(A, row, col)
				left := getLeft(A, row, col)
				right := getRight(A, row, col)
				upRight := getUpRight(A, row, col)
				upLeft := getUpLeft(A, row, col)
				downRight := getDownRight(A, row, col)
				downLeft := getDownLeft(A, row, col)

				neighbors := []float64{up, down, left, right, upRight, upLeft, downRight, downLeft}
				surrouningSum := sumNeighbors(neighbors)
				fmt.Printf("row %d ", row)
				fmt.Printf("col %d \n", col)
				fmt.Printf("sum of neighbors %f \n", surrouningSum)
				if surrouningSum < 4 {
					// 9 as a placeholder to mark Xs in the solution
					output.Set(row, col, 9)
				}
			}
		}
	}
	solution := buildSolution(output)
	for _, array := range solution {
		fmt.Println(array)
	}
}

func buildSolution(A *mat.Dense) []string {
	var solution []string
	for row := range 10 {
		var rowString string
		for col := range 10 {
			element := A.At(row, col)
			switch element {
			case 1.0:
				rowString += "@"
			case 9.0:
				rowString += "X"
			default:
				rowString += "."
			}
		}
		solution = append(solution, rowString)
	}
	return solution
}

func sumNeighbors(surroundings []float64) float64 {
	surroundingSum := 0.0
	for _, element := range surroundings {
		surroundingSum += element
	}
	return surroundingSum

}

func getUp(A *mat.Dense, row int, col int) float64 {
	// same col, one row up
	if row == 0 {
		return 0
	}
	return A.At(row-1, col)
}

func getDown(A *mat.Dense, row int, col int) float64 {
	// same col, one row down
	if row == 9 {
		return 0
	}
	return A.At(row+1, col)
}

func getRight(A *mat.Dense, row int, col int) float64 {
	// same row, one col up
	if col == 9 {
		return 0
	}
	return A.At(row, col+1)
}

func getLeft(A *mat.Dense, row int, col int) float64 {
	if col == 0 {
		return 0
	}
	return A.At(row, col-1)
}

func getUpRight(A *mat.Dense, row int, col int) float64 {
	// one row up, one col up
	if row == 0 || col == 9 {
		return 0
	}
	return A.At(row-1, col+1)
}

func getUpLeft(A *mat.Dense, row int, col int) float64 {
	// one row up, one col down
	if row == 0 || col == 0 {
		return 0
	}
	return A.At(row-1, col-1)
}

func getDownRight(A *mat.Dense, row int, col int) float64 {
	// one row down, one col up
	if row == 9 || col == 9 {
		return 0
	}
	return A.At(row+1, col+1)
}

func getDownLeft(A *mat.Dense, row int, col int) float64 {
	// one row down, one col down
	if row == 9 || col == 0 {
		return 0
	}
	return A.At(row+1, col-1)
}
