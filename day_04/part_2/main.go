package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const filepath = "day_04/part_2/input.txt"

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
	var allValidatorRules = map[string]*regexp.Regexp{}
	allValidatorRules["byr"] = regexp.MustCompile(`^(19[2-9][0-9])|(200[0-2])$`)
	allValidatorRules["iyr"] = regexp.MustCompile(`^(201[0-9]|2020)$`)
	allValidatorRules["eyr"] = regexp.MustCompile(`^(202[0-9]|2030)$`)
	allValidatorRules["hgt"] = regexp.MustCompile(`^(1([5-8][0-9]|9[0-3])cm)|((59|6[0-9]|7[0-6])in)$`)
	allValidatorRules["hcl"] = regexp.MustCompile(`^(#([0-9a-f]){6})$`)
	allValidatorRules["ecl"] = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	allValidatorRules["pid"] = regexp.MustCompile(`^([0-9]{9})$`)

	validPasses := 0
	for _, pass := range passports {
		valid := true
		if len(pass) != 7 {
			continue
		}
		for key, field := range pass {
			rule := allValidatorRules[key]
			if !rule.MatchString(field) {
				valid = false
			}
		}
		if valid {
			validPasses++
		}
	}
	fmt.Printf("P1: %v\n", time.Now().Sub(p1Start))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("number of valid passports: %d\n", validPasses)

}
