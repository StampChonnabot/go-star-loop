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

	if num > 0 {
		for i := 1; i <= num; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	}

}

func diamondStar(num int) {
	var i, k, j int

	if num > 0 {
		for i = 1; i <= num; i++ {
			for k = num - i; k >= 1 && i != num; k-- {
				fmt.Printf(" ")
			}
			for j = 1; j < i+i; j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
		for i = num - 1; i >= 1; i-- {
			for k = num - i; k >= 1 && i != num; k-- {
				fmt.Printf(" ")
			}
			for j = 1; j < i+i; j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	}

}
