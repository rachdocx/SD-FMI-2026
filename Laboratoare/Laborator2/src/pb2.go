package main

import (
	"fmt"
	"sync"
)

var ch = make(chan string)

func recv_hello_world(_wg *sync.WaitGroup) {
	var ms string

	ms = <-ch

	fmt.Println(ms)

	_wg.Done()
}

func main() {

	wg := sync.WaitGroup{}

	go recv_hello_world(&wg)

	wg.Add(1)

	ch <- "hello world!"

	wg.Wait()

}
