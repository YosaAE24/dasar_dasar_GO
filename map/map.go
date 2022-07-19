package main

import "fmt"

func main()  {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 7
	// myMap["javascripts"] = 9
	// myMap["go"] = 10

	// fmt.Println(myMap)

	myKap := map[string]string {
		"ruby": "cantik", 
		"go": "sangat cepat", 
		"javascript": "sungguh kuat",
	}
	for index, kap := range myKap {
		fmt.Println(index, kap)
	}
	delete(myKap, "ruby")
	fmt.Println("====================================")

	valeus, apaada := myKap["python"]

	fmt.Println(apaada)
	fmt.Println(valeus)

	for index, kap := range myKap {
		fmt.Println(index, kap)
	}
}