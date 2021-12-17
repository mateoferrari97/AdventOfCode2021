package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var increments int

	depths, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	length := len(depths)
	for i := 1; i < length; i++ {
		if depths[i] > depths[i-1] {
			increments++
		}
	}

	fmt.Println(increments)
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
