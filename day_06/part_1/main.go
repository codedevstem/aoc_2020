package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const filepath = "day_06/part_1/input.txt"

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

	scanner := bufio.NewScanner(file)
	var sumYesAll = 0
	yesMapCurrentGroup := map[rune]int{}
	scanner.Text()
	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			sumYesAll += len(yesMapCurrentGroup)
			yesMapCurrentGroup = map[rune]int{}
			continue
		}
		for i := range raw {
			yesMapCurrentGroup[rune(raw[i])] = 1
		}
	}
	sumYesAll += len(yesMapCurrentGroup)

	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Sum answered yes: %d\n", sumYesAll)
}
