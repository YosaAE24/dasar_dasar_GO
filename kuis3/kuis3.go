package main

import "fmt"

func main()  {
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// var total int
	
	// for _, score := range scores {
	// 	total = total + score
	// }
	// lenght := len(scores)
	// average := float64(total) / float64(lenght) 
	// fmt.Println(average)

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int
	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)

}