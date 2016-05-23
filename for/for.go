// for
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("index ", i)
	}
	//
	for {
		fmt.Println("loop...")
		break
	}
}
