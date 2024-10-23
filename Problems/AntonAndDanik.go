package main

import "fmt"

func main() {

	var num int
	var str string
	anton := 0
	danik := 0
	fmt.Scan(&num)
	fmt.Scan(&str)
	for i := 0; i < num; i++ {
		if str[i] == 'A' {
			anton++
		} else {
			danik++
		}
	}
	if anton == danik {
		fmt.Println("Friendship")
	} else if anton > danik {
		fmt.Println("Anton")
	} else {
		fmt.Println("Danik")
	}

}
