package main

import "fmt"
// import "sort"
import "strings"

func main(){
var str1, str2 string
fmt.Scan(&str1, &str2)

	lowerString1 := strings.ToLower(str1)
	lowerString2 := strings.ToLower(str2)
	if lowerString1 < lowerString2 {
		fmt.Println(-1)
	}else if lowerString1 > lowerString2{
		fmt.Println(1)
	}else{
		fmt.Println(0)
	}

}

// func SortString(s string) string {
//     r := []rune(s)
//     sort.Sort(sortRunes(r))
//     return string(r)
// }