package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

const filepath = "input.csv"

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
	var inputs []int32
	for _, record := range records {
		value, err := strconv.Atoi(record[0])
		if err != nil {
			panic("could not convert record entry: " + err.Error())
		}
		inputs = append(inputs, int32(value))
	}
	found := false
	for i := 0; i < len(inputs); i++ {
		for j := i; j < len(inputs); j++ {
			if inputs[i]+inputs[j] == 2020 {
				fmt.Printf("The answer is %d * %d = %d\n", inputs[i], inputs[j], inputs[i]*inputs[j])
				found = true
				return
			}
		}
		if found {
			return
		}
	}
	if !found {
		fmt.Printf("No answer found")
	}

}
