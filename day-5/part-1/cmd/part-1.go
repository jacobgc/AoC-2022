package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var arrayOfColumns [][]string

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataArray := strings.Split(dataString, "\n")

	// Figure out how many boxes there should be
	lineWithBoxCount := 0
	boxCount := 0
	for i, s := range dataArray {
		if strings.HasPrefix(s, " 1") {
			splitText := strings.Split(s, " ")
			boxCount, _ = strconv.Atoi(splitText[len(splitText)-2])
			lineWithBoxCount = i
			break
		}
	}

	// Make arrayOfColumns
	for i := lineWithBoxCount - 1; i >= 0; i-- {
		lineText := dataArray[i]
		chars := []rune(lineText)
		for j := 0; j < boxCount; j++ {
			if i == lineWithBoxCount-1 {
				arrayOfColumns = append(arrayOfColumns, []string{})
			}
			position := (j * 4) + 1
			char := string(chars[position])
			if char != " " {
				arrayOfColumns[j] = append(arrayOfColumns[j], char)
			}
		}

	}

	// Parse Instructions
	intRegex := regexp.MustCompile(`[0-9]{1,2}`)
	for i := lineWithBoxCount + 2; i < len(dataArray); i++ {
		result := intRegex.FindAllString(dataArray[i], -1)
		howManyToMove, _ := strconv.Atoi(result[0])
		from, _ := strconv.Atoi(result[1])
		to, _ := strconv.Atoi(result[2])

		for i := 0; i < howManyToMove; i++ {
			arrayOfColumns = moveBox(arrayOfColumns, from, to)
		}
	}

	for i := 0; i < len(arrayOfColumns); i++ {
		print(arrayOfColumns[i][len(arrayOfColumns[i])-1])
	}
}

func moveBox(array [][]string, from int, to int) [][]string {
	char := array[from-1][len(array[from-1])-1]
	array[from-1] = remove(array[from-1])
	array[to-1] = append(array[to-1], char)
	return array
}

// Helper method to remove last entry from slice
func remove(slice []string) []string {
	return slice[:len(slice)-1]
}
