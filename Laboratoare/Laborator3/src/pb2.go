package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}
func (l *LamportClock) Send() {
	l.counter++
}
func (l *LamportClock) Receive(ts int) {
	l.counter = max(l.counter, ts)
	l.counter++
}

type Message struct {
	sender, ts int
}

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

func proces(events *[]Event, process_id int, ch_in chan Message, ch_out chan Message, _wg *sync.WaitGroup, _mutex *sync.Mutex) {
	clock := LamportClock{counter: 0}

	n := rand.Intn(10)

	for i := 0; i < n; i++ {
		clock.Tick()

		_mutex.Lock()
		*events = append(*events, Event{process: process_id, internal_counter: clock.counter, kind: "internal"})
		_mutex.Unlock()
	}

	if process_id == 1 {
		clock.Send()

		_mutex.Lock()
		*events = append(*events, Event{process: process_id, internal_counter: clock.counter, kind: "send"})
		_mutex.Unlock()

		ch_out <- Message{sender: process_id, ts: clock.counter}
	}

	ms := <-ch_in

	clock.Receive(ms.ts)
	*events = append(*events, Event{process: process_id, internal_counter: clock.counter, kind: "recv"})

	if process_id != 1 {
		clock.Send()

		_mutex.Lock()
		*events = append(*events, Event{process: process_id, internal_counter: clock.counter, kind: "send"})
		_mutex.Unlock()

		ch_out <- Message{sender: process_id, ts: clock.counter}
	}

	_wg.Done()
}
func less(a, b Event) bool {
	if a.internal_counter < b.internal_counter {
		return true
	}
	return false
}

func main() {

	ch1_2 := make(chan Message, 1)
	ch2_3 := make(chan Message, 1)
	ch3_1 := make(chan Message, 1)

	var events []Event
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(3)

	go proces(&events, 1, ch3_1, ch1_2, &wg, &mutex)
	go proces(&events, 2, ch1_2, ch2_3, &wg, &mutex)
	go proces(&events, 3, ch2_3, ch3_1, &wg, &mutex)
	/*
		Proces: 1 | Tick uri: 1 | Kind: internal
		Proces: 3 | Tick uri: 1 | Kind: internal
		Proces: 2 | Tick uri: 1 | Kind: internal
		Proces: 1 | Tick uri: 2 | Kind: internal
		Proces: 3 | Tick uri: 2 | Kind: internal
		Proces: 2 | Tick uri: 2 | Kind: internal
		Proces: 1 | Tick uri: 3 | Kind: internal
		Proces: 2 | Tick uri: 3 | Kind: internal
		Proces: 1 | Tick uri: 4 | Kind: internal
		Proces: 2 | Tick uri: 4 | Kind: internal
		Proces: 2 | Tick uri: 5 | Kind: internal
		Proces: 1 | Tick uri: 5 | Kind: internal
		Proces: 1 | Tick uri: 6 | Kind: send <- la tick ul 6 a trimis P1
		Proces: 2 | Tick uri: 6 | Kind: internal
		Proces: 2 | Tick uri: 7 | Kind: internal
		Proces: 2 | Tick uri: 8 | Kind: internal
		Proces: 2 | Tick uri: 9 | Kind: internal
		Proces: 2 | Tick uri: 10 | Kind: recv <- la tick ul 9 l a primit p2 si a facut max(counter_intern, timp la care a fost primit) + 1 => tick ul 10
		Proces: 2 | Tick uri: 11 | Kind: send
		Proces: 3 | Tick uri: 12 | Kind: recv    | <- exact acelasi proces
		Proces: 3 | Tick uri: 13 | Kind: send
		Proces: 1 | Tick uri: 14 | Kind: recv
	*/
	wg.Wait()

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	for _, e := range events {
		fmt.Printf("Thread: %v | Tick uri: %v | Kind: %v\n", e.process, e.internal_counter, e.kind)
	}
}
