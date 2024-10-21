package main

import (
	"fmt"
	"sort"
)
func main(){
	var str string
	fmt.Scan(&str)
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool{
		return bytes[i]<bytes[j]
	})
	var count = 0
	sortStr := string(bytes)
	length := len(sortStr)-1
	for i:=0 ; i<length; i++{
		if sortStr[i]!=sortStr[i+1]{
			count++
		}
	}
	if (count+1)%2 == 0{
		fmt.Println("CHAT WITH HER!")
	}else{
		fmt.Println("IGNORE HIM!")
	}
}