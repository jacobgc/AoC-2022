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

	losingSolutions := make(map[string]string)
	losingSolutions["A"] = "Z" // Rock - Scissors
	losingSolutions["B"] = "X" // Paper - Rock
	losingSolutions["C"] = "Y" // Scissors - Rock

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

		switch roundArray[1] {
		case "X":
			print("Converting to win: ")
			roundArray[1] = losingSolutions[roundArray[0]]
			break
		case "Y":
			print("Converting to draw: ")
			roundArray[1] = drawingSolutions[roundArray[0]]
			break
		case "Z":
			print("Converting to loss: ")
			roundArray[1] = winningSolutions[roundArray[0]]
			break

		}

		if winningSolutions[roundArray[0]] == roundArray[1] {
			println("win")
			totalScore += 6
		} else if drawingSolutions[roundArray[0]] == roundArray[1] {
			println("draw")
			totalScore += 3
		} else {
			println("loss")
		}

		totalScore += winningPoints[roundArray[1]]

	}

	println(totalScore)

}
