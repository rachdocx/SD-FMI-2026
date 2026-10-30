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

	ch1 <- ms // 3 revieve primeste mesajul, il trimite catre main si se termina
}

func main() {

	var res string

	go receive() // 1 recieve asteapta un mesaj

	go send() // 2 send trimite un mesaj

	res = <-ch1 // 4 main primeste mesajul

	fmt.Println(res)
}
