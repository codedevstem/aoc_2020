package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Printf("duration: %v\n", time.Now().Sub(start))
	file, err := os.Open("day_01/part_2/input.csv")
	if err != nil {
		panic(err.Error())
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic("could not read records: " + err.Error())
	}
	var inputs []int64
	for _, record := range records {
		value, err := strconv.Atoi(record[0])
		if err != nil {
			panic("could not convert record entry: " + err.Error())
		}
		inputs = append(inputs, int64(value))
	}
	found := false
	for i := 0; i < len(inputs); i++ {
		for j := i; j < len(inputs); j++ {
			for k := j; k < len(inputs); k++ {
				if inputs[i]+inputs[j]+inputs[k] == 2020 {
					fmt.Printf("The answer is %d * %d * %d = %d\n", inputs[i], inputs[j], inputs[k], inputs[i]*inputs[j]*inputs[k])
					found = true
					return
				}
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
