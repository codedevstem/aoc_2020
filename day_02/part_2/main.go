package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
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
	valid, invalid := 0, 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for _, record := range records {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *sync.Mutex, record string) {
			match := lineRegex.FindStringSubmatch(record)
			paramsMap := make(map[string]string)
			for i, name := range lineRegex.SubexpNames() {
				paramsMap[name] = match[i]
			}
			first, err := strconv.Atoi(paramsMap["First"])
			if err != nil {
				panic("Min not a number: " + err.Error())
			}
			second, err := strconv.Atoi(paramsMap["Second"])
			if err != nil {
				panic("Max not a number: " + err.Error())
			}
			letter := paramsMap["Letter"]
			password := paramsMap["Password"]

			m.Lock()
			if string(password[first-1]) == letter && string(password[second-1]) == letter {
				invalid++
			} else if string(password[first-1]) == letter || string(password[second-1]) == letter {
				valid++
			} else {
				invalid++
			}
			m.Unlock()

			wg.Done()
		}(&wg, &m, record[0])
	}
	wg.Wait()
	fmt.Printf("valid passwords: %d, invalid %d\n", valid, invalid)
}
