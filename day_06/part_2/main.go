package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const filepath = "day_06/part_2/input.txt"

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
	peopleInCurrentGroup := 0
	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			sumYesAll += countCurrentGroupMap(yesMapCurrentGroup, peopleInCurrentGroup)
			peopleInCurrentGroup = 0
			yesMapCurrentGroup = map[rune]int{}
			continue
		}
		for i := range raw {
			yesMapCurrentGroup[rune(raw[i])]++
		}
		peopleInCurrentGroup++
	}
	sumYesAll += countCurrentGroupMap(yesMapCurrentGroup, peopleInCurrentGroup)

	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Sum answered yes: %d\n", sumYesAll)
}

func countCurrentGroupMap(yesMapCurrentGroup map[rune]int, peopleInCurrentGroup int) int {
	yesAllGroup := 0
	for _, num := range yesMapCurrentGroup {
		if num == peopleInCurrentGroup {
			yesAllGroup++
		}
	}
	return yesAllGroup

}
