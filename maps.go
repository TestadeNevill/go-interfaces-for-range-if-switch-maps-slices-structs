package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Testa",
		LastName:  "De Nevill",
	}

	myMap["me"] = me
	log.Println(myMap["me"].FirstName, myMap["me"].LastName)

}
