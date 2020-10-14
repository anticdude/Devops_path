package main
import "fmt"
func main() {
	var(
		a = 2
		b = 5
		c = 8
		d = c * a
		e = a - b
		f = d / e
	)
	fmt.Println("A is = ",a)
	fmt.Println("A is = ",b)
	fmt.Println("A is = ",c)
	fmt.Println("A is = ",d)
	fmt.Println("A is = ",e)
	fmt.Println("A is = ",f)
//		"b is = "b
//		"c is = "c
//		"* of c and a is = "d
//		"- of a and b is = "e
//		"/ of d and e is = "f )
}
