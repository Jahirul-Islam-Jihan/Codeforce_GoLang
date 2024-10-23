package main
import "fmt"
func main() {
	var str string
	count := 0
	fmt.Scan(&str)
	length := len(str)
	for i := 0; i < length-1; i++ {
		if str[i] == str[i+1] {
			count++
		} else if str[i] != str[i+1] && count < 6 {
			count = 0
		}
	}
	if count >= 6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}