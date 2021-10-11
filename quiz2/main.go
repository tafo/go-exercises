package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()
	csvFile, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Could not load the file %s", *csvFileName)
		return
	}
	fmt.Println("Please try to guess the following questions")
	defer csvFile.Close()
	var correct, incorrect int32 = 0, 0
	csvLines, _ := csv.NewReader(csvFile).ReadAll()
	problems := parseLines(csvLines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problemLoop:
	for _, problem := range problems {
		fmt.Printf("%s = ", problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
			case <-timer.C:
				fmt.Println()
				break problemLoop
		case answer:= <- answerCh:
			if problem.answer != answer {
				incorrect++
			} else {
				correct++
			}
		}
	}
	fmt.Printf("Correct Answers : %d, Incorrect answers : %d", correct, incorrect)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem {
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer string
}