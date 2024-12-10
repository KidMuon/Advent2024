package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("puzzle_input_04.txt")
	if err != nil {
		fmt.Println("Error importing puzzle input")
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n")

	wsm := convertLinesToWordSearchMatrix(lines)

	totalCount := countInWordSearch(wsm, "MAS")
	masCrossed := countCrossesInWordSearch(wsm, "MAS")

	fmt.Printf("%d\n", totalCount)
	fmt.Printf("%d\n", masCrossed)
}

func convertLinesToWordSearchMatrix(lines []string) wordSearchMatrix {
	var wsm wordSearchMatrix
	for _, line := range lines {
		lineRunes := []rune{}
		for _, r := range line {
			lineRunes = append(lineRunes, r)
		}
		wsm = append(wsm, lineRunes)
	}
	return wsm
}

func findInWordSearch(wsm wordSearchMatrix, answer string) []wordSearchVector {
	foundPositions := []wordSearchVector{}

	var first_letter rune
	for _, r := range answer {
		first_letter = r
		break
	}

	for y := 0; y < len(wsm); y++ {
		for x := 0; x < len(wsm[y]); x++ {
			if wsm[y][x] == first_letter {
				for _, direction := range getAllVectorDirections() {
					wsv := wordSearchVector{
						position: wordSearchPosition{
							x: x,
							y: y,
						},
						direction: direction,
					}
					if answerPresent(wsm, &wsv, 0, answer) {
						foundPositions = append(foundPositions, wordSearchVector{
							position: wordSearchPosition{
								x: x,
								y: y,
							},
							direction: direction,
						})
					}
				}
			}
		}
	}

	return foundPositions
}

func countInWordSearch(wsm wordSearchMatrix, answer string) int {
	return len(findInWordSearch(wsm, answer))
}

func findCrossesInWordSearch(wsm wordSearchMatrix, answer string) map[wordSearchPosition]int {
	crossLocations := make(map[wordSearchPosition]int)

	var string_length int
	for range answer {
		string_length++
	}
	if string_length%2 == 0 {
		return crossLocations
	}

	positions := findInWordSearch(wsm, answer)
	positions = progressAllVectors(positions, string_length/2)

	for _, wsvector := range positions {
		if wsvector.direction == north ||
			wsvector.direction == east ||
			wsvector.direction == west ||
			wsvector.direction == south {
			continue
		}
		if _, ok := crossLocations[wsvector.position]; !ok {
			crossLocations[wsvector.position] = 1
		} else {
			crossLocations[wsvector.position]++
		}
	}

	for k, v := range crossLocations {
		if v == 1 {
			delete(crossLocations, k)
		}
	}

	return crossLocations
}

func progressAllVectors(wsvectors []wordSearchVector, n int) []wordSearchVector {
	progressedVectors := []wordSearchVector{}
	for _, wsvector := range wsvectors {
		for i := 0; i < n; i++ {
			wsvector.progressAlong()
			progressedVectors = append(progressedVectors, wsvector)
		}
	}
	return progressedVectors
}

func countCrossesInWordSearch(wsm wordSearchMatrix, answer string) int {
	return len(findCrossesInWordSearch(wsm, answer))
}

func answerPresent(wsm wordSearchMatrix, wsv *wordSearchVector, expected_idx int, answer string) bool {
	if expected_idx == len(answer) {
		return true
	}

	var expected_rune rune
	for i, r := range answer {
		if i == expected_idx {
			expected_rune = r
			break
		}
	}

	if wsm.searchForExpected(wsv, expected_rune) {
		wsv.progressAlong()
		return answerPresent(wsm, wsv, expected_idx+1, answer)
	}

	return false
}

type vectorDirection int

const (
	west vectorDirection = iota
	northwest
	north
	northeast
	east
	southeast
	south
	southwest
)

func getAllVectorDirections() []vectorDirection {
	return []vectorDirection{
		west, northwest, north, northeast,
		east, southeast, south, southwest,
	}
}

type wordSearchPosition struct {
	x int
	y int
}

type wordSearchVector struct {
	position  wordSearchPosition
	direction vectorDirection
}

func (mv *wordSearchVector) progressAlong() {
	var new_x, new_y int
	switch mv.direction {
	case west:
		new_x = mv.position.x - 1
		new_y = mv.position.y
	case northwest:
		new_x = mv.position.x - 1
		new_y = mv.position.y - 1
	case north:
		new_x = mv.position.x
		new_y = mv.position.y - 1
	case northeast:
		new_x = mv.position.x + 1
		new_y = mv.position.y - 1
	case east:
		new_x = mv.position.x + 1
		new_y = mv.position.y
	case southeast:
		new_x = mv.position.x + 1
		new_y = mv.position.y + 1
	case south:
		new_x = mv.position.x
		new_y = mv.position.y + 1
	case southwest:
		new_x = mv.position.x - 1
		new_y = mv.position.y + 1
	}
	mv.position.x, mv.position.y = new_x, new_y
}

type wordSearchMatrix [][]rune

func (ws wordSearchMatrix) searchForExpected(wsv *wordSearchVector, expected rune) bool {
	x_max := len(ws[0])
	y_max := len(ws)
	if wsv.position.x < 0 || wsv.position.x >= x_max {
		return false
	}
	if wsv.position.y < 0 || wsv.position.y >= y_max {
		return false
	}
	return ws[wsv.position.y][wsv.position.x] == expected
}
