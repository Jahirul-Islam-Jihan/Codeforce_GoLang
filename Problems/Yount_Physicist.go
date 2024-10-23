package main

import "fmt"

func main() {
	var n, sum1, sum2, sum3 int
	fmt.Scan(&n)
	size := n * 3
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	for j := 0; j < size; j = j + 3 {
		sum1 = sum1+ arr[j]
		sum2 = sum2+ arr[j+1]
		sum3 = sum3+ arr[j+2]
	}
	if sum1 == 0 && sum2 == 0 && sum3 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
