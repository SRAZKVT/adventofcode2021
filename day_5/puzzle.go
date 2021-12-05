package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const SIZE = 1000

func main() {
	puzzle1()
	puzzle2()
}

func puzzle2() {
	fmt.Println("puzzle 2:", "not yet implemented")
}

func puzzle1() {
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)

	var mapOfVents [SIZE][SIZE]int

	for scanner.Scan() {
		positions := strings.Split(scanner.Text(), " -> ")
		mapOfVents = markPositionsHorizontalVertical(mapOfVents, positions)
	}
	count := 0
	for _, line := range mapOfVents {
		for _, spot := range line {
			if spot >= 2 {
				count++
			}
		}
	}
	fmt.Println("puzzle 1:", count)
}

func markPositionsHorizontalVertical(mapOfVents [SIZE][SIZE]int, positions []string) [SIZE][SIZE]int {
	pos1 := getPosition(positions[0])
	pos2 := getPosition(positions[1])
	if pos1[0] == pos2[0] {
		for i := 0; i < SIZE; i++ {
			if isWithinBounds(i, pos1[1], pos2[1]) {
				mapOfVents[i][pos1[0]]++
			}
		}
	}
	if pos1[1] == pos2[1] {
		for i := 0; i < SIZE; i++ {
			if isWithinBounds(i, pos1[0], pos2[0]) {
				mapOfVents[pos1[1]][i]++
			}
		}
	}
	return mapOfVents
}

func getPosition(s string) []int {
	var result []int
	values := strings.Split(s, ",")
	for _, val := range values {
		pos, _ := strconv.Atoi(val)
		result = append(result, pos)
	}
	return result
}

func isWithinBounds(i, pos1, pos2 int) bool {
	return (i >= pos2 && i <= pos1) || (i <= pos2 && i >= pos1)
}
