// slices
package main

import (
	"fmt"
)

func main() {
	a := make([]string, 2)
	//
	a[0] = "s1";
	a[1] = "s2";
	//
	a = append(a, "s3","s4", "s5", "s6")
	//
	fmt.Println("Print a ", a)
	fmt.Println("Print a len ", len(a))
	//
	b := make([]string, len(a))
	copy(b, a)
	//
	fmt.Println("Print b ", b)
	fmt.Println("Print b len ", len(b))
	//
	fmt.Println("Print Slice ", b[1:2])
	fmt.Println("Print Slice ", b[:2])
	fmt.Println("Print Slice ", b[1:])
	//
	c := make([][]int, 2)
	for i:=0; i<2; i++ {
		innerLen := i+1
		c_inner := make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			c_inner[j] = j
		}
		c[i] = c_inner;
	}
	fmt.Println("Print c ", c)
}
