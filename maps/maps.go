// maps
package main

import (
	"fmt"
)

func main() {
	//
	m := make(map[string]string)
	//
	m["name"] = "gavin"
	m["age"] = "18"
	//
	fmt.Println("Print map ", m)
	//
	fmt.Println("Print name ", m["name"])
	//
	fmt.Println("Print size ", len(m))
	//
	delete(m, "age")
	//
	fmt.Println("Print map ", m)
	//
	for k,v := range m {
		fmt.Printf("Key %s | Value %s\n", k, v)
	}
}
