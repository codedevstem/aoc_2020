package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const filepath = "day_03/part_2/input.csv"

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
	if err != nil || len(forest) == 0 {
		panic(err)
	}
	treeSymbol := "#"
	rightIndex, currentRightIndex, downIndex, currentDownIndex, foundIndex := 0, 1, 2, 3, 4
	var searchGroups = map[int][]int{}
	searchGroups[0] = []int{1, 1, 1, 1, 0}
	searchGroups[1] = []int{3, 3, 1, 1, 0}
	searchGroups[2] = []int{5, 5, 1, 1, 0}
	searchGroups[3] = []int{7, 7, 1, 1, 0}
	searchGroups[4] = []int{1, 1, 2, 2, 0}

	for i := 1; i < len(forest); i++ {
		for _, searchGroup := range searchGroups {
			if i%searchGroup[downIndex] != 0 {
				continue
			}
			searchLoc := searchGroup[currentRightIndex] % len(forest[i])
			if string(forest[i][searchLoc]) == treeSymbol {
				searchGroup[foundIndex]++
			}

			if i%searchGroup[downIndex] == 1 {
				searchGroup[currentDownIndex] = searchGroup[currentDownIndex] + searchGroup[downIndex]
			}
			searchGroup[currentRightIndex] = searchGroup[currentRightIndex] + searchGroup[rightIndex]
		}
	}
	multiplied := int64(1)
	for _, inst := range searchGroups {
		fmt.Printf("found %d trees in slope [%d,%d]\n", inst[foundIndex], inst[rightIndex], inst[downIndex])
		multiplied = multiplied * int64(inst[foundIndex])
	}
	fmt.Printf("Trees Found: %d\n", multiplied)

}
