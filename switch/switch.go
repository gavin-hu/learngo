// switch
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	//
	switch i {
		case 1 :{
			fmt.Println("One")
		}
		case 2 :{
			fmt.Println("Two")
		}
	}
	//
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday : {
			fmt.Println("idleday")
		}
		default : {
			fmt.Println("workday")
		}
	}
	//
	t := time.Now()
	switch {
		case t.Hour() < 12 : {
			fmt.Println("it's before noon")
		}
		default : {
			fmt.Println("it's after noon")
		}
	}
}
