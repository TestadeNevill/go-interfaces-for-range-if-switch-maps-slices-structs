package main

import "log"

type myStruct struct {
	FirstName string
}

// reciver ties function to struct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Testa"

	myVar2 := myStruct{
		FirstName: "Melissa",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
