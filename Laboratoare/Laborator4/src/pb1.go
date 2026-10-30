package main

import (
	"fmt"
	"sync"
)

type Node struct {
	id   int
	ch   chan int
	next *Node
}

func send(node *Node, _wg *sync.WaitGroup) {
	node.next.ch <- node.id
	_wg.Done()
}

func recv(node *Node, _wg *sync.WaitGroup) {
	temp := <-node.ch

	if temp > node.id {
		node.id = temp
	}
	_wg.Done()
}

func main() {
	head := &Node{id: 1, ch: make(chan int, 1), next: nil}
	tail := head

	for i := 2; i <= 10; i++ {

		var temp Node

		if i == 2 {
			temp = Node{id: 67, ch: make(chan int, 1), next: nil}
		} else {
			temp = Node{id: i, ch: make(chan int, 1), next: nil}
		}

		head.next = &temp

		head = &temp
	}

	head.next = tail

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		curr := tail

		for {
			wg.Add(1)
			go send(curr, &wg)
			curr = curr.next
			if curr == tail {
				break
			}
		}

		wg.Wait()
		curr = tail

		for {
			wg.Add(1)
			go recv(curr, &wg)
			curr = curr.next
			if curr == tail {
				break
			}
		}
		wg.Wait()
	}

	curr := tail

	for curr != head {
		fmt.Println(curr.id)

		curr = curr.next
	}

}
