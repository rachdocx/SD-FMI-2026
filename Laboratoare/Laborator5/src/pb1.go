package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

type Point struct {
	dx, dy float64
}

type Robot struct {
	pos Point
	id  int
	ch  chan Point
}

func send(robot *Robot, _wg *sync.WaitGroup) {
	id := robot.id

	for i := 0; i < n; i++ {
		if my_topology[id][i] == 1 {
			robots[i].ch <- robot.pos
		}
	}

	_wg.Done()
}

func recv(robot *Robot, _wg *sync.WaitGroup) {

	var sum_dx float64
	var sum_dy float64

	sum_dx = 0.0
	sum_dy = 0.0
	no_neigh := 0

	id := robot.id

	for i := 0; i < n; i++ {
		if my_topology[id][i] == 1 {
			point := <-robot.ch
			sum_dx += point.dx
			sum_dy += point.dy
			no_neigh += 1
		}
	}

	new_position_dx := (robot.pos.dx + float64(sum_dx)) / (1 + float64(no_neigh))
	new_position_dy := (robot.pos.dy + float64(sum_dy)) / (1 + float64(no_neigh))

	robot.pos.dx = new_position_dx
	robot.pos.dy = new_position_dy

	_wg.Done()
}

func distance(robot1 *Robot, robot2 *Robot, raza float64) bool {

	d := math.Sqrt((robot1.pos.dx-robot2.pos.dx)*(robot1.pos.dx-robot2.pos.dx) + (robot1.pos.dy-robot2.pos.dy)*(robot1.pos.dy-robot2.pos.dy))
	if d <= raza {
		return true
	} else {
		return false
	}
}

var my_topology [][]int
var robots []Robot

const n = 10
const r = 50

func main() {
	var wg sync.WaitGroup

	robots = make([]Robot, n)
	my_topology = make([][]int, n)

	for i := 0; i < n; i++ {
		my_topology[i] = make([]int, n)
		robots[i] = Robot{pos: Point{dx: rand.Float64() * 100.0, dy: rand.Float64() * 100.0}, id: i, ch: make(chan Point, n)}
	}

	iterations := 50

	for iterations != 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j && distance(&robots[i], &robots[j], r) {
					my_topology[i][j] = 1
				} else {
					my_topology[i][j] = 0
				}
			}
		}

		for i := 0; i < n; i++ {
			wg.Add(1)
			go send(&robots[i], &wg)
		}

		wg.Wait()

		for i := 0; i < n; i++ {
			wg.Add(1)
			go recv(&robots[i], &wg)
		}

		wg.Wait()

		fmt.Printf("\nAt iteration %v the positions are:\n", iterations)
		for i := 0; i < n; i++ {
			fmt.Printf("The robot number %v is at %v dx and %v dy\n", robots[i].id, robots[i].pos.dx, robots[i].pos.dy)
		}

		iterations--
	}
	/*
		//cu raza 15 mi se aduna in niste clustere locale
		At iteration 1 the positions are:
		The robot number 0 is at 16.852494855967077 dx and 16.852494855967077 dy
		The robot number 1 is at 16.852494855967077 dx and 16.852494855967077 dy
		The robot number 2 is at 16.852494855967077 dx and 16.852494855967077 dy
		The robot number 3 is at 45.00000000000001 dx and 45.00000000000001 dy
		The robot number 4 is at 45.00000000000001 dx and 45.00000000000001 dy
		The robot number 5 is at 45.00000000000001 dx and 45.00000000000001 dy
		The robot number 6 is at 45.00000000000001 dx and 45.00000000000001 dy
		The robot number 7 is at 73.14750514403293 dx and 73.14750514403293 dy
		The robot number 8 is at 73.14750514403293 dx and 73.14750514403293 dy
		The robot number 9 is at 73.14750514403293 dx and 73.14750514403293 dy
	*/
}
