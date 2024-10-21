package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	n := x % 5
	p := x - n
	if n == 0 {
		count := (p / 5)
		fmt.Println(count)
	} else {
		count := (p / 5) + 1
		fmt.Println(count)
	}

}
