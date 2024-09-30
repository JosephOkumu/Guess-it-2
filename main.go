package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error parsing input")
			continue
		}
		numbers = append(numbers, num)
	}
	fmt.Println(numbers)
}
