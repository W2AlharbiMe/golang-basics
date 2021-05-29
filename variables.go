package main

import "fmt"

// global variable
var appName string = "Variables"

func main() {
	// strings
	// declare variable using ex
	var firstName string = "Abdullah"
	// read the type based on the value (infer type)
	var lastName = "Alharbi"
	// declare variable using short declaration
	mario := "mario"

	fmt.Println(appName + "\n");
	fmt.Println(firstName, lastName, "\n------------\n" + mario)

	// change variable value

	firstName = "Saleh"

	fmt.Println(firstName)


	// integers
	var x int = 20
	var y = 30
	z := x + y

	fmt.Printf("%v + %v = %v\n", x, y, z)


	// bits & memory
	// for more details about bit types https://golang.org/pkg/builtin/#float32
	// for more details about ranges https://golang.org/ref/spec#Numeric_types
	var numOne int8 = 25
	var numTwo int8 = -128
	// unsigned int
	var numThree uint8 = 255

	var scoreOne float64 = 89.98
	var scoreTwo float64 = 948490312.99
	scoreThree := 409.9


	fmt.Println(numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree)

}
