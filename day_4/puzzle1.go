package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbersToCome []int
var pos int

type bingoBoard struct {
	numbers       [5][5]int
	markedNumbers [5][5]bool
}

var allBoards []bingoBoard

func main() {
	file, err := os.Open("day_4/input1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	lineNumbers := scanner.Text()
	args := strings.Split(lineNumbers, ",")
	for _, number := range args {
		value, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err.Error())
		}
		numbersToCome = append(numbersToCome, value)
	}
	scanner.Scan()
	for scanner.Scan() {
		var currentTable bingoBoard
		for i := 0; i < 5; i++ {
			args = strings.Split(scanner.Text(), " ")
			index := 0
			for _, number := range args {
				if number == "" {
					continue
				}
				value, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err.Error())
				}
				currentTable.numbers[i][index] = value
				index++
			}
			scanner.Scan()
		}
		allBoards = append(allBoards, currentTable)
		_, err := fmt.Scan()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	for !isABoardWinning() {
		markAllBoards(numbersToCome[pos])
		pos++
	}
	lastNumber := numbersToCome[pos-1]
	winningBoard := getWinningBoard()
	sumAllUnmarked := sumAllUnmarked(winningBoard)
	fmt.Println(lastNumber, "last")
	fmt.Println(sumAllUnmarked, "sum")
	fmt.Println(lastNumber * sumAllUnmarked)
}

func sumAllUnmarked(board bingoBoard) int {
	result := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.markedNumbers[i][j] {
				result += board.numbers[i][j]
			}
		}
	}
	return result
}

func getWinningBoard() bingoBoard {
	for _, board := range allBoards {
		if isWinningBoard(board) {
			return board
		}
	}
	return bingoBoard{}
}

func markAllBoards(number int) {
	for index := 0; index < len(allBoards); index++ {
		board := allBoards[index]
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if board.numbers[i][j] == number {
					board.markedNumbers[i][j] = true
				}
			}
		}
		allBoards[index] = board
	}
}

func isABoardWinning() bool {
	for _, board := range allBoards {
		if isWinningBoard(board) {
			return true
		}
	}
	return false
}

func isWinningBoard(board bingoBoard) bool {
	// check all horizontals
	for i := 0; i < 5; i++ {
		if board.markedNumbers[i][0] && board.markedNumbers[i][1] && board.markedNumbers[i][2] && board.markedNumbers[i][3] && board.markedNumbers[i][4] {
			return true
		}
	}
	// check all verticals
	for i := 0; i < 5; i++ {
		if board.markedNumbers[0][i] && board.markedNumbers[1][i] && board.markedNumbers[2][i] && board.markedNumbers[3][i] && board.markedNumbers[4][i] {
			return true
		}
	}
	return false
}
