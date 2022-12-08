package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read File

	var trees [][]int
	treeScores := make(map[int]map[int]int)

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := strings.Split(dataString, "\r\n")

	for i, s := range dataStringArray {
		rowItems := strings.Split(s, "")
		trees = append(trees, []int{})
		treeScores[i] = make(map[int]int)
		for _, item := range rowItems {
			itemInt, _ := strconv.Atoi(item)
			trees[i] = append(trees[i], itemInt)
		}
	}

	totalScore := 0

	for row, treeRow := range trees {
		for col := range treeRow {
			if row == 0 || row == len(trees)-1 {
				continue
			}
			if col == 0 || col == len(treeRow)-1 {
				continue
			}

			ans := calculateTreeScore(trees, row, col)
			if ans > totalScore {
				totalScore = ans
			}

		}
	}

	println(totalScore)
}

func calculateTreeScore(trees [][]int, row, col int) int {

	canSeeUp := 0
	canSeeDown := 0
	canSeeLeft := 0
	canSeeRight := 0

	treeInQuestion := trees[row][col]

	loopUpAmount := 0
	loopDownAmount := len(trees) - 1
	loopLeftAmount := 0
	loopRightAmount := len(trees[0]) - 1

	// Left - Broken
	for i := col - 1; i >= loopLeftAmount; i-- {
		comparisonTree := trees[row][i]
		canSeeLeft++
		if treeInQuestion <= comparisonTree {
			break
		}
	}

	// Right
	for i := col + 1; i <= loopRightAmount; i++ {
		comparisonTree := trees[row][i]
		canSeeRight++
		if treeInQuestion <= comparisonTree {
			break
		}
	}

	// Up - Broken
	for i := row - 1; i >= loopUpAmount; i-- {
		comparisonTree := trees[i][col]
		canSeeUp++
		if treeInQuestion <= comparisonTree {
			break
		}
	}

	// Down
	for i := row + 1; i <= loopDownAmount; i++ {
		comparisonTree := trees[i][col]
		canSeeDown++
		if treeInQuestion <= comparisonTree {
			break
		}
	}
	totalScore := canSeeLeft * canSeeRight * canSeeUp * canSeeDown
	return totalScore

}
