package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) { // named return values
	x = sum * 4 / 9
	y = sum - x
	return // naked returns should only be used in short functions as to not arm readability
}

var i, j int = 1, 2 // variable with initializers and type omitted (inferred)

const formula = "1 + 2 ="

func main() {
	sendMessage := true // short assignment statements can only be used in function bodys
	a, b := swap("GO!", "Hello")

	if sendMessage {
		fmt.Println(a, b)
		fmt.Printf("Type of a is: %T\n", a)
		rand.Seed(2)
		fmt.Println("My favorite number", rand.Intn(10))
		fmt.Printf("Square root %v yo\n", math.Sqrt(9))
		fmt.Println(formula, add(1, 2))
		fmt.Println(split(17))
	}

	var c, python, java = true, false, "no!" // initialized vars without `int`
	fmt.Println(i, j, c, python, java)
}

// to run: go run main.go
// to generate mod file: go mod init
// to compile: go build
// ro run compiled code: ./hello-go
// := syntax allows you to emit the type as go can infer it
