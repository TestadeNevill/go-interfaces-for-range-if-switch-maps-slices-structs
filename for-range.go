package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Testa", "De Nevill", "T@gmail", 30})
	users = append(users, User{"Director", "of Works", "Dow@gmail", 30})
	users = append(users, User{"Test", "Nick", "TN@gmail", 100})
	users = append(users, User{"Melissa", "De Nevill", "MM@gmail", 27})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
