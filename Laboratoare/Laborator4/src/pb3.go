package main

import (
	"fmt"
	"sync"
)

type Node2 struct {
	id         float64
	avg        float64
	in_no      int //cum fac daca nu stiu cate noduri intra intr un nod fara deadlock, ar trebui sa foloesc o coada de mesaje nu un canal?
	ch         chan float64
	neighbours []*Node2
}

func send2(node *Node2, _wg *sync.WaitGroup) {
	for _, n := range node.neighbours {
		n.ch <- node.avg
	}
	_wg.Done()
}

func recv2(node *Node2, _wg *sync.WaitGroup) {
	var temp_avg float64
	var temp_number float64

	for i := 0; i < node.in_no; i++ {
		temp_number = <-node.ch
		temp_avg += temp_number
	}

	temp_avg += node.avg

	temp_avg = temp_avg / (float64(node.in_no) + 1)

	node.avg = temp_avg

	_wg.Done()
}

func main() {

	//grafu din curs
	n9 := &Node2{id: 9, avg: 9}
	n2 := &Node2{id: 2, avg: 2}
	n6 := &Node2{id: 6, avg: 6}
	n14 := &Node2{id: 14, avg: 14}
	n4 := &Node2{id: 4, avg: 4}
	n3 := &Node2{id: 3, avg: 3}
	n17 := &Node2{id: 17, avg: 17}
	n8 := &Node2{id: 8, avg: 8}
	n22 := &Node2{id: 22, avg: 22}
	n19 := &Node2{id: 19, avg: 19}

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

	my_topology := []*Node2{n9, n2, n6, n14, n4, n3, n17, n8, n22, n19}

	n9.ch = make(chan float64, n9.in_no)
	n2.ch = make(chan float64, n2.in_no)
	n6.ch = make(chan float64, n6.in_no)
	n14.ch = make(chan float64, n14.in_no)
	n4.ch = make(chan float64, n4.in_no)
	n3.ch = make(chan float64, n3.in_no)
	n17.ch = make(chan float64, n17.in_no)
	n8.ch = make(chan float64, n8.in_no)
	n22.ch = make(chan float64, n22.in_no)
	n19.ch = make(chan float64, n19.in_no)

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

	iteration := 500
	n := len(my_topology)

	var wg sync.WaitGroup

	for iteration != 0 {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go send2(my_topology[i], &wg)
		}

		wg.Wait()

		for i := 0; i < n; i++ {
			wg.Add(1)
			go recv2(my_topology[i], &wg)
		}

		wg.Wait()

		iteration -= 1
	}

	for i := 0; i < n; i++ {
		fmt.Printf("the node %v thinks the average %v\n", my_topology[i].id, my_topology[i].avg)
	}
	/*
		// MEDIA E 10.4
		the node 9 thinks the average 10.498762376237629
		the node 2 thinks the average 10.498762376237629
		the node 6 thinks the average 10.498762376237629
		the node 14 thinks the average 10.498762376237629
		the node 4 thinks the average 10.498762376237629
		the node 3 thinks the average 10.498762376237629 | <- toate tind spre medie
		the node 17 thinks the average 10.498762376237629
		the node 8 thinks the average 10.498762376237629
		the node 22 thinks the average 10.498762376237629
		the node 19 thinks the average 10.498762376237629
	*/
}
