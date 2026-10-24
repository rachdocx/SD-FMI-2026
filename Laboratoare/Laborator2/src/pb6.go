package main

import (
	"fmt"
	"time"
)

type Robot struct {
	name   string
	dx, dy int
}

func (r *Robot) go_right() {
	r.dx += 1
}

func (r *Robot) go_left() {
	r.dx -= 1
}

func (r *Robot) go_up() {
	r.dy += 1
}

func (r *Robot) go_down() {
	r.dx -= 1
}

func (r *Robot) start(dx_finish int, dy_finish int, ch chan string) {

	for r.dx != dx_finish || r.dy != dy_finish {

		time.Sleep(20 * time.Millisecond) //timp cu care face un pas, eventual cu random

		if r.dx < dx_finish {
			r.go_right()
		} else if r.dx > dx_finish {
			r.go_left()
		}

		if r.dy < dy_finish {
			r.go_up()
		} else if r.dy > dy_finish {
			r.go_down()
		}

	}
	ch <- r.name
}

func main() {
	robot1 := Robot{name: "mugurel", dx: 0, dy: 0}
	robot2 := Robot{name: "trifoiel", dx: 0, dy: 0}

	ch_r1 := make(chan string)
	ch_r2 := make(chan string)

	go robot1.start(6, 7, ch_r1)
	go robot2.start(6, 7, ch_r2)

	select {
	case msg := <-ch_r1:
		fmt.Println(msg + " robot 1 has won the race")
	case msg := <-ch_r2:
		fmt.Println(msg + " robot 2 has won the race")
	}

}
