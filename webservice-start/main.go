package main

import (
	"net/http"

	"github.com/alexb760/training-go/webservice-start/controller"
)

func main() {
	controller.RegisterControllers()
	//Front Controller - Back-Controller pattern
	http.ListenAndServe(":3000", nil)
	//looping()
}

// func looping() {
// 	//traditional loop
// 	var i int = 5
// 	for i < 5 {
// 		println(i)
// 		i++
// 	}

// 	//Implicit definition.
// 	for j := 7; j < 7; j++ {
// 		println(j)
// 	}

// 	// Both lopps we can use break and continue wey words.
// 	// nfinite loop
// 	var k int
// 	println("==Infinitive loop==")
// 	for {
// 		//just added this to demostrate and breack this and not my machine
// 		if k == 8 {
// 			break
// 		}
// 		println(k)
// 		k++

// 	}
// }
