package main

import (
	"fmt"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	IsActive  bool
}

func main() {
	fmt.Println("Hello Belajar Golang")

	user := User{1, "Nur", "Ichsan", "iccankblog@gmail.com", true}
	user1 := User{1, "Ika", "Apriliya", "ika@gmail.com", true}

	fmt.Println(user)
	fmt.Println(user1)

	result := displayUser(user)
	fmt.Println(result)

	// How to iterate through an array (Use _ if a value isn't used)

	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cariBilanganBulat(numSlice)

}

func displayUser(user User) string {
	return fmt.Sprintf("Nama lengkap : %s %s, Email: %s", user.Firstname, user.Lastname, user.Email)
}

func cariBilanganBulat(numbers []int) {
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
	}
}
