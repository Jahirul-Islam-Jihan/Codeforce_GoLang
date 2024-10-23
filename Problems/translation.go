package main

import "fmt"

func main(){
	var str1, str2 string
	fmt.Scan(&str1,&str2)
	if str2==reverseString(str1){
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}

}
func reverseString(s string) string{
	length:=len(s)
	reversed:=make([]byte, length)
	for i := 0; i < length; i++ {
		reversed[i]=s[length-i-1]
	}

	return string(reversed)
}