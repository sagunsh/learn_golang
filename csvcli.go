package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func get_random_numbers(n, x int) []int {
	random_numbers := make([]int, n)
	if x <= n {
		for i := 0; i < n; i++ {
			random_numbers[i] = i
		}
	} else {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			random_numbers[i] = rand.Intn(x)
		}
	}
	return random_numbers
}

func main() {
	args := os.Args[2:]
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	headers := records[0]
	data := records[1:]

	for idx, element := range args {
		el := strings.ToLower(element)
		if el == "help" {
			fmt.Println("Usage     - ./csvcli test.csv cli_arg [additional args]")
			fmt.Println("Example 1 - ./csvcli test.csv headers")
			fmt.Println("Example 2 - ./csvcli test.csv tail 5")
			fmt.Println("\nArguments")
			fmt.Println("args     - description - example")
			fmt.Println("describe - describes the csv file i.e. print headers, number of rows and columns, count, etc - ./csvcli test.csv describe")
			fmt.Println("headers  - prints the header - ./csvcli test.csv headers")
			fmt.Println("count    - prints the count - ./csvcli test.csv headers")
			fmt.Println("head     - prints first n rows if n > total rows otherwise prints all rows - ./csvcli test.csv head 5")
			fmt.Println("tail     - prints last n rows if n > total rows otherwise prints all rows - ./csvcli test.csv tail 5")
			fmt.Println("sample   - prints randomly selected n rows if n > total rows otherwise prints all rows - ./csvcli test.csv sample 5")
		} else if el == "describe" {
			fmt.Println("row count:", len(data))
			fmt.Println("col count:", len(headers))
			fmt.Println("headers:", strings.Join(headers, ","))
		} else if el == "headers" {
			for idx, text := range headers {
				fmt.Printf("%d %s\n", idx, text)
			}
		} else if el == "count" {
			fmt.Println(len(data))
		} else if el == "head" {
			sample_size := 5
			if len(args) > idx+1 {
				sample_size, err = strconv.Atoi(args[idx+1])
				if err != nil {
					panic(err)
				}
			}
			sample_size = min(sample_size, len(data))
			fmt.Println(strings.Join(headers, ","))
			for i := 0; i < sample_size; i++ {
				fmt.Println(strings.Join(data[i], ","))
			}
		} else if el == "tail" {
			sample_size := 5
			if len(args) > idx+1 {
				sample_size, err = strconv.Atoi(args[idx+1])
				if err != nil {
					panic(err)
				}
			}
			sample_size = min(sample_size, len(data))
			fmt.Println(strings.Join(headers, ","))
			for i := len(data) - sample_size; i < len(data); i++ {
				fmt.Println(strings.Join(data[i], ","))
			}
		} else if el == "sample" {
			sample_size := 5
			if len(args) > idx+1 {
				sample_size, err = strconv.Atoi(args[idx+1])
				if err != nil {
					panic(err)
				}
			}
			sample_size = min(sample_size, len(data))
			indexes := get_random_numbers(sample_size, len(data))
			fmt.Println(strings.Join(headers, ","))
			for _, i := range indexes {
				fmt.Println(strings.Join(data[i], ","))
			}
		}
	}
}
