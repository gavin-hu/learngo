package main

// Callback function
func each(elements []string, callback func(string)) {
	for _, element := range elements {
		callback(element)
	}
}

// Callback interface
func each2(elements []string, callback Callback) {
	for _, element := range elements {
		callback.Handle(element)
	}
}

type Callback interface {

	Handle(string)

}

type DefaultCallback struct {

}

func (dc *DefaultCallback) Handle(value string) {
	println(value)
}


func main() {
	//
	var elements = []string{"gavin", "aimy", "arron"};
	each(elements, func(element string){
		println(element)
	})
	//
	each2(elements, &DefaultCallback{})
}
