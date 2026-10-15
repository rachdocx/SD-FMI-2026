package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Andrei ": 190,
		"Maria ":  165,
		"Ion":     182,
		"Elena ":  172,
		"George ": 178,
		"Ana":     165,
	}

	// cum sa sortez un map ca in c++?
	for i := range len(myMap) {

		var mi int = 500
		var maxName string = ""

		for name, height := range myMap {
			if height < mi {
				mi = height
				maxName = name
			}
		}

		delete(myMap, maxName)

		fmt.Printf("%v: Name: %v, Age: %v\n", i, maxName, mi)
	}
}
