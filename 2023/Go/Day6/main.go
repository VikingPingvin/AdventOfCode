package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	filePath = "input.txt"
)

type RaceData struct {
	time     int
	distance int
}

func main() {
	lines, err := readFileToLines(filePath)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	raceData := getRaceData(lines)

	p1Result := solveP1(raceData)
	fmt.Printf("Result for Part1: %v\n", p1Result)

	raceDataP2 := getRaceDataP2(lines)
	p2Result := solveQuadratic(raceDataP2)
	// P1 solver can be used, but longer
	// p2Result := solveP1([]RaceData{raceDataP2})
	fmt.Printf("Result for Part2: %v\n", p2Result)
}

func solveP1(data []RaceData) int {
	sumOfPossibilities := 1

	for _, race := range data {
		currTime := race.time
		recordDist := race.distance

		numOfWins := 0

		for timeHeld := 0; timeHeld < currTime; timeHeld++ {
			speed := timeHeld
			runTime := currTime - timeHeld

			distanceCovered := speed * runTime

			if distanceCovered >= recordDist {
				numOfWins++
			}
		}
		if numOfWins >= 1 {
			sumOfPossibilities *= numOfWins
		}
	}

	return sumOfPossibilities
}

func solveQuadratic(data RaceData) int {
	// sumOfPossibilities := 0
	time := data.time
	dist := data.distance
	// s = v * t
	// t = T - v
	// s = v * (T - v) = vT - v^2
	//////// -v^2 + vT - s = 0 //////////
	// x = -b +- (b^2 - 4ac)^1/2  / 2a
	// x = -T +- (T^2 - 4 * 1 * D)^1/2 / -2
	// D = (b^2 - 4ac) ^1/2
	D := math.Pow((float64(time)*float64(time) - (float64(4) * float64(-1) * float64(-dist))), 0.5)
	p1 := (float64(-time) + D) / (-2)
	p2 := (float64(-time) - D) / (-2)
	return int(p2 - p1)
}

func getRaceData(lines []string) []RaceData {
	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	raceData := make([]RaceData, 0)
	for i := 0; i < len(times); i++ {
		currTime, _ := strconv.Atoi(times[i])
		currDist, _ := strconv.Atoi(distances[i])
		raceData = append(raceData, RaceData{time: currTime, distance: currDist})
	}

	return raceData
}

func getRaceDataP2(lines []string) RaceData {
	re := regexp.MustCompile(`\b(\d+)\b`)
	time := re.FindAllString(lines[0], -1)
	distance := re.FindAllString(lines[1], -1)

	combinedTime := strings.Join(time, "")
	combinedDistance := strings.Join(distance, "")

	timeNum, _ := strconv.Atoi(combinedTime)
	distanceNum, _ := strconv.Atoi(combinedDistance)

	raceData := RaceData{
		time:     timeNum,
		distance: distanceNum,
	}
	// for i := 0; i < len(times); i++ {
	// 	currTime, _ := strconv.Atoi(times[i])
	// 	currDist, _ := strconv.Atoi(distances[i])
	// 	raceData = append(raceData, RaceData{time: currTime, distance: currDist})
	// }

	return raceData
}

func readFileToLines(fileName string) ([]string, error) {
	filePath, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	data, err := readLines(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return data, nil
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
