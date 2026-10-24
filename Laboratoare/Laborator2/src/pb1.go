package main

import "fmt"

var ch = make(chan string)

func hello_world() {
	ch <- "hello world!"
}
func main() {
	go hello_world()
	ms, _ := <-ch

	fmt.Println(ms)
}
