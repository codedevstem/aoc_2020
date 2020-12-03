package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const filepath = "day_03/part_1/input.csv"

func main() {
	start := time.Now()
	defer fmt.Printf("duration: %v\n", time.Now().Sub(start))
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
	var forest []string
	for scanner.Scan() {
		forest = append(forest, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	increment := 3
	searchIndex := increment
	treeSymbol := "#"
	treesFound := 0
	for i := 1; i < len(forest); i++ {
		searchLoc := searchIndex % len(forest[i])
		if string(forest[i][searchLoc]) == treeSymbol {
			treesFound++
		}
		searchIndex = searchIndex + increment
	}
	fmt.Printf("found: %d trees", treesFound)

}
