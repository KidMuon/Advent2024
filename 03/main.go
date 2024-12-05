package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("puzzle_input_03.txt")
	if err != nil {
		fmt.Println("error importing puzzle input")
		os.Exit(1)
	}

	var candidates []mulCandidate
	for _, condChunk := range getConditionalChunks(string(fileBytes)) {
		candidateTexts := strings.Split(condChunk.text, "mul")
		for _, cText := range candidateTexts {
			newCandidate := mulCandidate{
				text:    cText,
				enabled: condChunk.enabled,
			}
			candidates = append(candidates, newCandidate)
		}
	}

	var sum, enabledSum int
	for _, candidate := range candidates {
		res := candidate.getResult()
		sum += res
		if candidate.enabled {
			enabledSum += res
		}
	}
	fmt.Printf("%d\n", sum)
	fmt.Printf("%d\n", enabledSum)
	os.Exit(0)
}

type conditionalChunk struct {
	text    string
	enabled bool
}

func getConditionalChunks(s string) []conditionalChunk {
	var conditionalChunks []conditionalChunk
	doChunks := strings.Split(s, "do()")

	for _, doChunk := range doChunks {
		containedConditionChunks := strings.SplitN(doChunk, "don't()", 2)
		if containedConditionChunks[0] != "" {
			conditionalChunks = append(conditionalChunks, conditionalChunk{
				text:    containedConditionChunks[0],
				enabled: true,
			})
		}
		if len(containedConditionChunks) > 1 {
			conditionalChunks = append(conditionalChunks, conditionalChunk{
				text:    containedConditionChunks[1],
				enabled: false,
			})
		}
	}

	return conditionalChunks
}

type mulCandidate struct {
	text         string
	firstNumber  int64
	secondNumber int64
	enabled      bool
}

func (m *mulCandidate) getResult() int {
	splitFirstParen := strings.SplitN(m.text, "(", 2)
	if splitFirstParen[0] != "" || len(splitFirstParen) == 1 {
		return 0
	}

	splitFirstComma := strings.SplitN(splitFirstParen[1], ",", 2)
	if len(splitFirstComma) == 1 {
		return 0
	}
	firstNumber, err := strconv.ParseInt(splitFirstComma[0], 10, 64)
	if err != nil {
		return 0
	}
	m.firstNumber = firstNumber

	splitSecondParen := strings.SplitN(splitFirstComma[1], ")", 2)
	if len(splitSecondParen) == 1 {
		return 0
	}
	secondNumber, err := strconv.ParseInt(splitSecondParen[0], 10, 64)
	if err != nil {
		return 0
	}
	m.secondNumber = secondNumber

	return int(m.firstNumber) * int(m.secondNumber)
}
