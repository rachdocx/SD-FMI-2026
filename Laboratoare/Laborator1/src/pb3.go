package main

import "fmt"

func main() {

	// 3
	const SIZE1 = 10
	var arr1 [SIZE1]int32 = [SIZE1]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	intSlice := []int32{}

	for i := 0; i < SIZE1; i++ {
		if arr1[i]%2 == 0 {
			intSlice = append(intSlice, arr1[i])
		}
	}

	fmt.Println(intSlice)

}
