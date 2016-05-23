// pointers
package main

import (
	"fmt"
	"runtime"
)

func zeroval(v int) {
	v = 0
}

func zeroptr(v *int) {
	*v = 0
}

func main() {
	i:=1
	//
	zeroval(i)
	fmt.Println(i)
	//
	zeroptr(&i)
	fmt.Println(i)
	//
	fmt.Println("Pointer ", &i)
	//
}
