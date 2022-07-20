package main

import (
	"dasar_dasar_GO/calculation"
	"fmt"
)

func main()  {
	fmt.Println("halo, belajr Golang")

	result := calculation.Add(8, 9)
	result1 := calculation.Multiply(8, 9)

	fmt.Println("Ini penjumlahan", result)
	fmt.Println("Ini perkalian", result1)
}