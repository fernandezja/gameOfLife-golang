//from https://tour.golang.org/welcome/1
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Game of life!")

	var g = CreateNeighborhood(3, 3)

	print(" Step 1 ", g)

}

//CreateNeighborhood is for create a new neighborhood with lifeless (empty) cells
func CreateNeighborhood(x, y int) [][]int {

	var neighborhood = make([][]int, x)

	for i := range neighborhood {
		neighborhood[i] = make([]int, y)
	}

	return neighborhood
}

//SetCell is for set life in a particular cell (useful for unit tests)
func SetCell(neighborhood [][]int, x, y int) {
	neighborhood[x][y] = 1
}

//GetCell is for Get the value of a particular cell
func GetCell(neighborhood [][]int, x, y int) int {
	return neighborhood[x][y]
}

//NeighborsCount is a neighbors counter (Simple... )
func NeighborsCount(neighborhood [][]int, x, y int) int {

	var neighbors = 0

	//fmt.Printf("CELL %d %d  \n", x, y)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			var row = x + i
			var column = y + j

			//fmt.Printf("Iteration %d %d > row %d column %d \n", i, j, row, column)

			//fmt.Printf("len(neighborhood) = %d \n", len(neighborhood))

			if row >= 0 && row < len(neighborhood) {

				if column >= 0 && column < len(neighborhood[row]) {

					//fmt.Printf("len(neighborhood[row] = %d \n", len(neighborhood))
					//fmt.Printf("row %d column %d > value > %d \n", row, column, neighborhood[row][column])

					if x != row || y != column {
						neighbors = neighbors + neighborhood[row][column]
					}
				}
			}

		}
	}

	//var neighbors0 = neighborhood[x-1][y-1] + neighborhood[x-1][y] + neighborhood[x-1][y+1]
	//var neighbors1 = neighborhood[x][y-1] + neighborhood[x][y+1]
	//var neighbors2 = neighborhood[x+1][y-1] + neighborhood[x+1][y] + neighborhood[x+1][y+1]

	//fmt.Printf("neighbors count: %d", neighbors)
	return neighbors
}

//GrowUp is for... GrowUp all neighborhood
//func GrowUp(neighborhood [][]int) [][]int {
//
//	var new = CreateNeighborhood(len(neighborhood), len(neighborhood[0]))
//
//	for i := 0; i < len(neighborhood); i++ {
//		for j := 0; j < len(neighborhood[i]); j++ {
//
//			new[i][j] = GrowUpCell(neighborhood[i][j])
//
//			//fmt.Printf("%d ", neighborhood[i][j])
//		}
//	}
//
//	return new
//}

//GrowUpCell is for makes a cell grow (depending on its neighbors)
func GrowUpCell(neighborhood [][]int, x, y int) int {

	var actualValue = neighborhood[x][y]
	var neighborsCount = NeighborsCount(neighborhood, x, y)

	if actualValue == 1 {
		switch {
		case neighborsCount < 2:
			return 0
		case neighborsCount > 3:
			return 0
		case neighborsCount == 3:
			return 1
		}
	} else if neighborsCount == 3 {
		return 1
	}

	return 0
}

func print(s string, x [][]int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Printf("%d ", x[i][j])
		}
		fmt.Println("")
	}
}
