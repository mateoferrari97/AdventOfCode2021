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
	bingo, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	score := solvePartTwo(bingo)
	fmt.Println(score)
}

type Bingo struct {
	boards  []Board
	numbers []int
}

type Board struct {
	elements [][]Element
	won      bool
}

type Element struct {
	Value int
	Seen  bool
}

func getInput() (Bingo, error) {
	f, err := os.Open("./day4/bingo.txt")
	if err != nil {
		return Bingo{}, err
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)
	_ = scanner.Scan()
	ns := strings.Split(scanner.Text(), ",")
	_ = scanner.Scan()

	numbers := make([]int, 0, len(ns))
	for _, n := range ns {
		number, _ := strconv.Atoi(n)
		numbers = append(numbers, number)
	}

	var boards []Board
	var rows [][]Element
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			var board Board
			for _, row := range rows {
				board.elements = append(board.elements, row)
			}

			boards = append(boards, board)
			rows = nil
		} else {
			numbers := strings.Split(text, " ")
			row := make([]Element, 0, len(numbers))
			for _, number := range numbers {
				if number == "" {
					continue
				}

				number, _ := strconv.Atoi(number)
				row = append(row, Element{
					Value: number,
					Seen:  false,
				})
			}

			rows = append(rows, row)
		}
	}

	if scanner.Text() == "" {
		var board Board
		for _, row := range rows {
			board.elements = append(board.elements, row)
		}

		boards = append(boards, board)
		rows = nil
	}

	return Bingo{
		boards:  boards,
		numbers: numbers,
	}, nil
}

func solvePartOne(bingo Bingo) int {
	for _, number := range bingo.numbers {
		for _, board := range bingo.boards {
			for _, elements := range board.elements {
				for i := range elements {
					if elements[i].Value == number {
						elements[i].Seen = true
						if haveSeenAllNumbers(board, elements, i) {
							return elements[i].Value * sumNotSeenNumbers(board)
						}
					}
				}
			}
		}
	}

	return 0
}

func solvePartTwo(bingo Bingo) int {
	var wins int
	totalBoards := len(bingo.boards)
	for _, number := range bingo.numbers {
		for bi, board := range bingo.boards {
			for _, elements := range board.elements {
				for i := range elements {
					if elements[i].Value == number {
						elements[i].Seen = true
						if haveSeenAllNumbers(board, elements, i) {
							if !board.won {
								wins++
								bingo.boards[bi].won = true
							}

							if totalBoards-wins == 0 {
								fmt.Println(elements[i].Value)
								return elements[i].Value * sumNotSeenNumbers(board)
							}
						}
					}
				}
			}
		}
	}

	return 0
}

func haveSeenAllNumbers(board Board, elements []Element, index int) bool {
	vSeen, hSeen := true, true
	for _, e := range board.elements {
		if !e[index].Seen {
			vSeen = false
		}
	}

	for _, element := range elements {
		if !element.Seen {
			hSeen = false
		}
	}

	return vSeen || hSeen
}

func sumNotSeenNumbers(board Board) int {
	var result int
	for _, elements := range board.elements {
		for _, element := range elements {
			if !element.Seen {
				result += element.Value
			}
		}
	}

	return result
}
