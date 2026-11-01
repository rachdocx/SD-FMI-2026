package main

import (
	"fmt"
	"sync"
)

type Node struct {
	id     int
	leader int
	ch     chan int
	next   *Node
}

func send(node *Node, _wg *sync.WaitGroup) {

	//trimitem fie id-ul curent fie id-ul leader ului stiut de un nod oarecare
	if node.id > node.leader {
		node.next.ch <- node.id
	} else {
		node.next.ch <- node.leader
	}

	_wg.Done()
}

func recv(node *Node, _wg *sync.WaitGroup, _main_signal chan int) {
	temp := <-node.ch

	//daca primim un id mai mare il consideram lider

	//daca id-ul este egal cu id-ul nostru inseamna ca algoritmul s-a terminat si
	//trimitem un mesaj catre main
	if temp > node.id {
		node.leader = temp
	} else if temp == node.id {
		node.leader = temp
		_main_signal <- temp
	}
	_wg.Done()
}

func main() {

	//facem topologia de tip inel fara sa stim cate noduri sunt
	head := &Node{id: 1, ch: make(chan int, 1), next: nil}
	tail := head

	for i := 2; i <= 10; i++ {

		var temp Node

		if i == 2 {
			temp = Node{id: 67, leader: -1, ch: make(chan int, 1), next: nil}
		} else {
			temp = Node{id: i, leader: -1, ch: make(chan int, 1), next: nil}
		}

		head.next = &temp

		head = &temp
	}

	head.next = tail

	var wg sync.WaitGroup
	main_signal := make(chan int, 1)
	ok := true

	//rulam pana primim un semnal de la un thread
	for ok {
		curr := tail

		//toate thread urile trimit
		for {
			wg.Add(1)
			go send(curr, &wg)
			curr = curr.next
			if curr == tail {
				break
			}
		}
		//thread-urile se asteapta pana sa faca recieve
		wg.Wait()
		curr = tail

		//toate thread urile primesc
		for {
			wg.Add(1)
			go recv(curr, &wg, main_signal)
			curr = curr.next
			if curr == tail {
				break
			}
		}
		//toate thread-urile se asteapta
		wg.Wait()

		//daca am primit mesaj ca leader-ul a fost gasit, tratam acest caz
		select {
		case ms := <-main_signal:
			{
				fmt.Printf("the leader was elected and is %v\n", ms)
				ok = false
			}
		default:
		}
	}

	curr := tail

	for curr != head {
		fmt.Printf("the node: %v  knows the leader is %v\n", curr.id, curr.leader)

		curr = curr.next
	}
	/*
		the leader was elected and is 67
		the node: 1  knows the leader is 67
		the node: 67  knows the leader is -1
		the node: 3  knows the leader is 67
		the node: 4  knows the leader is 67
		the node: 5  knows the leader is 67
		the node: 6  knows the leader is 67
		the node: 7  knows the leader is 67
		the node: 8  knows the leader is 67
		the node: 9  knows the leader is 67
	*/
}
