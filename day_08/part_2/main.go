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

const filepath = "day_08/part_2/input.txt"

type Instruction struct {
	fun          string
	value        int
	visited      bool
	triedChanged bool
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
	var originalInstructions []Instruction
	for scanner.Scan() {
		rawParts := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(rawParts[1])
		originalInstructions = append(originalInstructions, Instruction{
			fun:   rawParts[0],
			value: value,
		})
	}
	fmt.Printf("Read: %v\n", time.Now().Sub(read))
	p2 := time.Now()
	accumulator := 0
	for true {
		accumulator = 0
		successRun := true
		triedToChangeInRun := false
		copyInstructions := append([]Instruction{}, originalInstructions...)
		for i := 0; i < len(originalInstructions); i++ {
			if copyInstructions[i].visited == true {
				successRun = false
				break
			}
			copyInstructions[i].visited = true
			command := copyInstructions[i].fun
			value := copyInstructions[i].value
			if command == "acc" {
				accumulator += value
			} else if command == "jmp" {
				if !triedToChangeInRun && copyInstructions[i].triedChanged == false {
					copyInstructions[i].fun = "nop"
					originalInstructions[i].triedChanged = true
					triedToChangeInRun = true
				} else {
					i += value - 1
				}
			} else if command == "nop" {
				if !triedToChangeInRun && copyInstructions[i].triedChanged == false {
					copyInstructions[i].fun = "jmp"
					originalInstructions[i].triedChanged = true
					triedToChangeInRun = true
					i += value - 1
				}
			}
		}

		if successRun {
			break
		}
	}

	fmt.Printf("P2: %v\n", time.Now().Sub(p2))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Accumulator: %d\n", accumulator)
}
