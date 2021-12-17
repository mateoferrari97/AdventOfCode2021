package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	depths, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	solution := solvePartTwo(depths)
	fmt.Println(solution)
}

func getInput() ([]int, error) {
	f, err := os.Open("./day1/depths.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var depths []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}

	return depths, nil
}

func solvePartOne(depths []int) int {
	length := len(depths)
	if length < 1 {
		return 0
	}

	var increments int
	for i := 1; i < length; i++ {
		if depths[i] > depths[i-1] {
			increments++
		}
	}

	return increments
}

func solvePartTwo(depths []int) int {
	length := len(depths)
	if length < 3 {
		return 0
	}

	prevMeasurement := depths[0] + depths[1] + depths[2]

	var increments int
	for i := 1; i+2 < length; i++ {
		currentMeasurement := depths[i] + depths[i+1] + depths[i+2]
		if currentMeasurement > prevMeasurement {
			increments++
		}

		prevMeasurement = currentMeasurement
	}

	return increments
}
