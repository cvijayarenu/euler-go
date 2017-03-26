package main

const buggerSize int = 10

func fibGenerator() <-chan int {
	ch := make(chan int, buggerSize)
	go func() {
		for i, j := 1, 1; ; i, j = i+j, i {
			ch <- i
		}
	}()

	return ch
}
