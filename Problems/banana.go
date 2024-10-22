package main

import "fmt"

func main(){
	var w,k,n int
	cost:=0
	fmt.Scan(&k,&n,&w)
	for i := 1; i <= w; i++ {
		cost=cost+(i*k)
	}
	if(cost>n){
		fmt.Println(cost-n)
	}else{
		fmt.Println(0)
	}
	
}