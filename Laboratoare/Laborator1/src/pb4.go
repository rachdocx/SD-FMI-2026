package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	intSlice1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

	index, _ := strconv.Atoi(data)

	intSlice1 = append(intSlice1[0:index], intSlice1[index+1:]...)

	fmt.Println(intSlice1)
}
