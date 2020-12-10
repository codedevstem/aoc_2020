package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	filepath = "day_10/part_1/test.txt"
)

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
	var adapters []int
	for scanner.Scan() {
		intNumber, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, intNumber)
	}
	readFinished := time.Now().Sub(read)
	p1 := time.Now()
	sort.Ints(adapters)
	currentJoltage := 0
	differenceOne, differenceThree := 0, 0
	for i := 0; i < len(adapters); i++ {
		nextJoltage := adapters[i]
		if nextJoltage-currentJoltage == 1 {
			differenceOne++
		} else {
			differenceThree++
		}
		currentJoltage = nextJoltage
	}
	differenceThree++

	fmt.Printf("Read: %v\n", readFinished)
	fmt.Printf("P1: %v\n", time.Now().Sub(p1))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Answeres multiplied: %d\n", differenceOne*differenceThree)
}
