package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("puzzle_input_01.txt")
	if err != nil {
		fmt.Println("Error importing puzzle input")
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n")

	var leftList, rightList []int

	for _, line := range lines {
		leftValue, rightValue, err := getValuesFromLine(line)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	fmt.Printf("%d\n", listDistance(leftList, rightList))
	fmt.Printf("%d\n", listSimilarityScore(leftList, rightList))
	os.Exit(0)
}

func getValuesFromLine(line string) (leftValue, rightValue int, err error) {
	values := strings.Fields(line)
	left, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Problem Parsing Value")
	}
	right, err := strconv.ParseInt(values[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Problem Parsing Value")
	}
	leftValue = int(left)
	rightValue = int(right)
	return leftValue, rightValue, nil
}

type byValue []int

func (a byValue) Len() int {
	return len(a)
}

func (a byValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byValue) Less(i, j int) bool {
	return a[i] < a[j]
}

func listDistance(leftList, rightList []int) int {
	sort.Sort(byValue(leftList))
	sort.Sort(byValue(rightList))

	var sum int

	for i := 0; i < len(leftList); i++ {
		difference := leftList[i] - rightList[i]
		sum += int(math.Abs(float64(difference)))
	}

	return sum
}

func listSimilarityScore(leftList, rightList []int) int {
	score := 0

	commonValues := make(map[int]int)
	for _, val := range leftList {
		commonValues[val] = 0
	}

	for _, val := range rightList {
		if _, ok := commonValues[val]; !ok {
			continue
		}
		commonValues[val] += 1
	}

	for _, val := range leftList {
		score += val * commonValues[val]
	}

	return score
}
