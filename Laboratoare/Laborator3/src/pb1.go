package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {

	delta = (t4.Sub(t1)) - (t3.Sub(t2))
	theta = ((t2.Sub(t1)) + (t3.Sub(t4))) / 2
	return theta, delta
}
func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {

	//request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	//la ce timp a ajuns request ul
	t2 := s_clock.Now()

	//procesare pe server
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	//response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	var ceas_client, ceas_server Clock

	ceas_server = Clock{offset: 0}
	ceas_client = Clock{offset: 1000 * time.Millisecond}

	var theta, delta time.Duration

	theta, delta = measure(&ceas_client, &ceas_server)

	fmt.Printf("a)Theta: %v | Delta: %v\n", theta, delta)
	ceas_client.offset += theta

	fmt.Printf("Offset ul pentru client a devenit: %v\n", ceas_client.offset)

	var ceas_client_ex, ceas_server_ex Clock

	ceas_server_ex = Clock{offset: 0}
	ceas_client_ex = Clock{offset: 1000 * time.Millisecond}

	var theta_ex, delta_ex time.Duration

	theta_ex = 10000000 * time.Millisecond
	delta_ex = 10000000 * time.Millisecond

	for i := 0; i < 10; i++ {
		var theta1, delta1 time.Duration

		theta1, delta1 = measure(&ceas_client_ex, &ceas_server_ex)

		//este corect? ar trebui sa pastrez aceleasi din accelasi
		//experiment si pe theta si pe delta nu seaparat

		if delta1 < delta_ex {
			delta_ex = delta1
		}

		if theta1 < theta_ex {
			theta_ex = theta1
		}

	}

	fmt.Printf("\nb) Theta: %v | Delta: %v\n", theta, delta)
	fmt.Printf("Theta min dupa experiment: %v | Delta minim dupa experiment: %v\n", theta_ex, delta_ex)
}
