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
}
