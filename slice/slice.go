package main

import "fmt"

func main()  {
	var gemingConsoles []string

	gemingConsoles = append(gemingConsoles, "PlayStation")
	gemingConsoles = append(gemingConsoles, "Nintendo Switch")
	gemingConsoles = append(gemingConsoles, "Xbox one")

	for _, console := range gemingConsoles {
		fmt.Println(console)
	}
}