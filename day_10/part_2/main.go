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
	filepath = "day_10/part_2/test_2.txt"
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
		intNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Could not parse text to int")
		}
		adapters = append(adapters, intNumber)
	}
	readFinished := time.Now().Sub(read)
	p2 := time.Now()
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	maxNumberOfAdapters := len(adapters)
	minNumberOfAdapters := 0
	currentJoltage := 0
	for i := 0; i < len(adapters); i++ {
		if i+3 < len(adapters) && adapters[i+3]-currentJoltage <= 3 {
			i += 3
		}
		if i+2 < len(adapters) && adapters[i+2]-currentJoltage <= 3 {
			i += 1
		}
		if i+1 < len(adapters) && adapters[i+1]-currentJoltage <= 3 {
			i += 1
		}
		currentJoltage = adapters[i]
		minNumberOfAdapters++

	}

	fmt.Printf("Read: %v\n", readFinished)
	fmt.Printf("P2: %v\n", time.Now().Sub(p2))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Possible Arrangements: %f\n", (float64(maxNumberOfAdapters - minNumberOfAdapters)))
}
