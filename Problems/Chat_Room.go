package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	length := len(str)
	str2 := "hello"
	count := 0
	checker:=0

	for i := 0; i < 5; i++ {
		for j := checker; j < length; j++ {
			if str2[i] == str[j] {
				checker=j+1
				count++
				break
			}
		}
	}
	if count == 5 {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}

}
