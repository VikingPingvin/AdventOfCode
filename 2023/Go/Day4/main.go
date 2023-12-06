package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

const (
	inputFilePath = "input.txt"
)

type Card struct {
	// cardNumber     int
	winningNumbers []int
	handNumbers    []int
}

type CardCollection struct {
	cards map[int]Card
}

func main() {
	data, err := readFile(inputFilePath)
	if err != nil {
		fmt.Printf("err %v", err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")

	cardCollection := &CardCollection{}
	// cardCollection.cards = make(map[int]Card)
	cardCollection.cards = map[int]Card{
		1: {},
	}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		cardNum, winningNumbers, handNumbers := splitLine(line)
		cardCollection.cards[cardNum] = Card{winningNumbers, handNumbers}
	}
	// copyOfMapP1 := new(CardCollection)
	// maps.Copy[]()
	// copyOfMapP1.cards = cardCollection.cards //KURVAANYÁD GO
	// copyOfMapP2 := *cardCollection

	p1SumTotal := solveP1(cardCollection)
	fmt.Printf("P1 Result: %d\n", p1SumTotal)

	// I suck at copying structs/maps in Go, uncomment p1 solver
	// so cardCollection is not modified/reduced during p1solver
	p2SumTotal := solveP2(cardCollection)
	fmt.Printf("P2 Result: %d\n", p2SumTotal)

}

func solveP1(cards *CardCollection) int {
	sumTotal := 0
	for _, card := range cards.cards {
		sumForCard := calculatePointsForCard(&card, p1PointCalculation)
		sumTotal += sumForCard
	}

	return sumTotal
}

func solveP2(cards *CardCollection) int {
	sumOfScratchCards := 0
	var copiesMap = make(map[int]int)

	// We need to sort here...
	sortedKeys := maps.Keys(cards.cards)
	slices.Sort(sortedKeys)

	for _, key := range sortedKeys {
		copiesMap[key]++
	}

	for _, key := range sortedKeys {
		copies := copiesMap[key]

		card := cards.cards[key]
		numOfMatches := calculatePointsForCard(&card, p2PointCalculation)
		for i := 0; i < copies; i++ {
			for i := key + 1; i <= key+numOfMatches; i++ {
				copiesMap[i]++
			}
		}
	}

	for _, value := range copiesMap {
		sumOfScratchCards += value
	}

	return sumOfScratchCards
}

type PointCalculationFunction func(int) int

func p1PointCalculation(input int) int {
	if input == 0 {
		input += 1
	} else {
		input *= 2
	}
	return input
}

func p2PointCalculation(input int) int {
	return input + 1
}

func calculatePointsForCard(card *Card, pointCalculationFunc PointCalculationFunction) int {
	sumForCard := 0
	for _, currentWinningNumber := range card.winningNumbers {
		for {
			// fmt.Printf("Checking winning number: %d\n", currentWinningNumber)
			index := slices.Index(card.handNumbers, currentWinningNumber)
			if index == -1 {
				break
			}

			sumForCard = pointCalculationFunc(sumForCard)

			// Remove found element from slice
			card.handNumbers = append(card.handNumbers[:index], card.handNumbers[index+1:]...)
		}
	}
	return sumForCard
}

func (cc *CardCollection) Copy() CardCollection {
	cards := cc.cards

	copy := &CardCollection{
		cards: cards,
	}

	return *copy
}

func splitLine(input string) (cardNum int, winningNumbers []int, handNumbers []int) {
	if len(input) <= 0 {
		return 0, nil, nil
	}
	colonSplit := strings.Split(input, ":")
	cardString := colonSplit[0]

	barSplit := strings.Split(string(colonSplit[1]), "|")
	winningNumbersString := barSplit[0]
	handNumbersString := barSplit[1]

	reNumbers := regexp.MustCompile(`(\d+)`)
	cardNumReMatches := reNumbers.FindAllString(cardString, -1)
	cardNum, err := strconv.Atoi(cardNumReMatches[0])
	if err != nil {
		fmt.Printf("error converting cardnumber %v", err)
		os.Exit(1)
	}

	winningNumberReMatches := reNumbers.FindAllString(winningNumbersString, -1)
	winningNumbers, err = stringSliceToIntSlice(winningNumberReMatches)
	if err != nil {
		fmt.Printf("error converting winningnumber %v", err)
		os.Exit(1)
	}
	handNumberReMatches := reNumbers.FindAllString(handNumbersString, -1)
	handNumbers, err = stringSliceToIntSlice(handNumberReMatches)
	if err != nil {
		fmt.Printf("error converting handnumber %v", err)
		os.Exit(1)
	}

	return cardNum, winningNumbers, handNumbers
}

func stringSliceToIntSlice(input []string) ([]int, error) {
	var intSlice = make([]int, len(input))

	for i, str := range input {
		converted, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intSlice[i] = converted
	}

	return intSlice, nil
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
