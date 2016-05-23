// interfaces
package main

import (
	"fmt"
)

type Animal interface {
	Say()
}

type Person struct {
}

func (p Person) Say() {
	fmt.Println("hi...")
}

type Dog struct {
	
}

func (d Dog) Say() {
	fmt.Println("wo...")
}

func say(animal Animal) {
	animal.Say()
}

func main() {
	//
	p := new(Person)
	d := new(Dog)
	//
	say(p)
	say(d)
}
