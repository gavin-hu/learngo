// range
package main

import (
	"fmt"
)

func main() {
	//
	a := []int {1, 2, 3, 4, 5}
	for i, e := range a {
		fmt.Println(i, e)
	}
	//
	m := map[string]string{"name":"gavin", "age":"18"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	//
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	
}
