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

func CreateNeighborhood(x, y int) [][]int {

	var neighborhood = make([][]int, x)

	for i := range neighborhood {
		neighborhood[i] = make([]int, y)
	}

	return neighborhood
}

func SetCell(neighborhood [][]int, x, y int) {
	neighborhood[x][y] = 1
}

func GetCell(neighborhood [][]int, x, y int) int {
	return neighborhood[x][y]
}

func print(s string, x [][]int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Printf("%d ", x[i][j])
		}
		fmt.Println("")
	}
}
