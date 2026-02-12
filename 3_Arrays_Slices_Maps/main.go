package main

import "fmt"

func main() {
	/*
		1) ARRAYS
		 - fixed size
		 - same type
		 - static
		 - contguous
	*/

	// declaration
	var strs [2]string
	strs[0] = "hello"
	strs[1] = "world"
	fmt.Println(strs)
	
	// declaration + initialization
	nums := [5]int{10, 20, 30, 40, 50} 
	// or var nums [5]int = [5]int{10, 20, 30, 40, 50}
	// or num := [...]int{10, 20, 30, 40, 50} // compiler will count the size automatically
	fmt.Println(nums)

	// slice a array
	slicedNums := nums[1 : 4]  // from index 1 to index 3 (last one is excluded)
	fmt.Println(slicedNums)
	// IMP -> slicing does not create a new array, it just creates a view over the existing array
	// any changes in the sliced array will reflect in the original array & vice versa. Eg ->
	slicedNums[2] = 1000 
	fmt.Println(nums) // nums is changed at index 3


	/*
		2) SLICES
		 - a dynamic data structure that is built on top of an array
		 - dynamic size
		 - same type
		 
		VIMP -> When you define a slice variable in Go, such as var mySlice []int, 
		Go creates a data structure (internally Go slice is basically a struct often called slice header)
		that contains three pieces of information:

		i) A pointer to an underlying array that holds the slice elements
		ii) A length that represents the number of elements in the slice
		iii) A capacity that represents the maximum number of elements that the underlying array can hold

		Eg -> 
		mySlice := []int{2, 3,4}

		go creates this under the hood_ 

		+---------+---+
		| 2| 3| 4|
		+---------+--- +
		len=3, cap=4
	*/

	// initialization is kind of same as arrays
	intSlice := []int{1, 2, 3} // Or var intSlice []int = []int{1, 2, 3}	
	// another type of initialization
	intSlice2 := make([]int, 3, 6)   
	                // type, len, cap
	fmt.Println(intSlice2)

	// append elements to slice
	intSlice = append(intSlice, 4)
	// VIMP -> The append built-in function appends elements to the end of a slice. 
	// If it has sufficient capacity, the destination is resliced to accommodate the new elements. 
	// If it does not, a new underlying array will be allocated. Append returns the updated slice.
	// TC is amortized O(1)

	// emptySlice := []string{} 
	
	// functions ->
	// len() -> return no. of elements in slice 
	// cap() -> return size of slice 

	// spread operator in go as in js/ts ->
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5}
	arr1 = append(arr1, arr2...)

	// IMP 
	slower := make([]int, 0)
	faster := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		slower = append(slower, i)
	}
	for i := 0; i < 10000; i++ {
		faster = append(faster, i)
	}
	// appending els to the first one will be slower
	// Reason -> 2nd array has capacity of pre initialized, so no need to reallocate memory 

	/*
		3) MAPS
		 - <key -> value> pairs
	*/
	
	mpp := make(map[string]int)  // or var mpp map[string]int = make(map[string]int)
				 // map[<key>]<value>
	
	// add element
	mpp["apple"] = 100
	mpp["banana"] = 200
	fmt.Println(mpp)  // map[apple:100 banana:200]

	// delete element
	delete(mpp, "apple")

	// check if key present ?
	_, isThere := mpp["apple"]  // blank identifier to ignore the key
	if isThere {
		fmt.Println("present")
	} else {
		fmt.Println("not present")
	}
	// IMP
	fmt.Println(mpp["apple"]) // returns 0 
	// in go maps, if u try to access a key which is not present, 
	// it returns the default value of the value type. 
	// Eg -> for a map which is <int, string>, it returns "".

	/*
		4) LOOPS
	*/

	// for loop
	for i := 0; i < 10; i++ {  // NOTE_ traditional way of writing loops with () does not exist in go  
		// do something
	}

	// while loop
	// go does not have while loop, but u can achieve the same using for loop like this ->
	i := 10
	for i >= 0 {
		// do something
		i--
	}

	// for-each loops
	// 1) Arrays or slices_
	for index, str := range strs {
		fmt.Printf("Index: %v, Value: %v \n", index, str)
		// NOTE_ u can use %v which is called versatile specifier 
		// used to print the value of any data type in its default format
	}

	// 2) maps_
	for name, age := range mpp {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}
	
}