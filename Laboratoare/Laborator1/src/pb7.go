package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("cuvinte7.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	fmt.Println(scanner.Text())

	words := []string{}

	words = strings.Split(scanner.Text(), " ")

	sol := make(map[string]uint8)

	for _, word := range words {

		_, ok := sol[word]

		if !ok {
			sol[word] = 1
		} else {
			sol[word] += 1
		}

	}

	for word, count := range sol {
		fmt.Printf("The word %v has %v appearances\n", word, count)
	}

}
