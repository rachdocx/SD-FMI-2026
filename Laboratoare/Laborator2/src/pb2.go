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

	_wg.Done() // <- anuntam ca routine ul e gata
}

func main() {

	wg := sync.WaitGroup{} //<- facem un wait group

	wg.Add(1)

	go recv_hello_world(&wg) // <- pornim goroutine ul

	ch <- "hello world!"

	wg.Wait() // asteptam pana routine ul e gata

}
