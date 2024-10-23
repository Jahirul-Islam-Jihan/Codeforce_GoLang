package main
import "fmt"
func main() {
	var str string
	count := 0
	fmt.Scan(&str)
	length := len(str)
	if length < 7 {
		fmt.Println("NO")
	} else {
		for i := 0; i < length-1; i++ {
			if str[i] == str[i+1] {
				count++
				if count == 6 {
					fmt.Println("YES")
					return 
				}
			} else {
				count = 0
			}
		}
		fmt.Println("NO")
	}
}
