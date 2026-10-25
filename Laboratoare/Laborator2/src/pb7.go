package main

import (
	"fmt"
	"sync"
	"time"
)

type Point struct {
	dx, dy   int
	finished bool
}

type Robot1 struct {
	name     string
	location Point
}

func (r *Robot1) go_right() {
	r.location.dx += 1
}

func (r *Robot1) go_left() {
	r.location.dx -= 1
}

func (r *Robot1) go_up() {
	r.location.dy += 1
}

func (r *Robot1) go_down() {
	r.location.dy -= 1
}

func (r *Robot1) start(dx_finish int, dy_finish int, in chan Point, out chan Point, ms chan bool, _wg *sync.WaitGroup) {

	me := false
	other := false

	//cat timp nu au ajuns ambii roboti
	for !me || !other {

		if !me {

			time.Sleep(1000 * time.Millisecond) //timp cu care face un pas, eventual cu random

			if r.location.dx < dx_finish {
				r.go_right()
			} else if r.location.dx > dx_finish {
				r.go_left()
			}

			if r.location.dy < dy_finish {
				r.go_up()
			} else if r.location.dy > dy_finish {
				r.go_down()
			}

			if r.location.dx == dx_finish && r.location.dy == dy_finish {
				me = true
				//anuntam observer-ul ca am terminat
				ms <- true
				fmt.Printf("Robot %v arrived at destination\n", r.name)
			}
		}

		//trimitem locatia robotului curent
		r.location.finished = me
		out <- r.location

		//primim locatia celuilalt robot
		temp := <-in
		other = temp.finished

		if !me {
			fmt.Printf("Robot %v is moving and got notified the other is at %v, %v\n", r.name, temp.dx, temp.dy)
		} else if !other {
			fmt.Printf("Robot %v is waiting and got notified the other is at %v, %v\n", r.name, temp.dx, temp.dy)
		}
	}

	_wg.Done()
}

func main() {
	start1 := Point{dx: 0, dy: 0, finished: false}
	start2 := Point{dx: 0, dy: 0, finished: false}

	robot1 := Robot1{name: "mugurel", location: start1}
	robot2 := Robot1{name: "trifoiel", location: start2}

	//canalele de comunicare intre roboti
	var ch1 = make(chan Point, 1)
	var ch2 = make(chan Point, 1)

	//canalele de comunicare intre roboti si observer
	var ms1 = make(chan bool)
	var ms2 = make(chan bool)

	wg1 := sync.WaitGroup{}

	wg1.Add(2)

	go robot1.start(6, 7, ch1, ch2, ms1, &wg1)
	go robot2.start(6, 12, ch2, ch1, ms2, &wg1)

	var state_r1, state_r2 bool

	state_r1 = false
	state_r2 = false

	for !state_r1 || !state_r2 {
		select {
		case state_r1 = <-ms1:
			fmt.Printf("The observer has been notified that mugurel finished the race\n")
		case state_r2 = <-ms2:
			fmt.Printf("The observer has been notified that trifoiel finished the race\n")
		}
	}

	wg1.Wait()
}
