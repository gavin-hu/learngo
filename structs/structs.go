// structs
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Say(message string) {
	fmt.Println(p.Name, " Say ", message)
}

func main() {
	//
	p := new(Person)
	p.Name = "gavin"
	p.Age = 18
	//
	fmt.Println(p)
	//
	p.Say("hello world")
}
