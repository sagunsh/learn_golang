package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//var limit int
	//flag.IntVar(&limit, "limit", 50, "time limit for the quiz in seconds")

	var csv_file string
	flag.StringVar(&csv_file, "csv", "problems.csv", "csv file in the format question,answer")
	flag.Parse()

	f, err := os.Open(csv_file)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	csvReader := csv.NewReader(f)

	i := 1
	score := 0
	total := 0

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		question := rec[0]
		answer := rec[1]
		fmt.Printf("Problem %d: %s = ", i, question)
		//_, err = fmt.Scan(&user_answer)

		reader := bufio.NewReader(os.Stdin)
		user_answer, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(strings.TrimSpace(user_answer)) == strings.ToLower(strings.TrimSpace(answer)) {
			score++
		}
		total++
	}
	fmt.Printf("You scored %d out of %d\n", score, total)
	os.Exit(0)
}
