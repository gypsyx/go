
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i string = "i am a string"
	j := "I am a string too"
	var k = "I am a third string"

	fmt.Println(i,"\n", j, "\n", k)
	fmt.Println("TypeOf(i) ", reflect.TypeOf(i))


	fmt.Println("Starting infinite loop")
	a := 0
	
	for {
		fmt.Printf("%d ", a)
		a++

		if a > 100 {
			break
		}
	}

	fmt.Printf("\n\n")

	// Another for loop, with init and condition
	for b:=0; b <= 50; b++ {
		fmt.Printf("%d ", b)
	}
	fmt.Printf("\n")

	//strings
	string_demos()
}

func string_demos() {
	var fname string = "Dinesh"
	var lname = "Ketumakha"
	middle := "P"
	fmt.Printf("\n%s .%s %s\n", fname, middle, lname)
}







