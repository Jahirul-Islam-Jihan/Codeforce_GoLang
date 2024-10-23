package main

import "fmt"

func main(){
	var n,k int
	digit:=0

	fmt.Scan(&n,&k)

	for i:=0;i<k;i++{
		digit=n%10
		if(digit==0){
			n=n/10
		}else{
			n=n-1
		}
	}
	fmt.Println(n)
}