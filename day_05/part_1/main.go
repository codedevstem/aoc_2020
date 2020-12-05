package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const filepath = "day_05/part_1/input.txt"

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
	scanner.Text()
	var maxSeatId int

	for scanner.Scan() {
		raw := scanner.Text()
		currentSeatRow := ParseIndex(raw[:7], 0, 127, "F", "B")
		currentSeatNumber := ParseIndex(raw[7:], 0, 7, "L", "R")
		currentSeatId := currentSeatRow*8 + currentSeatNumber
		if currentSeatId > maxSeatId {
			maxSeatId = currentSeatId
		}
	}

	fmt.Printf("Total: %v\n", time.Now().Sub(start))
	fmt.Printf("max seat number is: %d\n", maxSeatId)

}

func ParseIndex(s string, minAxis int, maxAxis int, down string, up string) int {
	if len(s) == 1 {
		if string(s[0]) == down {
			return minAxis
		} else if string(s[0]) == up {
			return maxAxis
		}
	}
	if string(s[0]) == down {
		return ParseIndex(s[1:], minAxis, maxAxis-(int(math.Ceil(float64(maxAxis)/float64(2)))-int(math.Ceil(float64(minAxis)/float64(2)))), down, up)
	} else if string(s[0]) == up {
		return ParseIndex(s[1:], int(math.Ceil(float64(maxAxis)/float64(2)))+int(math.Ceil(float64(minAxis)/float64(2))), maxAxis, down, up)
	} else {
		panic("Fucked up row parsing")
	}
}
