package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const filepath = "day_04/part_1/input.txt"

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
	var passports []map[string]string
	var passport = map[string]string{}
	scanner.Text()
	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}
		for _, field := range strings.Split(raw, " ") {
			keyValuePair := strings.Split(field, ":")
			if keyValuePair[0] == "cid" {
				continue
			}
			passport[keyValuePair[0]] = keyValuePair[1]
		}
	}
	passports = append(passports, passport)
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read: %v\n", time.Now().Sub(start))

	p1Start := time.Now()

	validPasses := 0
	for _, pass := range passports {
		if len(pass) == 7 {
			validPasses++
		}
	}
	fmt.Printf("P1: %v\n", time.Now().Sub(p1Start))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("number of valid passports: %d\n", validPasses)

}
