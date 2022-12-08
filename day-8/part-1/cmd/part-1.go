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
	visibleTrees := make(map[int]map[int]bool)

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := strings.Split(dataString, "\r\n")

	for i, s := range dataStringArray {
		rowItems := strings.Split(s, "")
		trees = append(trees, []int{})
		visibleTrees[i] = make(map[int]bool)
		for _, item := range rowItems {
			itemInt, _ := strconv.Atoi(item)
			trees[i] = append(trees[i], itemInt)
		}
	}

	totalScore := 0

	for row, treeRow := range trees {
		for col := range treeRow {
			if row == 0 || row == len(trees)-1 {
				visibleTrees[row][col] = true
				totalScore++
				continue
			}
			if col == 0 || col == len(treeRow)-1 {
				visibleTrees[row][col] = true
				totalScore++
				continue
			}
			visibleTrees[row][col] = isTreeVisible(trees, row, col)
			if visibleTrees[row][col] == true {
				totalScore++
			}
		}
	}

	println(totalScore)
}

func isTreeVisible(trees [][]int, row, col int) bool {

	isHidden := make(map[string]bool)
	isHidden["up"] = false
	isHidden["down"] = false
	isHidden["left"] = false
	isHidden["right"] = false

	treeInQuestion := trees[row][col]

	loopUpAmount := 0
	loopDownAmount := len(trees) - 1
	loopLeftAmount := 0
	loopRightAmount := len(trees[0]) - 1

	// Left - Broken
	for i := col - 1; i >= loopLeftAmount; i-- {
		comparisonTree := trees[row][i]
		if treeInQuestion <= comparisonTree {
			isHidden["left"] = true
		}
	}

	// Right
	for i := col + 1; i <= loopRightAmount; i++ {
		comparisonTree := trees[row][i]
		if treeInQuestion <= comparisonTree {
			isHidden["right"] = true
		}
	}

	// Up - Broken
	for i := row - 1; i >= loopUpAmount; i-- {
		comparisonTree := trees[i][col]
		if treeInQuestion <= comparisonTree {
			isHidden["up"] = true
		}
	}

	// Down
	for i := row + 1; i <= loopDownAmount; i++ {
		comparisonTree := trees[i][col]
		if treeInQuestion <= comparisonTree {
			isHidden["down"] = true
		}
	}

	if isHidden["up"] == false || isHidden["down"] == false || isHidden["left"] == false || isHidden["right"] == false {
		return true
	} else {
		return false
	}

}
