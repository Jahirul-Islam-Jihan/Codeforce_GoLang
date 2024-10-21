package main
import (
	"fmt"
	//"strings"
)
func main(){
	var str string
	fmt.Scan(&str)
	if str[0]>= 'a' && str[0] <= 'z'{
		result:=string(str[0]-32)
		fmt.Print(result)
	}else{
		fmt.Print(string(str[0]))
	}
	for i:=1 ; i<len(str);i++{
		fmt.Print(string(str[i]))
	}
	fmt.Println("")

}