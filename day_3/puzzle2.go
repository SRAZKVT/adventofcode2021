package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_3/input2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	var toSearch []string
	pos := 0
	for scanner.Scan() {
		line := scanner.Text()
		toSearch = append(toSearch, line)
		pos++
	}
	err = file.Close()
	if err != nil {
		return
	}
	fmt.Println(lifeSupportRating(toSearch))
}

func lifeSupportRating(toSearch []string) int {
	wordOxygen := getWordOxygen(toSearch)
	wordCo2 := getWordCo2(toSearch)

	oxygen := getDecimalFromBinary(wordOxygen)
	co2 := getDecimalFromBinary(wordCo2)

	fmt.Println("Oxygen:", oxygen, "Co2:", co2)
	return co2 * oxygen
}

func getDecimalFromBinary(word string) int {
	result, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return 0
	}
	return int(result)
}

func getWordOxygen(toSearch []string) string {
	fmt.Println("Oxygen :")
	pos := 0
	for len(toSearch) != 1 {
		toSearch = removeWrongWordsOxygen(toSearch, pos)
		pos++
	}
	fmt.Println(toSearch[0])
	return toSearch[0]
}

func removeWrongWordsOxygen(toSearch []string, pos int) []string {
	fmt.Println(toSearch)
	var result []string
	var characterAtPosMostCommon uint8
	amtLow, amtHigh := getAmtCharsAtPos(toSearch, pos)
	if amtHigh >= amtLow {
		characterAtPosMostCommon = '1'
	} else {
		characterAtPosMostCommon = '0'
	}
	for _, word := range toSearch {
		if word[pos] == characterAtPosMostCommon {
			result = append(result, word)
		}
	}
	return result
}

func getAmtCharsAtPos(toSearch []string, pos int) (int, int) {
	amtLow := 0
	amtHigh := 0
	for _, line := range toSearch {
		if line[pos] == '0' {
			amtLow++
		} else {
			amtHigh++
		}
	}
	return amtLow, amtHigh
}

func getWordCo2(toSearch []string) string {
	fmt.Println("Co2")
	pos := 0
	for len(toSearch) != 1 {
		toSearch = removeWrongWordsCo2(toSearch, pos)
		pos++
	}
	fmt.Println(toSearch[0])
	return toSearch[0]
}

func removeWrongWordsCo2(toSearch []string, pos int) []string {
	fmt.Println(toSearch)
	var result []string
	var characterAtPosMostCommon uint8
	amtLow, amtHigh := getAmtCharsAtPos(toSearch, pos)
	if amtHigh < amtLow {
		characterAtPosMostCommon = '1'
	} else {
		characterAtPosMostCommon = '0'
	}
	for _, word := range toSearch {
		if word[pos] == characterAtPosMostCommon {
			result = append(result, word)
		}
	}
	return result
}
