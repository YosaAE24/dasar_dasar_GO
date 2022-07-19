package main

import "fmt"

func main() {
	// var languages [...]]string
	// languages[0] = "Go"
	// languages[1] = "Python"
	// languages[2] = "Rust"
	// languages[3] = "C#"
	// languages[4] = "Ruby"

	// fmt.Println(languages)

	bahasa := [...]string {"go", "ruby", "python", "rust", "c#", "javascript", "php"}

	for index, lang := range bahasa {
		fmt.Println("index : ", index, "bahasa : ", lang)
		
	}
}