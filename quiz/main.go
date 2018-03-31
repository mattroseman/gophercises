package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// read in default csv file problems.csv
	problemsCsv, err := os.Open("./problems.csv")
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

		// TODO prompt the problem and read input
		fmt.Print(problem + ": ")
		answer, _ := stdInReader.ReadString('\n')
		// knock off the newline
		answer = answer[:len(answer)-1]

		// TODO compare input to answer and move on to next problem
		numQuestions++
		if answer == correctAnswer {
			numCorrect++
		}
	}

	fmt.Println("")
	fmt.Println("Number Correct =", numCorrect)
	fmt.Println("Number Wrong = ", numQuestions-numCorrect)
}
