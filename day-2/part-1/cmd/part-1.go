package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	winningSolutions := make(map[string]string)
	winningSolutions["A"] = "Y" // Rock - Paper
	winningSolutions["B"] = "Z" // Paper - Scissors
	winningSolutions["C"] = "X" // Scissors - Rock

	drawingSolutions := make(map[string]string)
	drawingSolutions["A"] = "X" // Rock - Rock
	drawingSolutions["B"] = "Y" // Paper - Paper
	drawingSolutions["C"] = "Z" // Scissors - Scissors

	winningPoints := make(map[string]int)
	winningPoints["X"] = 1 // Rock
	winningPoints["Y"] = 2 // Paper
	winningPoints["Z"] = 3 // Scissors

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}

	dataString := string(dat)

	dataArray := strings.Split(dataString, "\n")

	totalScore := 0
	for _, lineString := range dataArray {
		roundArray := strings.Split(lineString, " ")

		if winningSolutions[roundArray[0]] == roundArray[1] {
			totalScore += 6
		} else if drawingSolutions[roundArray[0]] == roundArray[1] {
			totalScore += 3
		}

		totalScore += winningPoints[roundArray[1]]

	}

	println(totalScore)

}
