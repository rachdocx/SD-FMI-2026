package main

import (
	"fmt"
	"sync"
)

type Node1 struct {
	id         int
	max_id     int
	in_no      int
	ch         chan int
	neighbours []*Node1
}

func send1(node *Node1, _wg *sync.WaitGroup) {

	for _, n := range node.neighbours {
		n.ch <- node.max_id
	}

	_wg.Done()
}

func recv1(node *Node1, _wg *sync.WaitGroup) {
	ma := -1
	for i := 0; i < node.in_no; i++ {
		temp := <-node.ch

		if temp > ma {
			ma = temp
		}
	}

	if ma > node.max_id {
		node.max_id = ma
	}

	_wg.Done()
}

func main() {

	//grafu din curs
	n9 := &Node1{id: 9, max_id: 9}
	n2 := &Node1{id: 2, max_id: 2}
	n6 := &Node1{id: 6, max_id: 6}
	n14 := &Node1{id: 14, max_id: 14}
	n4 := &Node1{id: 4, max_id: 4}
	n3 := &Node1{id: 3, max_id: 3}
	n17 := &Node1{id: 17, max_id: 17}
	n8 := &Node1{id: 8, max_id: 8}
	n22 := &Node1{id: 22, max_id: 22}
	n19 := &Node1{id: 19, max_id: 19}

	n9.in_no = 2
	n2.in_no = 2
	n6.in_no = 1
	n14.in_no = 1
	n4.in_no = 3
	n3.in_no = 1
	n17.in_no = 2
	n8.in_no = 2
	n22.in_no = 2
	n19.in_no = 1

	my_topology := []*Node1{n9, n2, n6, n14, n4, n3, n17, n8, n22, n19}

	n9.ch = make(chan int, n9.in_no)
	n2.ch = make(chan int, n2.in_no)
	n6.ch = make(chan int, n6.in_no)
	n14.ch = make(chan int, n14.in_no)
	n4.ch = make(chan int, n4.in_no)
	n3.ch = make(chan int, n3.in_no)
	n17.ch = make(chan int, n17.in_no)
	n8.ch = make(chan int, n8.in_no)
	n22.ch = make(chan int, n22.in_no)
	n19.ch = make(chan int, n19.in_no)

	n9.neighbours = append(n9.neighbours, n2)
	n2.neighbours = append(n2.neighbours, n6, n4)
	n6.neighbours = append(n6.neighbours, n14)
	n14.neighbours = append(n14.neighbours, n4, n17)
	n4.neighbours = append(n4.neighbours, n3, n22)
	n3.neighbours = append(n3.neighbours, n17, n8)
	n17.neighbours = append(n17.neighbours, n8, n9)
	n8.neighbours = append(n8.neighbours, n22)
	n22.neighbours = append(n22.neighbours, n19, n2)
	n19.neighbours = append(n19.neighbours, n9, n4)

	iteration := 5
	n := len(my_topology)

	var wg sync.WaitGroup

	for iteration != 0 {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go send1(my_topology[i], &wg)
		}

		wg.Wait()

		for i := 0; i < n; i++ {
			wg.Add(1)
			go recv1(my_topology[i], &wg)
		}

		wg.Wait()

		iteration -= 1
	}

	for i := 0; i < n; i++ {
		fmt.Printf("the node %v thinks the leader is %v\n", my_topology[i].id, my_topology[i].max_id)
	}
}
