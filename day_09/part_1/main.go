package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	filepath = "day_09/part_1/input.txt"
)

func main() {
	start := time.Now()
	read := time.Now()
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
	var numbers []int64
	for scanner.Scan() {
		intNumber, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, int64(intNumber))
	}
	readFinished := time.Now().Sub(read)
	p1 := time.Now()
	weakNumber := int64(-1)
	weakNess := 25
	for i := weakNess; i < len(numbers); i++ {
		found := false
		suspect := numbers[i]
		for j := i - weakNess; j < i; j++ {
			first := numbers[j]
			for k := j + 1; k < i; k++ {
				if first+numbers[k] == suspect {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found == false {
			weakNumber = numbers[i]
			break
		}
	}

	fmt.Printf("Read: %v\n", readFinished)
	fmt.Printf("P1: %v\n", time.Now().Sub(p1))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("First Weak Number: %d\n", weakNumber)
}
