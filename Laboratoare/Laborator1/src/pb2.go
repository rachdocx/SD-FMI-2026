package main

import "fmt"

func main() {

	// 2
	const SIZE = 5
	var arr [SIZE]int32 = [SIZE]int32{5, 4, 3, 2, 1}
	var aux int32

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			if arr[i] > arr[j] {
				aux = arr[i]
				arr[i] = arr[j]
				arr[j] = aux
			}
		}
	}

	fmt.Println(arr)

}
