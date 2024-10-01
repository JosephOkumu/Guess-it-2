package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"guessit/calcs"
)

func main() {
	if len(os.Args) != 1 {
		return
	}

	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}
		numbers = append(numbers, num)
		if len(numbers) > 1 {
			lower, upper := calcs.PredictRange(numbers)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
