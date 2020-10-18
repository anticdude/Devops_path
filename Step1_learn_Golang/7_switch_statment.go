package main
import "fmt"
func main() {
	fmt.Print("Please enter last digit of your mobile and we will give you a nice name : ")
	var num int64
	fmt.Scanf("%d", &num)
	fmt.Print("Please enter your gender F/M : ")
	var gen string
	fmt.Scanf("%s", &gen)
	if gen == "M" {
		switch num {
		case 0: fmt.Println("Jack")
		case 1: fmt.Println("babu")
		case 2: fmt.Println("charm")
		case 3: fmt.Println("king")
		case 4: fmt.Println("devil")
		case 5: fmt.Println("tresu")
		case 6: fmt.Println("Lalu")
		case 7: fmt.Println("bako")
		case 8: fmt.Println("jaan")
		case 9: fmt.Println("chikudo")
		}
	}else if gen == "F" {
		switch num {
                        case 0: fmt.Println("lazy")
                        case 1: fmt.Println("diku")
                        case 2: fmt.Println("baku")
                        case 3: fmt.Println("chaku")
                        case 4: fmt.Println("laku")
                        case 5: fmt.Println("daku")
                        case 6: fmt.Println("angu")
                        case 7: fmt.Println("kidi")
                        case 8: fmt.Println("jaanudi")
                        case 9: fmt.Println("laali")
		}
	}else {
		fmt.Print("Please select proper gender")
	}
	

}
