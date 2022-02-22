package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
}


func main() {
	
	person1 := new(Person)
	person1.name = "My name is Slim.."
	person1.age = 45

	// var person2 *Person = person1
	//	person2 := person1
	person2 := person1
	
	var person3 = new(Person)
	person3.name = "Susin"
	person3.age = 30

	//var person4 **Person
	//person4 = &person3
	person4 := &person3
	//	fmt.Println("Person address:&p\n", (obj Person))
	fmt.Println("Person1 address:&p\n", &person1)
	fmt.Println("Person2 :%p\n", person2)
	fmt.Println("Person2 address:%p\n", &person2)
	
	fmt.Println(person1)
	fmt.Println("Now using pointers:\n", (*person2).name)

	fmt.Println("Person 3:\n", person3)
	fmt.Println("Person 3 address:\n", &person3)

	fmt.Println("Person 4 :\n", person4)
	fmt.Println("Person 4 address:\n", &person4)
	fmt.Println("Person 4 *p:\n", *person4)

}
