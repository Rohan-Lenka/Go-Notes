package main

import (
	"fmt"
	"errors"
)  // do multiple imports like this

func sayHi(name string) string {  // return type shld be mentioned
	return "hi " + name
}

func fn(n int, d int) (int, int, error) {  // multiple return types
	var err error  // by default it is nil, nil is just like null
	// if statements can be written like this too in go or just use traditional way
	if d == 0 {
		err = errors.New("division by zero not allowed")
		return 0, 0, err
	}
	q := n / d
	r := n % d
	return q, r, err
}

func main() {
	fmt.Println(sayHi("John"))	
	q, r, err := fn(10, 3)
	if err != nil {
		fmt.Println("Error:", err.Error())  // print the error message like this using Error()
		return
	}
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	// functions inside functions & anonymous functions 
	func () {
		fmt.Println("Anonymous function")
	}() // declaration + execution is compulsory for anonymous functions

	sum := func (a int, b int) int {  // the anonymous function is assigned to var sum
		return a + b
	}
	fmt.Println(sum(1, 9))
}