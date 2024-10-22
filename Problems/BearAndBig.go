package main

import "fmt"

func main(){
	var a,b int
	count:= 0
	fmt.Scan(&a,&b)
	for true{
		if(a<=b){
			a=a*3
			b=b*2
			count++
		}else if (a>b){
			fmt.Println(count)
			break
		}
	}

}