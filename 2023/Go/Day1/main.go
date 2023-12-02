package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

const (
	inputFilePath = "input.txt"
)

var wordToDigitMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	data, err := readFile(inputFilePath)
	if err != nil {
		os.Exit(1)
	}

	calibrationSum := 0

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		calibrationSum += getCalibrationValues(line)
	}

	fmt.Printf("Calibration Value Sum: %d\n", calibrationSum)
}

func getCalibrationValues(line string) int {
	if len(line) == 0 {
		return 0
	}
	line = strings.ReplaceAll(line, "\r", "")

	firstDigit, lastDigit := "0", "0"

	line = convertLineToDigits(line)

	firstDigit = string(line[0])
	lastDigit = string(line[len(line)-1])

	retNum, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		fmt.Print("error? %v", err)
	}

	return retNum
}

func convertLineToDigits(line string) string {
	for stringNumber, numericValue := range wordToDigitMap {
		line = strings.ReplaceAll(line, stringNumber, numericValue)
	}

	var strippedLine []rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			strippedLine = append(strippedLine, char)
		}
	}

	return string(strippedLine)
}

func readFile(path string) (lines []byte, err error) {
	filePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
