package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	csvFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println(err)
		// Eğer hata olmuşsa csv dosyasını kapatmaya gerek var mı?
		return
	}
	fmt.Println("Please try to guess the following questions")
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {

		}
	}(csvFile)
	var correct, incorrect int32 = 0, 0
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	for _, line := range csvLines {
		fmt.Printf("%s = ", line[0])
		answer, _ := strconv.Atoi(line[1])
		var i int
		_, err := fmt.Scan(&i)
		if err != nil || i != answer {
			incorrect++
		} else {
			correct++
		}
	}
	fmt.Printf("Correct Answers : %d, Incorrect answers : %d", correct, incorrect)
}