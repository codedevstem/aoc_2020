package main_test

import (
	"math"
	"strconv"
	"testing"
)

func TestParseIndex(t *testing.T) {
	currentSeatNumber := ParseIndex("RRR", 0, 7, "L", "R")
	if currentSeatNumber != 7 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_1(t *testing.T) {
	currentSeatNumber := ParseIndex("RRL", 0, 7, "L", "R")
	if currentSeatNumber != 6 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_2(t *testing.T) {
	currentSeatNumber := ParseIndex("RLR", 0, 7, "L", "R")
	if currentSeatNumber != 5 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_3(t *testing.T) {
	currentSeatNumber := ParseIndex("RLL", 0, 7, "L", "R")
	if currentSeatNumber != 4 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_4(t *testing.T) {
	currentSeatNumber := ParseIndex("LRR", 0, 7, "L", "R")
	if currentSeatNumber != 3 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_5(t *testing.T) {
	currentSeatNumber := ParseIndex("LRL", 0, 7, "L", "R")
	if currentSeatNumber != 2 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_6(t *testing.T) {
	currentSeatNumber := ParseIndex("LLR", 0, 7, "L", "R")
	if currentSeatNumber != 1 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
}
func TestParseIndex_7(t *testing.T) {
	currentSeatNumber := ParseIndex("LLL", 0, 7, "L", "R")
	if currentSeatNumber != 0 {
		panic("WRONG: was " + strconv.Itoa(currentSeatNumber))
	}
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
