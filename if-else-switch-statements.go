package main

import "log"

// func main() {
// 	myNum := 100
// 	isTrue := false

// 	if myNum > 99 && !isTrue {
// 		log.Println("myNum is greater than 99 and isTrue is set to true")
// 	} else  if myNum < 100 || !isTrue {
// 		log.Println("myNum is greater less than 100 and isTrue is false")

// 	}
// }

func main() {
	myVar := "moster"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("dog is set to dog")
	case "fish":
		log.Println("fish is set to fish")

	default:
		log.Println("cat is set something else")
	}
}
