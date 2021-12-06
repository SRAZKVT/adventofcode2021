package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	amtDays := 80
	amtFishies := "3,3,2,1,4,1,1,2,3,1,1,2,1,2,1,1,1,1,1,1,4,1,1,5,2,1,1,2,1,1,1,3,5,1,5,5,1,1,1,1,3,1,1,3,2,1,1,1,1,1,1,4,1,1,1,1,1,1,1,4,1,3,3,1,1,3,1,3,1,2,1,3,1,1,4,1,2,4,4,5,1,1,1,1,1,1,4,1,5,1,1,5,1,1,3,3,1,3,2,5,2,4,1,4,1,2,4,5,1,1,5,1,1,1,4,1,1,5,2,1,1,5,1,1,1,5,1,1,1,1,1,3,1,5,3,2,1,1,2,2,1,2,1,1,5,1,1,4,5,1,4,3,1,1,1,1,1,1,5,1,1,1,5,2,1,1,1,5,1,1,1,4,4,2,1,1,1,1,1,1,1,3,1,1,4,4,1,4,1,1,5,3,1,1,1,5,2,2,4,2,1,1,3,1,5,5,1,1,1,4,1,5,1,1,1,4,3,3,3,1,3,1,5,1,4,2,1,1,5,1,1,1,5,5,1,1,2,1,1,1,3,1,1,1,2,3,1,2,2,3,1,3,1,1,4,1,1,2,1,1,1,1,3,5,1,1,2,1,1,1,4,1,1,1,1,1,2,4,1,1,5,3,1,1,1,2,2,2,1,5,1,3,5,3,1,1,4,1,1,4"
	lanternFishies := getAmtFishies(amtFishies)
	lanternFishies = populateFishiesForDays(lanternFishies, amtDays)
	fmt.Println("Puzzle 1:", getAmtFishiesInPopulation(lanternFishies))
}

func puzzle2() {
	amtDays := 256
	amtFishies := "3,3,2,1,4,1,1,2,3,1,1,2,1,2,1,1,1,1,1,1,4,1,1,5,2,1,1,2,1,1,1,3,5,1,5,5,1,1,1,1,3,1,1,3,2,1,1,1,1,1,1,4,1,1,1,1,1,1,1,4,1,3,3,1,1,3,1,3,1,2,1,3,1,1,4,1,2,4,4,5,1,1,1,1,1,1,4,1,5,1,1,5,1,1,3,3,1,3,2,5,2,4,1,4,1,2,4,5,1,1,5,1,1,1,4,1,1,5,2,1,1,5,1,1,1,5,1,1,1,1,1,3,1,5,3,2,1,1,2,2,1,2,1,1,5,1,1,4,5,1,4,3,1,1,1,1,1,1,5,1,1,1,5,2,1,1,1,5,1,1,1,4,4,2,1,1,1,1,1,1,1,3,1,1,4,4,1,4,1,1,5,3,1,1,1,5,2,2,4,2,1,1,3,1,5,5,1,1,1,4,1,5,1,1,1,4,3,3,3,1,3,1,5,1,4,2,1,1,5,1,1,1,5,5,1,1,2,1,1,1,3,1,1,1,2,3,1,2,2,3,1,3,1,1,4,1,1,2,1,1,1,1,3,5,1,1,2,1,1,1,4,1,1,1,1,1,2,4,1,1,5,3,1,1,1,2,2,2,1,5,1,3,5,3,1,1,4,1,1,4"
	lanternFishies := getAmtFishies(amtFishies)
	lanternFishies = populateFishiesForDays(lanternFishies, amtDays)
	fmt.Println("Puzzle 2:", getAmtFishiesInPopulation(lanternFishies))
}

func getAmtFishies(amtFishies string) [9]uint64 {
	var result [9]uint64
	for _, number := range strings.Split(amtFishies, ",") {
		value, _ := strconv.Atoi(number)
		result[value]++
	}
	return result
}

func populateFishiesForDays(fishiesPopulation [9]uint64, days int) [9]uint64 {
	days++
	for i := 1; i < days; i++ {
		temp := fishiesPopulation[0]
		for i := 0; i < 8; i++ {
			fishiesPopulation[i] = fishiesPopulation[i+1]
		}
		fishiesPopulation[6] += temp
		fishiesPopulation[8] = temp
	}
	return fishiesPopulation
}

func getAmtFishiesInPopulation(population [9]uint64) uint64 {
	var result uint64
	for i := 0; i < 9; i++ {
		result += population[i]
	}
	return result
}