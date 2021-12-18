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
	binaries, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	b := solvePartTwo(binaries)
	fmt.Println(b)
}

func getInput() ([]string, error) {
	f, err := os.Open("./day3/diagnostic.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var binaries []string
	for scanner.Scan() {
		binaries = append(binaries, scanner.Text())
	}

	return binaries, nil
}

func solvePartOne(binaries []string) int {
	n := len(binaries)
	m := len(binaries[0])
	if n == 0 || m == 0 {
		return 0
	}

	var gamma, epsilon []string
	for col := 0; col < m; col++ {
		var zerosFreq, onesFreq int
		for row := 0; row < n; row++ {
			switch binaries[row][col] {
			case '0':
				zerosFreq += 1
			case '1':
				onesFreq += 1
			}
		}

		if onesFreq > zerosFreq {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}

	g, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	e, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)

	return int(g * e)
}

func solvePartTwo(binaries []string) int {
	oxygenRating := calculateOxygenRating(binaries, 0)
	co2Rating := calculateCo2Rating(binaries, 0)

	oxygen, _ := strconv.ParseInt(strings.Join(oxygenRating, ""), 2, 64)
	co2, _ := strconv.ParseInt(strings.Join(co2Rating, ""), 2, 64)

	return int(oxygen * co2)
}

func calculateOxygenRating(binaries []string, index int) []string {
	length := len(binaries)
	if length == 1 {
		return binaries
	}

	zerosFreq, onesFreq := calculateFrequencyAtIndex(binaries, index)

	mostCommon := '0'
	if onesFreq >= zerosFreq {
		mostCommon = '1'
	}

	newBinaries := getBinariesWithValue(binaries, mostCommon, index)
	return calculateOxygenRating(newBinaries, index+1)
}

func calculateCo2Rating(binaries []string, index int) []string {
	length := len(binaries)
	if length == 1 {
		return binaries
	}

	zerosFreq, onesFreq := calculateFrequencyAtIndex(binaries, index)

	leastCommon := '0'
	if onesFreq < zerosFreq {
		leastCommon = '1'
	}

	newBinaries := getBinariesWithValue(binaries, leastCommon, index)
	return calculateCo2Rating(newBinaries, index+1)
}

func calculateFrequencyAtIndex(binaries []string, index int) (zerosFreq, onesFreq int) {
	for _, binary := range binaries {
		switch binary[index] {
		case '0':
			zerosFreq++
		case '1':
			onesFreq++
		}
	}

	return
}

func getBinariesWithValue(binaries []string, value rune, index int) []string {
	var result []string
	for _, binary := range binaries {
		if rune(binary[index]) == value {
			result = append(result, binary)
		}
	}

	return result
}
