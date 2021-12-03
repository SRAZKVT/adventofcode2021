package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day_3/input1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)

	epsilon := -1
	gamma := 0
	countLines := 0
	var amtBitsToHigh [12]int
	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		for i, character := range line {
			if character == '1' {
				amtBitsToHigh[i]++
			}
		}
		countLines++
	}
	err = file.Close()
	if err != nil {
		return
	}

	for i, amount := range amtBitsToHigh {
		if amount > countLines/2 {
			gamma += pow(2, len(amtBitsToHigh)-i)
		}
	}
	fmt.Println("gamma is", gamma)
	for i, amount := range amtBitsToHigh {
		if amount < countLines/2 {
			epsilon += pow(2, len(amtBitsToHigh)-i)
		}
	}
	fmt.Println("epsilon is", epsilon)
	fmt.Println(gamma * epsilon)
}
func pow(n int, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 0; i < m-2; i++ {
		result *= n
	}
	return result
}
