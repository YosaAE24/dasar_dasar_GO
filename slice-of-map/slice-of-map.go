package main

import "fmt"

func main()  {
	students := []map[string]string{
		{"name": "Yosa", "score": "90"},
		{"name": "sincan", "score": "60"},
		{"name": "ishaq", "score": "70"},
		{"name": "ulya", "score": "80"},
		{"name": "frendy", "score": "85"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}