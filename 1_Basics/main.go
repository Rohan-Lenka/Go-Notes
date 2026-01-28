/* 
	1) FEATURES
	
	- statically typed language (vars must have a type)
	- strongly typed language (operations u perform depends on the types of the vars)
	  Eg -> 1 + "1" IS NOT ALLOWED
	- compiled language and produces a binary file (.exe file) after compilation (hence faster)
	- built in concurrency i.e parallelism in go is builtin and no work arounds are needed
	- simplicity
	- garbage collector to freeze up unused memory

	2) FILE STRUCTURE 

	- package ->                     vs          module ->
	  collection of go files                     collection of go packages

*/


package main 
/*  1) special package in go which tells the compiler that 
	this is an executable program and not a library 
	2) every go program must have a main package
*/

import "fmt"

func sayHi() {
	fmt.Println("hi...")
}

func main() {
	// every executable Go program starts executing from main() function
	// Only one main function is allowed in a program
	fmt.Println("hello world")

	/* 
		2) DATA TYPES 

		i) int8, int16, int32, int64, int (by defaut it is int32)
		ii) uint8, uint16, uint32, uint64
		iii) float32, float64 (for more precision)
		iv) bool
	*/

	// v) Strings ->
	// just like js/ts, u can use "" or `` to define strings (not '')
	
	// string functions ->

	// 1) len() [IMP]
	var str string = "hello world"
	fmt.Println(len(str)) // returns the number of bytes not no. of characters
	
	// When len() is fine? 
	// Your string is ASCII only (English letters, numbers, symbols)

	// When not ?
	// Your string contains Unicode / non-ASCII characters
	// eg_ var hindiHi string = "हाय"
	// for such cases, use the utf8 package like this ->
	// import "unicode/utf8"
	// utf8.RuneCountInString(hindiHi)

	// NOTE -> more about runes in go later

	/* 
		3) DEFAULT VALUES

		i) all int = 0
		ii) bool = false
		iii) string = "" (empty string)
	*/

	/*
		there is const declaration of variables too in go
		they behave like const in js/ts
	*/

	/* 
		some more key points
	*/
	// i) automatic type inference
	var pi = 3.14  
	fmt.Println(pi)
	// ii) short-hand variable delaration
	a := 12  
	fmt.Println(a)
	// iii) multiple variable declaration
	b, c := 1, "bye"  // or var b, c = 1, "bye" 
	fmt.Println(c)
	fmt.Println(b)
}