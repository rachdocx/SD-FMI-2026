package main

import "fmt"

var ch = make(chan string, 2)

func two_messages() {

	ch <- "hello"
	ch <- "world"

}
func main() {
	var ms1, ms2 string

	go two_messages()
	ms1 = <-ch
	ms2 = <-ch

	fmt.Println(ms1)
	fmt.Println(ms2)

}
