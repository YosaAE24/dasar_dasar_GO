package main

import (
	"errors"
	"fmt"
)

// import "runtime"

func main() {
	// scores := []int{10, 5, 8, 9 ,7}
	// total := sum(scores)
	// fmt.Println(total)

	result, err := calculate(5, 6, "*")
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

// func sum(score []int) int {
// 	var total int
// 	for _, scores := range score{
// 		total = total + scores
// 	}
// 	return total
// }

func calculate(number, numbertwo int , operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numbertwo
	case "-":
		result = number - numbertwo
	case "*":
		result = number * numbertwo
	case "/":
		result = number / numbertwo
	default:
		errorResult = errors.New("unknown opration")
	}

	return result, errorResult
}