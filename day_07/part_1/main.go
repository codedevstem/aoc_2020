package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"time"
)

const filepath = "day_07/part_1/input.txt"

var bagPattern = regexp.MustCompile(`(\b(\w+\s\w+)\s(bag(\s|s\s|,|s,|\.|s\.)))`)

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
	bagMap := map[string][]string{}
	for scanner.Scan() {
		raw := scanner.Text()
		matches := bagPattern.FindAllStringSubmatch(raw, -1)
		if matches[1][2] == "no other" {
			continue
		}
		for i := 1; i < len(matches); i++ {
			bagMap[matches[0][2]] = append(bagMap[matches[0][2]], matches[i][2])
		}
	}
	fmt.Printf("Read: %v\n", time.Now().Sub(read))
	p1 := time.Now()

	var canHold []string
	for key, val := range bagMap {
		for _, branch := range val {
			if branch == "shiny gold" || CheckIfItemIsInMap("shiny gold", bagMap, branch) {
				canHold = append(canHold, key)
				break
			}
		}
	}
	sort.Strings(canHold)
	fmt.Printf("P1: %v\n", time.Now().Sub(p1))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("Sum can contain: %d\n", len(canHold))
}

func CheckIfItemIsInMap(searchItem string, bagPatternMap map[string][]string, leaf string) bool {
	branch := bagPatternMap[leaf]
	for _, foundLeaf := range branch {
		if foundLeaf == searchItem {
			return true
		}
		res := CheckIfItemIsInMap(searchItem, bagPatternMap, foundLeaf)
		if res {
			return true
		}
	}
	return false
}
