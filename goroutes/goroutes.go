package main

func main() {
	c := make(chan int, 10)
	//
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()
	//
	for {
		v := <-c
		print(v)
		if v==99 {
			break
		}
	}
}
