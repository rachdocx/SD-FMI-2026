package main

import "fmt"

func main() {

	// 5

	sliceString := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	var aux1 string
	var length = len(sliceString)

	for i := 0; i < length/2; i++ {
		aux1 = sliceString[i]
		sliceString[i] = sliceString[length-i-1]
		sliceString[length-i-1] = aux1
	}

	fmt.Println(sliceString)
}
