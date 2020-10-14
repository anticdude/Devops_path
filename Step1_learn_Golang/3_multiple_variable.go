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
	fmt.Println("value of a is = ",a)
	fmt.Println("value of b is = ",b)
	fmt.Println("value of c is = ",c)
	fmt.Println("value of d is = ",d)
	fmt.Println("value of e is = ",e)
	fmt.Println("value of f is = ",f)
}
