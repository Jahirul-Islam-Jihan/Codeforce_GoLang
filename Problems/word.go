package main

import "fmt"

func main() {
	var str string
	var countLower, countUpper int
	fmt.Scan(&str)
	length := len(str)
	for i := 0; i < length; i++ {
		if 'z' >= str[i] && str[i] >= 'a' {
			countLower++
		} else {
			countUpper++
		}
	}
	if countLower >= countUpper {
		for i := 0; i < length; i++ {
			if 'z' >= str[i] && str[i] >= 'a' {
				fmt.Print(string(str[i]))
			} else {
				fmt.Print(string(str[i] + 32))
			}

		}

	} else {
		for i := 0; i < length; i++ {
			if 'Z' >= str[i] && str[i] >= 'A' {
				fmt.Print(string(str[i]))
			} else {
				fmt.Print(string(str[i] - 32))
			}
		}
	}
}
