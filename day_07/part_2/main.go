package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

const filepath = "day_07/part_2/input.txt"

var bagPattern = regexp.MustCompile(`(^|(\d+)\s)(\b(\w+\s\w+)\s(bag(\s|s\s|,|s,|\.|s\.)))`)

type BagTuple struct {
	name     string
	contains int
}

func main() {
	start := time.Now()
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	read := time.Now()

	scanner := bufio.NewScanner(file)
	bagMap := map[string][]BagTuple{}
	for scanner.Scan() {
		raw := scanner.Text()
		matches := bagPattern.FindAllStringSubmatch(raw, -1)
		if len(matches) < 2 {
			continue
		}
		for i := 1; i < len(matches); i++ {
			numberOfBagsToContain, _ := strconv.Atoi(matches[i][2])
			bagMap[matches[0][4]] = append(bagMap[matches[0][4]], BagTuple{matches[i][4], numberOfBagsToContain})
		}
	}
	fmt.Printf("Read: %v\n", time.Now().Sub(read))
	p2 := time.Now()
	var numberOfBagsToHold = 0
	numberOfBagsToHold += CalculateNrBagsToHoldForKey(bagMap["shiny gold"], bagMap)
	fmt.Printf("P2: %v\n", time.Now().Sub(p2))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Sum can contain: %d\n", numberOfBagsToHold)
}

func CalculateNrBagsToHoldForKey(tuples []BagTuple, bagMap map[string][]BagTuple) int {
	nrForBag := 0
	for _, tuple := range tuples {
		nrForBag += tuple.contains * (1 + CalculateNrBagsToHoldForKey(bagMap[tuple.name], bagMap))
	}
	return nrForBag
}
