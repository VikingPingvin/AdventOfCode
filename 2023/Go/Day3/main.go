package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/maps"
)

const fileName = "input.txt"

func main() {
	filePath, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	data, err := readLines(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	p1Result, symbolCollection := solveP1(data)
	fmt.Printf("Result for P1: %d\n", p1Result)

	p2Result := solveP2(symbolCollection, data)
	fmt.Printf("Result for P2: %d\n", p2Result)

}
func solveP2(symbolCollection map[int]map[int][]int, lines []string) int {
	sumOfPartValues := 0
	for lineNumber, symbols := range symbolCollection {
		for symbolPosition, parts := range symbols {
			line := lines[lineNumber]
			symbol := line[symbolPosition]
			if string(symbol) != "*" {
				continue
			}
			if len(parts) == 2 {
				valueOfParts := parts[0] * parts[1]
				sumOfPartValues += valueOfParts
			}
			continue
		}
	}
	return sumOfPartValues
}

func solveP1(lines []string) (int, map[int]map[int][]int) {
	// Store found ID positions in a map
	var foundIDPositions = make(map[int][]int, 0)
	var foundPartNumbers = []int{}
	var partsOfSymbols = make(map[int]map[int][]int, 0)
	var foundPartsForCurrentSymbol = make([]int, 0)

	symbolCoordinates := getSymbolCoordinates(lines)

	// Sort map keys because why not
	keys := maps.Keys(symbolCoordinates)
	slices.Sort(keys)

	for _, symbolLine := range keys {
		symbols, ok := symbolCoordinates[symbolLine]
		if !ok {
			fmt.Printf("cant get coordinates from %v", symbolLine)
		}

		for _, symbolPos := range symbols {
			fullLineOfSymbol := []byte(lines[symbolLine])
			currentLine := symbolLine - 1
			for currentLine <= symbolLine+1 {
				if currentLine < 0 || currentLine > len(lines) {
					break
				}

				currentPos := symbolPos - 1
				charsLine := []byte(lines[currentLine])
				for currentPos >= symbolPos-1 && currentPos <= symbolPos+1 {
					currentChar := charsLine[currentPos]
					if unicode.IsNumber(rune(currentChar)) {
						if slices.Contains(foundIDPositions[currentLine], currentPos) {
							break
						}
						positions, partNumber := getPartNumbers(currentPos, charsLine)
						foundIDPositions[currentLine] = append(foundIDPositions[currentLine], positions...)
						foundPartNumbers = append(foundPartNumbers, partNumber)

						// Add to part number to * symbol collection
						if string(fullLineOfSymbol[symbolPos]) == "*" {
							foundPartsForCurrentSymbol = append(foundPartsForCurrentSymbol, partNumber)

						}
					}
					currentPos++
				}
				currentLine++
			}

			if len(foundPartsForCurrentSymbol) != 0 {
				symbolsForLine := partsOfSymbols[symbolLine]
				if _, exists := symbolsForLine[symbolPos]; !exists {
					if _, ok := partsOfSymbols[symbolLine]; !ok {
						partsOfSymbols[symbolLine] = make(map[int][]int, 0)
					}
					sliceCopy := make([]int, len(foundPartsForCurrentSymbol))
					copy(sliceCopy, foundPartsForCurrentSymbol)
					partsOfSymbols[symbolLine][symbolPos] = sliceCopy

				}
				foundPartsForCurrentSymbol = foundPartsForCurrentSymbol[:0]
			}
		}
	}

	// Sum machine IDs
	sum := 0
	for _, partNumber := range foundPartNumbers {
		sum += partNumber
	}

	return sum, partsOfSymbols
}

func getSymbolCoordinates(lines []string) map[int][]int {
	x, y := 0, 0
	var symbolCoordinates = make(map[int][]int, 0)
	for _, line := range lines {
		x = 0
		for _, char := range line {
			if isExtendedSymbol(char) {
				symbolCoordinates[y] = append(symbolCoordinates[y], x)
			}
			x++
		}
		y++
	}
	return symbolCoordinates
}

func getPartNumbers(pos int, line []byte) (bytePositions []int, numbers int) {
	bytePositions = append(bytePositions, pos)

	for i := pos - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(line[i])) {
			bytePositions = append(bytePositions, i)
		} else {
			break
		}
	}
	for i := pos + 1; i <= len(line)-1; i++ {
		if unicode.IsNumber(rune(line[i])) {
			bytePositions = append(bytePositions, i)
		} else {
			break
		}
	}
	slices.Sort(bytePositions)

	// Get the int value
	strSlice := make([]string, len(bytePositions))
	for i, position := range bytePositions {
		strSlice[i] = string(line[position])
	}
	numString := ""
	for _, str := range strSlice {
		numString += str
	}
	resultInt, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Printf("error converting part number to integer. %v", err)
	}

	return bytePositions, resultInt
}

func isExtendedSymbol(char rune) bool {
	if char >= 33 && char <= 65 && char != 46 && (char < 48 || char > 57) {
		return true
	}
	return false
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Trim(line, "\t\n\r"))
	}
	return lines, err
}
