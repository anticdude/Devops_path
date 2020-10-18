package main
import "fmt"
func main() {
	var a = 2
	var b = 2
	if a == b {
		var c =	b / a
		fmt.Println(c)
	}else { // if a >= b {
		var c = b * a
		fmt.Println(c)
	}

}
