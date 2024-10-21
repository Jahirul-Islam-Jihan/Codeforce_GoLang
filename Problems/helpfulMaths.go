package main
import(
	"fmt"
	"sort"
)
func main(){
	var str string
	fmt. Scan(&str)
	sortString := SortString(str)
	length:=(len(sortString))/2
	for i:=length; i<len(str); i++{
		fmt.Print(string(sortString[i]))
		
		if i!= len(str)-1{
			fmt.Print("+")
		} 
		
	}

}

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}