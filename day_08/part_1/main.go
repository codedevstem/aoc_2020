package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const filepath = "day_08/part_1/input.txt"

type Instruction struct {
	fun     string
	value   int
	visited bool
}

func main() {
	start := time.Now()
	read := time.Now()
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var instructions []Instruction
	for scanner.Scan() {
		rawParts := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(rawParts[1])
		instructions = append(instructions, Instruction{
			fun:     rawParts[0],
			value:   value,
			visited: false,
		})
	}
	readFinished := time.Now().Sub(read)
	p1 := time.Now()
	accumulator := 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i].visited == true {
			break
		}
		instructions[i].visited = true
		command := instructions[i].fun
		value := instructions[i].value
		if command == "acc" {
			accumulator += value
		} else if command == "jmp" {
			i += value - 1
		}

	}
	fmt.Printf("Read: %v\n", readFinished)
	fmt.Printf("P1: %v\n", time.Now().Sub(p1))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Accumulator: %d\n", accumulator)
}
