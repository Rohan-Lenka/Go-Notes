package main

import "fmt"

/*
	1) STRUCTS
*/

type Student struct {
	name string 
	age uint8
	class Class  // struct inside struct  
}

type Class struct {
	standard uint8
	section rune
}

// struct methods
func (student Student) getSection() string {
	return string(student.class.section)
}
// everything is same as a normal function only thing to add is the (student Student) part
// this allows the structs to access the method like this struct.method() 

/*
	2) INTERFACES
*/

type Lion struct {
	name string
	speed uint8
}

type Tiger struct {
	name string
	speed uint8
}

func (lion Lion) isFast() bool {
	return lion.speed > 50
}

func (tiger Tiger) isFast() bool {
	return tiger.speed > 50
}

type Animal interface {
	// can pass common methods inside  
	isFast() bool
}

func isAnimalFast(animal Animal) bool {
	return animal.isFast()
}

func main() {
	// structs
	var student Student
	student.name = "Rohan"
	student.age = 15    
	student.class.standard = 10 
	student.class.section = 'C' 
	fmt.Println(student)
	studentSection := student.getSection()
	fmt.Println(studentSection)

	john := Student{"John", 16, Class{11, 'A'}}
	fmt.Println(john)

	// anonymous struct
	var animal = struct {
		name string 
		legs uint8
	}{"Lion", 4}
	// advantage -> define and initalize a struct immediately
	// dis-advantage -> not reusable
	fmt.Println(animal)

	// interfaces
	lion1 := Lion{"lion1", 60}
	tiger1 := Tiger{"tiger1", 40}
	isLionFast := isAnimalFast(lion1)
	isTigerFast := isAnimalFast(tiger1)
	fmt.Printf("%v & %v \n", isLionFast, isTigerFast)
}