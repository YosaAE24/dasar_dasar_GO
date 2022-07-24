package main

import (
	"fmt"
)

type User struct {
	id int
	firstName string
	lastName string
	email string
	isActive bool
}
type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func (group Group) displayGroup()  {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println(" ")
	fmt.Printf("Member Count: %d", len(group.Users))
	fmt.Println(" ")

	fmt.Print("Name Users: ")
	for _, user := range group.Users {
		fmt.Print(user.firstName, " ", user.lastName, ", ")
	}
}

func main() {
	user1 := User{
		//mengisi field nya secara langsung
		id: 1,
		firstName: "Yosa",
		lastName: "Adytia",
		email: "yosa@gamil.com",
		isActive: true,
	}

	//dan bisa mengisinya secara tidak langsung
	user2 := User{}
	user2.id = 2
	user2.firstName = "Adytia"
	user2.lastName = "Pratama"
	user2.email = "adit@gmail.com"
	user2.isActive = true
	//ada juga cara ketiga yakni tanpa memanggil propertinya seperti ID tetapi ketika mengisi harus urut dari Id sampai isActive

	// displayUser1 := displayUser(user1)
	// displayUser2 := displayUser(user2)

	users := []User{user1, user2}
	group := Group{"Gamer", user1, users, true}

	group.displayGroup()

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)
}

//struct juga bisa menjadi atribut dari struct yang lain


//struct juga bisa dipakai untuk sebuah parameter
// func displayUser(user User) string {
// 	result := fmt.Sprintf("Name: %s %s Email: %s", user.firstName, user.lastName, user.email)
// 	return result
// }
