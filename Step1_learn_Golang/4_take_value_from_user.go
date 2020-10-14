package main
import "fmt"
func main() {
	fmt.Print("enter your money you want to double ;) ")
	var input float64
	fmt.Scanf("%f", &input)
	var output = input * 2
	fmt.Println("21 din me paisa double : ",output)
}
