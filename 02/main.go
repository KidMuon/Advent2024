package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("puzzle_input_02.txt")
	if err != nil {
		fmt.Println("error importing puzzle input")
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n")
	records, err := convertLinesToRecords(lines)
	if err != nil {
		fmt.Printf("error converting lines to records: %v", err)
		os.Exit(1)
	}

	var safeRecords, singleErrorRecords int

	for _, r := range records {
		if r.graduallyChanging() {
			safeRecords++
		}
		if r.graduallyChangingProblemDampened() {
			singleErrorRecords++
		}
	}

	fmt.Printf("%d\n", safeRecords)
	fmt.Printf("%d\n", singleErrorRecords)
}

type record []int64

func convertLinesToRecords(lines []string) ([]record, error) {
	result := []record{}
	for _, line := range lines {
		newRecord := record{}
		for _, valueString := range strings.Fields(line) {
			value, err := strconv.ParseInt(valueString, 10, 64)
			if err != nil {
				return []record{}, fmt.Errorf("line cannot be parsed to integers: %s", line)
			}
			newRecord = append(newRecord, value)
		}
		result = append(result, newRecord)
	}
	return result, nil
}

func (r record) graduallyChangingProblemDampened() bool {
	if r.graduallyChanging() {
		return true
	}

	for i := 0; i < len(r); i++ {
		if r.testAround(i) {
			return true
		}
	}

	return false
}

func (r record) testAround(i int) bool {
	r_ := record{}
	r_2 := record{}
	for rj := 0; rj < len(r); rj++ {
		if rj != i-1 {
			r_ = append(r_, r[rj])
		}
		if rj != i {
			r_2 = append(r_2, r[rj])
		}
	}
	return r_.graduallyChanging() || r_2.graduallyChanging()
}

func (r record) graduallyChanging() bool {
	direction := r[0] < r[1]
	for i := 1; i < len(r); i++ {
		currentDirection := r[i-1] < r[i]
		if direction != currentDirection {
			return false
		}
		if r[i-1] == r[i] {
			return false
		}
		if r[i-1]-r[i] > 3 || r[i]-r[i-1] > 3 {
			return false
		}
	}
	return true
}
