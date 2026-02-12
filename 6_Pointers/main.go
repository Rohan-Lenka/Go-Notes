package main

import "fmt"

func main() {
	/*
		1) BASICS IN POINTERS
	*/

	var ptr *int32 = new (int32)  // this means ptr is a pointer which will hold the address for a int32
	// this part new (int32) means -> 
	// memory is allocated for an int32, default value 0 is given to it & its address is returned  

	var emptyPtr *int32  // by defautlt a pointer stores nil if it is not referencing any variable yet
	fmt.Printf("%v \n", emptyPtr)   
	
	fmt.Printf("The value at pointer ptr is %v \n", *ptr)  // prints 0   
	// the value of the var referenced by the pointer can be accessed like this -> *ptr
	
	// fmt.Printf("%v \n", *emptyPtr)   
	// the above line will throw runtime error, coz you're trying to access memory through a nil pointer.

	integer := 200
	p := &integer  // or var p *int32 = &integer
	fmt.Printf("The value of integer variable is: %v \n", *p)

	// IMP
	k := 100
	integer = k // what happens here is that the value of k is copied into integer's memory location
	var slice []int32 = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 10  // slice array also changes 
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// Reason -> (same as in js)

	/*
		2) POINTERS IN FUNCTIONS
	*/
	
	i := 10
	change(&i)
	fmt.Println(i)
	
	/*
		3) POINTERS IN FUNCTIONS INVOLVING SLICES (VIMP)
	*/
	
	// https://medium.com/@ansujain/understanding-slices-in-go-pass-by-value-and-pass-by-pointer-9a52830c741e
	
}

func change(i *int) {
	*i = (*i) * 2
}