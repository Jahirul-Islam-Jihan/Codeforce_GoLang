package main

import "fmt"


func main() {
    var n, k, count int
    count = 0
    fmt.Scan(&n, &k)
    arrayOfInteger := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arrayOfInteger[i])
    }
	kthPlace := arrayOfInteger[k-1]
	for i := 0; i < n; i++ {
		if arrayOfInteger[i] >= kthPlace && arrayOfInteger[i]>0{
            count++ 
        }
    }
    fmt.Println(count)
}
