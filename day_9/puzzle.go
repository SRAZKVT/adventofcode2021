package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SIZEX = 100
const SIZEY = 100

func main() {
	puzzle1()
}

func puzzle1() {
	heightmap := parseInput()
	lowPointsValues := findLowPoints(heightmap)
	fmt.Println(lowPointsValues)
	result := 0
	for _, value := range lowPointsValues {
		result += value + 1
	}
	fmt.Println(result)
}

func parseInput() [SIZEX][SIZEY]int {
	file, err := os.Open("day_9/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	var result [SIZEX][SIZEY]int
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, character := range scanner.Text() {
			value, _ := strconv.Atoi(string(character))
			result[y][x] = value
		}
		y++
	}
	return result
}

func findLowPoints(heightmap [SIZEX][SIZEY]int) []int {
	var result []int
	for y, line := range heightmap {
		for x, value := range line {
			if y != 0 && value >= heightmap[y-1][x] {
				continue
			}
			if y != SIZEX-1 && value >= heightmap[y+1][x] {
				continue
			}
			if x != 0 && value >= heightmap[y][x-1] {
				continue
			}
			if x != SIZEY-1 && value >= heightmap[y][x+1] {
				continue
			}
			result = append(result, value)
		}
	}
	return result
}
