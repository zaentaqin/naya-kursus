package main

import "fmt"

func main() {
	ch := make(chan int)

	go action(ch, 1)

	go reactionLike(ch)
	go reactionComment(ch)

	for range ch {
	}
}

func action(ch chan<- int, a int) {
	ch <- a
	close(ch)
}

func reactionLike(ch <-chan int) {
	if <-ch == 0 {
		fmt.Println("Like")
	} else {
		return
	}
}

func reactionComment(ch <-chan int) {
	if <-ch == 1 {
		fmt.Println("Comment")
	}
}
