package main

import "fmt"

func main() {
	num := 0
	fmt.Printf("Enter Number : ")
	fmt.Scanf("%d", &num)
	fmt.Println("-------------------------")

	fmt.Println("Right Triangle Star")
	rightTriangleStar(num)
	fmt.Println("-------------------------")

	fmt.Println("Diamond Star")
	diamondStar(num)
}

func rightTriangleStar(num int) {

	var result string
	if num > 0 {
		for i := 1; i <= num; i++ {
			for j := 1; j <= i; j++ {
				result += "*"
			}
			result += "\n"
		}
	}
	fmt.Printf(result)
}

func diamondStar(num int) {
	var i, k, j int
	var result string
	if num > 0 {
		for i = 1; i <= num; i++ {
			for k = num - i; k >= 1 && i != num; k-- {
				result += " "
			}
			for j = 1; j < i+i; j++ {
				result += "*"

			}
			result += "\n"
		}
		for i = num - 1; i >= 1; i-- {
			for k = num - i; k >= 1 && i != num; k-- {
				result += " "
			}
			for j = 1; j < i+i; j++ {
				result += "*"
			}
			result += "\n"
		}
	}
	fmt.Printf(result)
}
