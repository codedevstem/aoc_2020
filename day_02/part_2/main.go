package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const filepath = "day_02/part_2/input.csv"

var lineRegex = regexp.MustCompile(`^(?P<First>\d+)-(?P<Second>\d+)\b\s(?P<Letter>\w):\s\b(?P<Password>\w+)$`)

func main() {
	start := time.Now()
	defer fmt.Printf("duration: %v\n", time.Now().Sub(start))
	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic("could not read records: " + err.Error())
	}
	valid := 0
	invalid := 0
	for _, record := range records {
		match := lineRegex.FindStringSubmatch(record[0])

		paramsMap := make(map[string]string)
		for i, name := range lineRegex.SubexpNames() {
			if i > 0 && i <= len(match) {
				paramsMap[name] = match[i]
			}
		}
		first, err := strconv.Atoi(paramsMap["First"])
		if err != nil {
			panic("Min not a numer: " + err.Error())
		}
		second, err := strconv.Atoi(paramsMap["Second"])
		if err != nil {
			panic("Max not a numer: " + err.Error())
		}
		letter := paramsMap["Letter"]
		password := paramsMap["Password"]

		if string(password[first-1]) == letter && string(password[second-1]) == letter {
			invalid++
			continue
		} else if string(password[first-1]) == letter || string(password[second-1]) == letter {
			valid++
			continue
		} else {
			invalid++
		}
	}
	fmt.Printf("valid passwords: %d, invalid %d\n", valid, invalid)
}
