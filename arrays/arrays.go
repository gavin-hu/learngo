// arrays
package main

import (
	"fmt"
)

func main() {
	//
	var a [5]int;
	fmt.Println("Print a ", a);
	//
	a[4] = 100;
	fmt.Println("Print a[4] ", a[4]);
	//
	fmt.Println("Print a len ", len(a));
	//
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Print b ", b)
	//
	var c [2][2]int
	for i:=0; i<2; i++ {
		for j:=0;j<2;j++ {
			c[i][j] = i+j
		}
	}
	fmt.Println("Print c ", c)
}
