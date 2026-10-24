package main

import "fmt"

var ch = make(chan string)
var ch1 = make(chan string)

func send() {
	ch <- "hello world"
}

func receive() {

	var ms string

	ms = <-ch

	ch1 <- ms
}

func main() {

	var res string

	go receive()

	go send()

	res = <-ch1

	fmt.Println(res)
}
