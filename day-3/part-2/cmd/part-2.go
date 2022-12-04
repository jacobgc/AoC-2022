package main

import (
	"log"
	"os"
	"strings"
)

func contains(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func containsMultiple(arrayOfRunes [][]rune, runeToSearchFor rune) bool {
	for _, runes := range arrayOfRunes {
		if !contains(runes, runeToSearchFor) {
			return false
		}
	}
	return true
}

func main() {

	scores := map[string]int{}
	scores["a"] = 1
	scores["b"] = 2
	scores["c"] = 3
	scores["d"] = 4
	scores["e"] = 5
	scores["f"] = 6
	scores["g"] = 7
	scores["h"] = 8
	scores["i"] = 9
	scores["j"] = 10
	scores["k"] = 11
	scores["l"] = 12
	scores["m"] = 13
	scores["n"] = 14
	scores["o"] = 15
	scores["p"] = 16
	scores["q"] = 17
	scores["r"] = 18
	scores["s"] = 19
	scores["t"] = 20
	scores["u"] = 21
	scores["v"] = 22
	scores["w"] = 23
	scores["x"] = 24
	scores["y"] = 25
	scores["z"] = 26
	scores["A"] = 27
	scores["B"] = 28
	scores["C"] = 29
	scores["D"] = 30
	scores["E"] = 31
	scores["F"] = 32
	scores["G"] = 33
	scores["H"] = 34
	scores["I"] = 35
	scores["J"] = 36
	scores["K"] = 37
	scores["L"] = 38
	scores["M"] = 39
	scores["N"] = 40
	scores["O"] = 41
	scores["P"] = 42
	scores["Q"] = 43
	scores["R"] = 44
	scores["S"] = 45
	scores["T"] = 46
	scores["U"] = 47
	scores["V"] = 48
	scores["W"] = 49
	scores["X"] = 50
	scores["Y"] = 51
	scores["Z"] = 52

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataStrig := string(dat)
	backpacks := strings.Split(dataStrig, "\n")
	totalScore := 0

	for i := 0; i < len(backpacks); i += 3 {
		var arrayOfRunes [][]rune
		arrayOfRunes = append(arrayOfRunes, []rune(backpacks[i+1]))
		arrayOfRunes = append(arrayOfRunes, []rune(backpacks[i+2]))
		firstCompartment := []rune(backpacks[i])

		println("First three lines: " + string(firstCompartment[0]) + string(arrayOfRunes[0][0]) + string(arrayOfRunes[1][0]))

		var usedRunes []rune
		for _, r := range firstCompartment {
			doesContain := containsMultiple(arrayOfRunes, r)
			runeUsed := contains(usedRunes, r)
			if doesContain && !runeUsed {
				println("All Contain: " + string(r))
				totalScore += scores[string(r)]
				usedRunes = append(usedRunes, r)
			}
		}

	}

	println(totalScore)

	//for _, backpack := range backpacks {
	//	chars := []rune(backpack)
	//	var firstCompartment []rune
	//	var secondCompartment []rune
	//	for i := 0; i < len(chars)/2; i++ {
	//		firstCompartment = append(firstCompartment, chars[i])
	//	}
	//	for i := len(chars) / 2; i < len(chars); i++ {
	//		secondCompartment = append(secondCompartment, chars[i])
	//	}
	//
	//	var usedRunes []rune
	//	for _, r := range firstCompartment {
	//		doesContain := contains(secondCompartment, r)
	//		runeUsed := contains(usedRunes, r)
	//		if doesContain && !runeUsed {
	//			println("Contains: " + string(r))
	//			totalScore += scores[string(r)]
	//			usedRunes = append(usedRunes, r)
	//		}
	//	}
	//
	//}
	//
	//println(totalScore)

}
