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
	filepath = "day_09/part_2/input.txt"
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
	p1Finished := time.Now().Sub(p1)

	p2 := time.Now()
	smallest, largest := -1, -1
	for i := 0; i < len(numbers); i++ {
		iNumber := numbers[i]
		currentSmallest := iNumber
		currentLargest := iNumber
		currentSum := iNumber
		for j := i + 1; j < len(numbers); j++ {
			jNumber := numbers[j]
			if jNumber > currentLargest {
				currentLargest = jNumber
			} else if jNumber < currentSmallest {
				currentSmallest = jNumber
			}
			currentSum += jNumber
			if currentSum == weakNumber {
				smallest = int(currentSmallest)
				largest = int(currentLargest)
			}
			if currentSum > weakNumber {
				break
			}

		}
	}

	fmt.Printf("Read: %v\n", readFinished)
	fmt.Printf("P1: %v\n", p1Finished)
	fmt.Printf("P2: %v\n", time.Now().Sub(p2))
	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("smallest: %d + largest: %d = %d\n", smallest, largest, smallest+largest)
}
