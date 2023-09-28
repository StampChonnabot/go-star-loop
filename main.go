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
	fmt.Println("-------------------------")

	fmt.Println("Diamond Star part2")
	diamondStarII(num)
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
	fmt.Print(result)
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
	fmt.Print(result)
}

func diamondStarII(num int) {
	var i, k, j int

	result := [][]string{}
	if num > 0 {
		for i = 1; i <= num; i++ {
			row := []string{}
			for k = num - i; k >= 1 && i != num; k-- {
				row = append(row, " ")
			}
			for j = 1; j < i+i; j++ {
				if j+num-1 == num || j+1 == i+i {
					row = append(row, "*")
				} else {
					row = append(row, " ")
				}
			}
			row = append(row, "\n")
			result = append(result, row)
		}
		for i = num - 1; i >= 1; i-- {
			row := []string{}
			for k = num - i; k >= 1 && i != num; k-- {
				row = append(row, " ")
			}
			for j = 1; j < i+i; j++ {
				if j+num-1 == num || j+1 == i+i {
					row = append(row, "*")
				} else {
					row = append(row, " ")
				}
			}
			row = append(row, "\n")
			result = append(result, row)
		}
	}

	for i := 0; i < len(result); i++ {
		for _, star := range result[i] {
			fmt.Print(star)
		}
	}
}
