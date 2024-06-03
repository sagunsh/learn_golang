package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// filename := "cities.csv"
	filename := os.Args[1]

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	max_lens := make([]int, len(records[0]))
	for _, row := range records {
		for col, value := range row {
			if len(value) > max_lens[col] {
				max_lens[col] = len(value)
			}
		}
	}

	total_len := 0
	for _, value := range max_lens {
		total_len += value + 3
	}
	total_len -= 1

	fmt.Print(" +" + strings.Repeat("-", total_len) + "+")
	fmt.Println()
	for row_num, row := range records {
		// separator for header (first row)
		if row_num == 1 {
			fmt.Print(" +" + strings.Repeat("-", total_len) + "+")
			fmt.Println()
		}
		for col_num, value := range row {
			extra := max_lens[col_num] - len(value)
			if col_num == 0 {
				fmt.Print(" | ")
			}
			fmt.Print(value + strings.Repeat(" ", extra) + " | ")
		}
		fmt.Println()
	}

	fmt.Print(" +" + strings.Repeat("-", total_len) + "+")
	fmt.Println()
}
