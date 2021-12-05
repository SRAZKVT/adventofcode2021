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
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)

	var mapOfVents [SIZE][SIZE]int

	for scanner.Scan() {
		positions := strings.Split(scanner.Text(), " -> ")
		mapOfVents = markPositionsHorizontalVertical(mapOfVents, positions)
		mapOfVents = markPositionsDiagonal(mapOfVents, positions)
	}
	count := 0
	for _, line := range mapOfVents {
		for _, spot := range line {
			if spot >= 2 {
				count++
			}
		}
	}
	fmt.Println("puzzle 2:", count)
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

func markPositionsDiagonal(mapOfVents [SIZE][SIZE]int, positions []string) [SIZE][SIZE]int {
	pos1 := getPosition(positions[0])
	pos2 := getPosition(positions[1])

	// straight line already done
	if pos1[0] == pos2[0] || pos1[1] == pos2[1] {
		return mapOfVents
	}

	// wait i'm dumb i can just iterate over x and just decrease or increse y, more complex is unnecessary
	shouldBeReverse := (pos1[0] > pos2[0] || pos1[1] > pos2[1]) && !(pos1[0] > pos2[0] && pos1[1] > pos2[1]) // bruh why no XOR

	var minPos, maxPos, startPos [2]int
	minPos[0], maxPos[0] = min(pos1[0], pos2[0]), max(pos1[0], pos2[0])
	minPos[1], maxPos[1] = min(pos1[1], pos2[1]), max(pos1[1], pos2[1])
	startPos[0], startPos[1] = minPos[0], minPos[1]
	if shouldBeReverse {
		startPos[1] = maxPos[1]
	}
	for i := 0; i < maxPos[1]-(minPos[1]-1); i++ {
		if shouldBeReverse {
			mapOfVents[startPos[1]-i][startPos[0]+i]++
		} else {
			mapOfVents[startPos[1]+i][startPos[0]+i]++
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

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}
