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
	depths, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	area := solvePartOne(depths)
	fmt.Println(area)
}

type Command struct {
	Name  string
	Value int
}

func NewCommand(s string) Command {
	t := strings.Split(s, " ")
	name := strings.ToLower(t[0])
	value, _ := strconv.Atoi(t[1])

	return Command{Name: name, Value: value}
}

func (c Command) UpdateMeasurementPartOne(forward, depth int) (newForward, newDepth int) {
	switch c.Name {
	case "down":
		return forward, depth + c.Value
	case "up":
		return forward, depth - c.Value
	case "forward":
		return forward + c.Value, depth
	default:
		return forward, depth
	}
}

func (c Command) UpdateMeasurementPartTwo(aim, forward, depth int) (newAim, newForward, newDepth int) {
	switch c.Name {
	case "down":
		return aim + c.Value, forward, depth
	case "up":
		return aim - c.Value, forward, depth
	case "forward":
		return aim, forward + c.Value, depth + (aim * c.Value)
	default:
		return aim, forward, depth
	}
}

func getInput() ([]Command, error) {
	f, err := os.Open("./day2/depths.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var depths []Command
	for scanner.Scan() {
		depths = append(depths, NewCommand(scanner.Text()))
	}

	return depths, nil
}

func solvePartOne(commands []Command) int {
	var forward, depth int
	for _, command := range commands {
		forward, depth = command.UpdateMeasurementPartOne(forward, depth)
	}

	return forward * depth
}

func solvePartTwo(commands []Command) int {
	var aim, forward, depth int
	for _, command := range commands {
		aim, forward, depth = command.UpdateMeasurementPartTwo(aim, forward, depth)
	}

	return forward * depth
}
