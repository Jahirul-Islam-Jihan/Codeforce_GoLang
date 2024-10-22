package main

import "fmt"

func main(){

	var n int 
	var str string
	count:=0
	fmt.Scan(&n,&str)
	for i:=0;i<n-1;i++{
		if str[i]==str[i+1]{
			count++
		}
	}
	fmt.Println(count)
}