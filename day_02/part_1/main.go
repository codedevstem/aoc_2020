package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const filepath = "input.csv"

var lineRegex = regexp.MustCompile(`^(?P<Min>\d+)-(?P<Max>\d+)\b\s(?P<Letter>\w):\s\b(?P<Password>\w+)$`)

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
				if i > 0 && i <= len(match) {
					paramsMap[name] = match[i]
				}
			}
			min, err := strconv.Atoi(paramsMap["Min"])
			if err != nil {
				panic("Min not a numer: " + err.Error())
			}
			max, err := strconv.Atoi(paramsMap["Max"])
			if err != nil {
				panic("Max not a numer: " + err.Error())
			}
			letter := paramsMap["Letter"]
			password := paramsMap["Password"]
			occurrences := strings.Count(password, letter)
			m.Lock()
			if occurrences >= min && occurrences <= max {
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
