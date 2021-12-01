package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1/input1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	var lastDepth = 9999
	var currentDepth = 9999
	var amtIncreases int
	for scanner.Scan() {
		lastDepth = currentDepth
		currentDepth, _ = strconv.Atoi(scanner.Text())
		if currentDepth > lastDepth {
			fmt.Println(currentDepth, ">", lastDepth, "Increase")
			amtIncreases++
		}
	}
	err = file.Close()
	if err != nil {
		return
	}
	fmt.Println(amtIncreases)
}
