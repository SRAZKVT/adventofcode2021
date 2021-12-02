package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_2/input1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	var x, depth, value, aim int
	var line string
	var args []string

	for scanner.Scan() {
		line = scanner.Text()
		args = strings.Split(line, " ")
		value, _ = strconv.Atoi(args[1])
		switch args[0] {
		case "forward":
			x += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		}
		fmt.Println(line, "new pos is", x, depth)
	}
	fmt.Println(x * depth)
}
