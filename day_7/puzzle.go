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
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("day_7/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	var crabbiesPos []int
	maxCrabbiePos := 0
	for _, value := range strings.Split(input, ",") {
		crabbiePos, _ := strconv.Atoi(value)
		crabbiesPos = append(crabbiesPos, crabbiePos)
		if crabbiePos > maxCrabbiePos {
			maxCrabbiePos = crabbiePos
		}
	}
	amtFuelMin := 9999999999999
	for i := 0; i < maxCrabbiePos; i++ {
		if amtFuel := getAmtFuelMovinCrabbies(crabbiesPos, i); amtFuel < amtFuelMin {
			amtFuelMin = amtFuel
		}
	}
	fmt.Println(amtFuelMin)
}

func getAmtFuelMovinCrabbies(crabbiesPos []int, i int) int {
	result := 0
	for _, crabbiePos := range crabbiesPos {
		if crabbiePos == i {
			continue
		}
		result += abs(crabbiePos - i)
	}
	return result
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func puzzle2() {
	file, err := os.Open("day_7/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	var crabbiesPos []int
	maxCrabbiePos := 0
	for _, value := range strings.Split(input, ",") {
		crabbiePos, _ := strconv.Atoi(value)
		crabbiesPos = append(crabbiesPos, crabbiePos)
		if crabbiePos > maxCrabbiePos {
			maxCrabbiePos = crabbiePos
		}
	}
	amtFuelMin := 9999999999999
	for i := 0; i < maxCrabbiePos; i++ {
		if amtFuel := getAmtFuelMovinCrabbies2(crabbiesPos, i); amtFuel < amtFuelMin {
			amtFuelMin = amtFuel
		}
	}
	fmt.Println(amtFuelMin)
}

func getAmtFuelMovinCrabbies2(crabbiesPos []int, i int) int {
	result := 0
	for _, crabbiePos := range crabbiesPos {
		if crabbiePos == i {
			continue
		}
		x := abs(crabbiePos - i)
		result += (x * (x + 1)) / 2
	}
	return result
}
