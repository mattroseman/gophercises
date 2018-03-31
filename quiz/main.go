package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	problemsFilePath := flag.String("csv", "problems.csv", "filepath to problems csv file")
	flag.Parse()

	// read in default csv file problems.csv
	problemsCsv, err := os.Open(*problemsFilePath)
	if err != nil {
		panic(err)
	}
	// iterate through and get problems and answers
	csvReader := csv.NewReader(problemsCsv)

	stdInReader := bufio.NewReader(os.Stdin)

	var numCorrect int
	var numQuestions int

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		problem := record[0]
		correctAnswer := record[1]

		// prompt the problem and read input
		fmt.Print("#" + strconv.Itoa(numQuestions) + ": " + problem + " = ")
		answer, _ := stdInReader.ReadString('\n')
		// knock off the newline
		answer = answer[:len(answer)-1]

		// compare input to answer and move on to next problem
		numQuestions++
		if answer == correctAnswer {
			numCorrect++
		}
	}

	fmt.Println("")
	fmt.Println("You scored", numCorrect, "out of", numQuestions)
}
