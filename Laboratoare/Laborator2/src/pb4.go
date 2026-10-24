package main

import (
	"fmt"
	"time"
)

var notify = make(chan bool)

func notify_main() {

	time.Sleep(100 * time.Millisecond)

	notify <- true
}

func main() {

	var ok bool

	go notify_main()

	ok = <-notify

	for ; ok != true; ok = <-notify {
		fmt.Println(ok)
	}

	fmt.Println("The go routine finally ended")

}
