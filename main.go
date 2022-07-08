package main

import (
	"fmt"
	"dasar_dasar_GO/calculation"
)

func main()  {
	fmt.Println("halo, belajr Golang")

	result := calculation.Add(8, 9)
	result1 := calculation.Multiply(8, 9)

	fmt.Println(result)
	fmt.Println(result1)
}