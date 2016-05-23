// functions
package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b	
}

func divide(a float32, b float32) float32 {
	return a / b
}

func multi_values() (int, int) {
	return 1 , 2
}

func sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func next_seq() func() int {
	id := 0;
	return func() int {
		id += 1
		return id
	}
}

func plus_plus(a int, b int) int {
	if a < 1000 && b < 1000{
		a = a + b
		b = a + b
		//
		return plus_plus(a, b)
	}
	return a + b
}

func main() {
	//
	fmt.Println("a + b = ", plus(1, 2))
	a, b := multi_values()
	//
	fmt.Println("Multi Return Values : ", a, b)
	//
	fmt.Println("Sum 1 3 4 5", sum(1, 3, 4, 5))
	//
	seq := next_seq()
	//
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	//
	fmt.Println(plus_plus(100, 10))
}
