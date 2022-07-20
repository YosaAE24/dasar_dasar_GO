package main

import "fmt"

func main() {
	// sentence := printMyResult("saya sedang")
	// fmt.Println(sentence)

	luas, keliling := calculate(5, 8)
	fmt.Println(luas)
	fmt.Println(keliling)
}

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " belajar golang"
// 	return newSentence
// }

func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}