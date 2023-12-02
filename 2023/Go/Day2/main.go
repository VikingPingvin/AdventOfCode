package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFilePath = "input.txt"
)

const (
	maxRedCubes   = 12
	maxGreenCubes = 13
	maxBlueCubes  = 14
)

var maxCubesMaps = map[string]int{
	"red":   maxRedCubes,
	"green": maxGreenCubes,
	"blue":  maxBlueCubes,
}

func main() {
	lines := readFileToLines(inputFilePath)

	solvePart1(lines)
	solvePart2(lines)

}

func solvePart1(lines []string) {
	sumOfIDs := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		gameLine := strings.Split(line, ":")
		gameId := getGameIdFromString(gameLine[0])

		sets := getGameSets(gameLine[1])
		setPossible := true

		for _, hand := range sets {
			handInfo := strings.Split(hand, " ")
			cubeNum, err := strconv.Atoi(handInfo[0])
			if err != nil {
				fmt.Printf("err %v, err")
			}
			cubeColor := handInfo[1]

			maxPossibleOfColor, ok := maxCubesMaps[cubeColor]
			if !ok {
				fmt.Printf("error getting color: %s", cubeColor)
			}

			if cubeNum > maxPossibleOfColor {
				setPossible = false
				break
			}

		}
		if setPossible {
			sumOfIDs += gameId
		}
	}
	fmt.Printf("Sum of possible Game IDs: %d\n", sumOfIDs)
}

func solvePart2(lines []string) {
	var minCubes = map[string]int{}
	sumOfPowers := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		minCubes = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		setPower := 1
		gameLine := strings.Split(line, ":")
		sets := getGameSets(gameLine[1])
		for _, hand := range sets {
			handInfo := strings.Split(hand, " ")
			cubeNum, err := strconv.Atoi(handInfo[0])
			if err != nil {
				fmt.Printf("err %v, err")
			}
			cubeColor := handInfo[1]

			if minCubes[cubeColor] < cubeNum {
				minCubes[cubeColor] = cubeNum
			}
		}
		for _, cubeValue := range minCubes {
			setPower *= cubeValue
		}
		sumOfPowers += setPower
	}
	fmt.Printf("Sum of Powers for Part2: %d", sumOfPowers)
}

func getGameSets(gameSetsLine string) []string {
	re := regexp.MustCompile(`(\d+\s(green|red|blue))`)
	matches := re.FindAllString(gameSetsLine, -1)
	return matches
}

func getGameIdFromString(input string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(input)
	if match != "" {
		gameId, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("err during game id regex: %v", err)
			return 0
		}
		return gameId
	}
	return 0
}

func readFileToLines(path string) []string {
	data, err := readFile(inputFilePath)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	return lines
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
