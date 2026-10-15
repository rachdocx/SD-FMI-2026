package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	newMap := map[int][]string{}

	for name, dept := range myMap {
		newMap[dept] = append(newMap[dept], name)
	}

	fmt.Println(newMap)
}
