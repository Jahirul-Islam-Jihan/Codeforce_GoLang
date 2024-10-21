package main

import (
	"fmt"
)

func main() {
	var matrix [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == 1 {
				x := abs(i-2)
				y := abs(j - 2)
				fmt.Println(abs(x+y))
				break
			}
		}
	}
}
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}


