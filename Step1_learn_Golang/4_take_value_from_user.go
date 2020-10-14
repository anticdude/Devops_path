package main
import (
"fmt"
"time"
)
func main() {
	fmt.Print("enter your money you want to double ;) ")
	var input float64
	fmt.Scanf("%f", &input)
	var output = input * 2
	fmt.Println("21 din me paisa double : ",output) 
	var time string = time.Now().Format(time.Stamp)
	fmt.Println("Lakhpati banne ka time is ",time)
}
