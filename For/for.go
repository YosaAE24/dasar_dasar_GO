package main

import "fmt"

func main() {
	title := "Aku sedang belajar Golang"

	for index, letter := range title {
		if index % 2 == 0 {
			fmt.Println("index :", index, "letter :", string(letter))
			fmt.Println("===========================================")
		}
	}

	judul := "Aku sEnang sekalI belajar gOlang"

	for index, letter := range judul {
		switch letter {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'U', 'I', 'E', 'O':
			fmt.Println("index: ", index, "letter", string(letter))	
		}
	}
}