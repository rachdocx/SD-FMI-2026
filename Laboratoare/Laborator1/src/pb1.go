package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "os"
	// "strconv"
)

func main() {

	// 1.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var data string = scanner.Text()

	nr, _ := strconv.Atoi(data)

	if nr < 2 {
		fmt.Println("Nu este prim")
	} else {
		for i := 2; i*i <= nr; i++ {
			if nr%i == 0 {
				fmt.Println("Nu este prim")
				return
			}
		}
	}

	fmt.Println("Este prim")

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

	// 4

	intSlice1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}

	scanner.Scan()
	data = scanner.Text()

	index, _ := strconv.Atoi(data)

	intSlice1 = append(intSlice1[0:index], intSlice1[index+1:]...)

	fmt.Println(intSlice1)

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

	// 6

	myMap := map[string]int{
		" Andrei ": 190,
		" Maria ":  165,
		"Ion":      182,
		" Elena ":  172,
		" George ": 178,
		"Ana":      165,
	}

	fmt.Println(myMap)

}
