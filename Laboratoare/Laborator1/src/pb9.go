package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// TWO SUM LETSGOOOlang

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	data := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(data[0])
	s, _ := strconv.Atoi(data[1])

	scanner.Scan()

	data = strings.Split(scanner.Text(), " ")

	myMap := make(map[int]int)

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(data[i])

		j, ok := myMap[s-num]

		if ok {
			fmt.Println(j, i)
		}

		myMap[num] = i
	}
}
