package main
import "fmt"
var z = 10
func main() {
	var a = "antic"
	fmt.Println(a)
	var b, c, d int = 1, 2, 3
	var e = b + c + d - z
	fmt.Println("total is", e)
}
