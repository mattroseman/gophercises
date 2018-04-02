package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	problemsFilePath := flag.String("csv", "problems.csv", "filepath to problems csv file")
	limit := flag.Int("limit", 30, "time limit for each question in seconds")
	flag.Parse()

	problems, err := readProblems(*problemsFilePath)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *problemsFilePath)
		os.Exit(1)
	}

	stdInReader := bufio.NewReader(os.Stdin)

	fmt.Print("Press enter to start")
	stdInReader.ReadString('\n')

	// start timer
	t := make(chan int)
	go startTimer(*limit, t)

	numQuestions := len(problems)
	numCorrect := startQuiz(problems, t)

	fmt.Println("")
	fmt.Println("You scored", numCorrect, "out of", numQuestions)
}

func startTimer(limit int, t chan int) {
	time.Sleep(time.Duration(limit) * time.Second)
	t <- limit
}

func startQuiz(problems []problem, t chan int) int {
	var numCorrect int

	// start quiz
	for i, problem := range problems {
		// prompt the problem and read input
		fmt.Print("#" + strconv.Itoa(i+1) + ": " + problem.question + " = ")

		a := make(chan string)
		go getAnswer(problem, a)

		select {
		case answer := <-a:
			// knock off the newline
			answer = answer[:len(answer)-1]

			// compare input to answer and move on to next problem
			if answer == problem.answer {
				numCorrect++
			}
		case <-t:
			return numCorrect
		}
	}

	return numCorrect
}

func getAnswer(p problem, a chan string) {
	stdInReader := bufio.NewReader(os.Stdin)
	answer, _ := stdInReader.ReadString('\n')
	a <- answer
}

func readProblems(problemsFilePath string) ([]problem, error) {
	// read in csv file problems.csv
	problemsCsv, err := os.Open(problemsFilePath)
	if err != nil {
		return nil, err
	}
	// iterate through and get problems and answers
	csvReader := csv.NewReader(problemsCsv)

	var problems []problem

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			return problems, nil
		}
		if err != nil {
			return nil, err
		}

		problems = append(problems, problem{question: record[0], answer: record[1]})
	}
}
