package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		1) STRINGS
	*/

	// in go, strings are actually an ARRAY OF BYTES
	// they are represented using UTF-8 encoding by default
	var str string = "résumé"
	var indexed = str[0];
	fmt.Printf("%v %T \n", indexed, indexed)  // %T to print type   
	// o/p -> 114 uint8
	// this is actually the byte value and byte type (uint8) respectively

	for ind, ch := range str {
		fmt.Println(ind, ch)
	}
	// o/p is actually weird ->
	// 0 114
	// 1 233
	// 3 115
	// 4 117
	// 5 109
	// 6 233

	// how str is stored ->
	//  r        é                 s        u        m        é
	// [01110010 11000011 10101001 01110011 01110101 01101101 11000011 10101001] 
	// -> UTF-8 encoding of each ch, this is how it is stored in memory

	// [114      233               115      117      109      233] 
	// -> unicodes of each ch

	// 1) in case of indexing ->
	firstbyte := str[1]	
	fmt.Println(firstbyte) // o/p -> 195
	// reason -> 'é' is 11000011 10101001 & first byte is 11000011 which is 195 in decimal
	// IMP -> indexing a string always gives the FIRST BYTE of the character which is in UTF-8 encoding  

	// 2) in case of for-each loop ->
	// here, go automatically decodes the UTF-8 encoded bytes to give the UNICODE value of each character
	// so for 'é', it gives 233 which is the unicode value of 'é'
	// In the for-each loop above, index 2 is skipped coz 'é' takes 2 bytes

	/*
		2) RUNES
	*/

	var runeExample rune = 'A';
	fmt.Println(runeExample)  // 65
	// rune stores the unicode value of a character
	// rune is just an alias for int32

	runeStr1 := []rune{'r', 'é', 's', 'u', 'm', 'é'}
	fmt.Println(runeStr1)  // [114 233 115 117 109 233]

	// to convert string to rune slice, just use type-casting ->
	runeStr2 := []rune(str)
	fmt.Println(runeStr2)  

	/*
		3) STRING BUILDER
	*/
	
	// strings are IMMUTABLE in go
	// str[0] = 'x' // wont work
	
	// so we use string builder package 
	var builder strings.Builder
	var sentence string = ""
	var words []string = []string{"Go", "is", "awesome"}
	for _, word := range words {
		// sentence += word + " "  // this creates a new string in memory each time, inefficient
		builder.WriteString(word)
		builder.WriteString(" ")
		// this is effecient
	}
	sentence = builder.String()
	fmt.Println(sentence)  // o/p -> Go is awesome

	
}