package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1/input2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	var lastDepth = 99999
	var currentDepth = 99999

	var currentDepth1, currentDepth2, currentDepth3 int

	var amtIncreases int
	for scanner.Scan() {
		currentDepth3 = currentDepth2
		currentDepth2 = currentDepth1

		currentDepth1, _ = strconv.Atoi(scanner.Text())
		currentDepth = currentDepth1 + currentDepth2 + currentDepth3
		if currentDepth3 == 0 {
			continue
		}
		if currentDepth > lastDepth {
			amtIncreases++
		}
		lastDepth = currentDepth
	}
	err = file.Close()
	if err != nil {
		return
	}
	fmt.Println(amtIncreases)
}
