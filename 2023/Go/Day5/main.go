package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"
)

const fileName = "input.txt"

func main() {
	lines, err := readFileToLines(fileName)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}

	p1TimingStart := time.Now()
	p1Result := solveP1(lines)
	p1TimingDuration := time.Since(p1TimingStart)
	fmt.Printf("Result of P1: %d\n--Took: %v\n", p1Result[0], p1TimingDuration)

	// P2 Brute
	p2BruteTimingStart := time.Now()
	p2BruteResult := solveP2BruteThread(lines)
	p2BruteTimingDuration := time.Since(p2BruteTimingStart)
	fmt.Printf("Result of P2: %d\n--Took: %v\n", p2BruteResult, p2BruteTimingDuration)

	// p2TimingStart := time.Now()
	// p2Result := solveP2(lines)
	// p2TimingDuration := time.Since(p2TimingStart)
	// fmt.Printf("Result of P2: %d\n--Took: %v\n", p2Result, p2TimingDuration)

}

func solveP2Brute(input []string) int {
	seedLine := input[0]
	seeds := getNumbers(seedLine)

	mapCollection := getMapCollection(input)
	lowestLocation := -1

	for seedRangeIndex := 0; seedRangeIndex < len(seeds); seedRangeIndex += 2 {
		currentSeedRangeStart := seeds[seedRangeIndex]
		currentSeedRange := seeds[seedRangeIndex+1]
		maxSeedForCurrentRange := currentSeedRangeStart + currentSeedRange - 1

		fmt.Printf("--DEBUG-- Solving for Seed range %v\n", currentSeedRangeStart)

		for seedIndex := 0; seedIndex <= maxSeedForCurrentRange; seedIndex++ {
			curentSeed := []int{currentSeedRangeStart + seedIndex}
			fmt.Printf("--DEBUG-- current seed %v\n", curentSeed)

			locationsForSeeds := getSeedLocationMap(curentSeed, mapCollection)
			if lowestLocation == -1 ||
				locationsForSeeds[curentSeed[0]] < lowestLocation {
				lowestLocation = locationsForSeeds[curentSeed[0]]
			}
		}
	}

	return lowestLocation
}

type safeInt struct {
	mu    sync.RWMutex
	value int64
}

func (s *safeInt) Value() int {
	return int(atomic.LoadInt64(&s.value))
}
func (s *safeInt) Set(val int) {

	atomic.StoreInt64(&s.value, int64(val))

}

func solveP2BruteThread(input []string) int {
	seedLine := input[0]
	seeds := getNumbers(seedLine)

	mapCollection := getMapCollection(input)
	lowestLocation := safeInt{value: -1}
	var wg sync.WaitGroup

	for seedRangeIndex := 0; seedRangeIndex < len(seeds); seedRangeIndex += 2 {
		currentSeedRangeStart := seeds[seedRangeIndex]
		currentSeedRange := seeds[seedRangeIndex+1]
		maxSeedForCurrentRange := currentSeedRangeStart + currentSeedRange - 1

		fmt.Printf("--DEBUG-- Solving for Seed range %v\n", currentSeedRangeStart)

		for seedIndex := 0; seedIndex <= maxSeedForCurrentRange; seedIndex++ {
			curentSeed := []int{currentSeedRangeStart + seedIndex}

			wg.Add(1)
			go func() {
				defer wg.Done()
				locationsForSeeds := getSeedLocationMap(curentSeed, mapCollection)
				safeLocationValue := lowestLocation.Value()
				if safeLocationValue == -1 ||
					locationsForSeeds[curentSeed[0]] < safeLocationValue {
					lowestLocation.Set(locationsForSeeds[curentSeed[0]])
				}
				fmt.Printf("--DEBUG--   Solved: %v\n", curentSeed)

			}()
		}
	}
	wg.Wait()

	return lowestLocation.Value()
}

func solveP2(input []string) int {
	seedLine := input[0]
	seedNumbers := getNumbers(seedLine)

	var seeds = &RangeSeeds{}

	for i := 0; i < len(seedNumbers); i += 2 {
		seedRange := &SeedRange{
			min: seedNumbers[i],
			max: seedNumbers[i] + seedNumbers[i+1] - 1,
		}
		seeds.Seeds = append(seeds.Seeds, *seedRange)
	}

	mapCollection := getMapCollection(input)
	lowestLocation := getSeedRangeLocationMap(seeds, mapCollection)

	return lowestLocation
}

func solveP1(input []string) []int {
	seedLine := input[0]
	seeds := getNumbers(seedLine)

	mapCollection := getMapCollection(input)
	locationsForSeeds := getSeedLocationMap(seeds, mapCollection)

	// Get min of seed-location mapcollection
	var values = make([]int, 0)
	for _, value := range locationsForSeeds {
		values = append(values, value)
	}
	sort.Ints(values)

	return values
}

type SourceDestinationMap struct {
	name string
	maps []Mapping
}
type Mapping struct {
	src  int
	dest int
	rng  int
}

// P2
type RangeSeeds struct {
	Seeds []SeedRange
}

// P2
type SeedRange struct {
	min int
	max int
}

func getMapCollection(input []string) map[int]SourceDestinationMap {
	indexOfMapCategory := 0
	var mapCollection = make(map[int]SourceDestinationMap)
	srcDstMap := &SourceDestinationMap{}

	for i := 1; i < len(input); i++ {
		currentLine := input[i]

		// Mapping border
		if len(currentLine) <= 0 {
			if len(srcDstMap.maps) != 0 {
				mapCollection[indexOfMapCategory] = *srcDstMap
				indexOfMapCategory++
			}
			continue
		}

		// If first character is non-numeric it is a SourceDestinationMap
		// beginning
		if unicode.IsLetter(rune(currentLine[0])) {
			splitString := strings.Split(currentLine, " ")
			srcDstMap = &SourceDestinationMap{}
			srcDstMap.name = string(splitString[0])
			continue
		} else {
			numbers := getNumbers(currentLine)
			mappingRange := &Mapping{
				dest: numbers[0],
				src:  numbers[1],
				rng:  numbers[2],
			}

			srcDstMap.maps = append(srcDstMap.maps, *mappingRange)
			continue
		}
	}

	// Add last map to collection
	mapCollection[indexOfMapCategory] = *srcDstMap
	return mapCollection
}

func getSeedLocationMap(seeds []int, mapCollection map[int]SourceDestinationMap) map[int]int {
	var locationsForSeeds = make(map[int]int)
	for _, seed := range seeds {
		interimLocation := seed
		for i := 0; i < len(mapCollection); i++ {
			currentMap := mapCollection[i]
			// for _, currentMap := range mapCollection {
			for _, mapping := range currentMap.maps {
				srcDstDiff := mapping.dest - mapping.src

				if interimLocation <= (mapping.src+mapping.rng-1) &&
					interimLocation >= mapping.src {
					interimLocation += srcDstDiff
					break
				}
			}
			// InterimLocation not found in mapping, it will stay the same
		}
		locationsForSeeds[seed] = interimLocation
	}
	return locationsForSeeds
}

func getSeedRangeLocationMap(seeds *RangeSeeds, mapCollection map[int]SourceDestinationMap) int {
	lowestLocation := -1
	var locationRanges = make([]int, 0)
	for _, seedRange := range seeds.Seeds {
		interimLocationRanges := []int{seedRange.min, seedRange.max}
		fmt.Printf("--DEBUG-- Solving for Seed range %v\n", interimLocationRanges)
		for iMap := 0; iMap < len(mapCollection); iMap++ {
			currentMap := mapCollection[iMap]
			skipIndex := []int{}
			for _, mapping := range currentMap.maps {
				srcDstDiff := mapping.dest - mapping.src
				sourceRangeMin := mapping.src
				sourceRangeMax := mapping.src + mapping.rng - 1

				// If an interval aws added, we don't want to retest it
				// as it already was modified
				for i := 0; i < len(interimLocationRanges); i += 2 {
					if slices.Contains(skipIndex, i) {
						continue
					}

					interimRangeMin := interimLocationRanges[i]
					interimRangeMax := interimLocationRanges[i+1]

					// // seedrange == sourcerange
					// if interimRangeMin == sourceRangeMin&&
					//   interimRangeMax == sourceRangeMax{

					// }

					// whole seedrange is contained within source range
					if interimRangeMin >= sourceRangeMin &&
						interimRangeMax <= sourceRangeMax {
						interimLocationRanges[i] += srcDstDiff
						interimLocationRanges[i+1] += srcDstDiff
						continue
					}
					// Seedrange contains whole source range
					// if interimRangeMin <= sourceRangeMin ||
					// 	interimRangeMax >= sourceRangeMax {
					// 	newValues := make([]int, 0)
					// 	// seedrange contains sourcerange no equality
					// 	if interimRangeMin < sourceRangeMin &&
					// 		interimRangeMax > sourceRangeMax {
					// 		newValues = []int{
					// 			interimRangeMin,
					// 			interimRangeMin + sourceRangeMin - 1,
					// 			sourceRangeMin + srcDstDiff,
					// 			sourceRangeMax + srcDstDiff,
					// 			sourceRangeMax + 1,
					// 			interimRangeMax,
					// 		}
					// 	} else if interimRangeMin == sourceRangeMin {
					// 		newValues = []int{
					// 			interimRangeMin + srcDstDiff,
					// 			sourceRangeMax + srcDstDiff,
					// 			sourceRangeMax + 1,
					// 			interimRangeMax,
					// 		}
					// 	} else if interimRangeMax == sourceRangeMax {
					// 		newValues = []int{
					// 			interimRangeMin,
					// 			sourceRangeMin - 1,
					// 			sourceRangeMin + srcDstDiff,
					// 			interimRangeMax + srcDstDiff,
					// 		}
					// 	}
					// 	interimLocationRanges = modifyRangeSlice(
					// 		interimLocationRanges, i, newValues)
					// 	skipIndex = append(skipIndex, i+2)
					// 	i += 4
					// 	continue
					// }

					// interim min < source max < interim max
					// bottom seedrange collision
					if interimRangeMin <= sourceRangeMax &&
						sourceRangeMax < interimRangeMax {
						// newRange := []int{interimRangeMin + srcDstDiff, sourceRangeMax + srcDstDiff}
						// oldRange := []int{sourceRangeMax + 1, interimRangeMax}
						newValues := []int{
							interimRangeMin + srcDstDiff,
							sourceRangeMax + srcDstDiff,
							sourceRangeMax + 1,
							interimRangeMax}
						interimLocationRanges = modifyRangeSlice(
							interimLocationRanges, i, newValues)

						// interimLocationRanges = []int{newRange[0], newRange[1], oldRange[0], oldRange[1]}
						skipIndex = append(skipIndex, i)
						i += 2
						continue
					}

					// interim max > source min > interim min
					// top seedrange collision
					// if interimRangeMax >= sourceRangeMin &&
					// 	sourceRangeMin > interimRangeMin {
					if interimRangeMax >= sourceRangeMin &&
						sourceRangeMin > interimRangeMin {
						// newRange := []int{sourceRangeMin + srcDstDiff, interimRangeMax + srcDstDiff}
						// oldRange := []int{interimRangeMin, sourceRangeMin - 1}
						newValues := []int{
							sourceRangeMin + srcDstDiff,
							interimRangeMax + srcDstDiff,
							interimRangeMin,
							sourceRangeMin - 1}
						interimLocationRanges = modifyRangeSlice(
							interimLocationRanges, i, newValues)
						// interimLocationRanges = []int{newRange[0], newRange[1], oldRange[0], oldRange[1]}
						skipIndex = append(skipIndex, i)
						i += 2
						continue
					}
					// // saving just in case
					// if interimRangeMax >= sourceRangeMin &&
					// 	sourceRangeMin > interimRangeMin {
					// 	newRange := []int{sourceRangeMin + srcDstDiff, interimRangeMax + srcDstDiff}
					// 	oldRange := []int{interimRangeMin, sourceRangeMin - 1}
					// 	interimLocationRanges = []int{newRange[0], newRange[1], oldRange[0], oldRange[1]}
					// 	skipIndex = append(skipIndex, i)
					// 	i += 2
					// }
				}
			}
		}
		locationRanges = interimLocationRanges
		// Get lowest location number
		for i := 0; i < len(locationRanges); i += 2 {
			if lowestLocation == -1 {
				lowestLocation = locationRanges[i]
				continue
			}
			if locationRanges[i] < lowestLocation {
				lowestLocation = locationRanges[i]
				fmt.Printf("--DEBUG--   lowest: %v\n", lowestLocation)
			}

		}
	}
	// // Get lowest location number
	// for i := 0; i < len(locationRanges); i += 2 {
	// 	if lowestLocation == -1 {
	// 		lowestLocation = locationRanges[i]
	// 		continue
	// 	}
	// 	if locationRanges[i] < lowestLocation {
	// 		lowestLocation = locationRanges[i]
	// 	}
	// }
	return lowestLocation
}

func getNumbers(line string) []int {
	pattern := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(line, -1)
	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("err during getseeds: %v", err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func removeElementFromSlice(input []interface{}, startIndex, endIndex int) []interface{} {
	return append(input[:startIndex], input[endIndex+1:]...)
}

// modifyRangeSlice remove values at [index, index+1] and
// append in their position
func modifyRangeSlice(slice []int, index int, newValues []int) []int {
	if index >= 0 && index < len(slice)-1 {
		slice = append(slice[:index], slice[index+2:]...)
		slice = append(slice[:index], append(newValues, slice[index:]...)...)

	}
	return slice
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
