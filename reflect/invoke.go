package main

import (
	"reflect"
	"fmt"
	"time"
)

type YourT1 struct {}
func (y YourT1) MethodBar() {
	//do something
	fmt.Println("method invoked too")
}

type YourT2 struct {}
func (y YourT2) MethodFoo(i int, oo string) {
	//do something
	fmt.Println("method invoked")
}

func Invoke(any interface{}, name string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func main() {
	begin := time.Now()
	for i:=0; i<100000; i++ {
		Invoke(YourT2{}, "MethodFoo", 10, "abc")
		/*obj := YourT2{}
		obj.MethodFoo(10, "abc")*/
		//Invoke(YourT1{}, "MethodBar")
	}
	//
	elapsed := time.Since(begin)
	//
	fmt.Println(elapsed)
}