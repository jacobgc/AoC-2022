package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func contains(s []int, r int) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func main() {

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	pairs := strings.Split(dataString, "\n")

	totalScore := 0

	for _, pair := range pairs {
		var array1 []int
		var array2 []int
		pairArray := strings.Split(pair, ",")

		splitSingle1 := strings.Split(pairArray[0], "-")
		num11, _ := strconv.Atoi(splitSingle1[0])
		num12, _ := strconv.Atoi(splitSingle1[1])

		splitSingle2 := strings.Split(pairArray[1], "-")
		num21, _ := strconv.Atoi(splitSingle2[0])
		num22, _ := strconv.Atoi(splitSingle2[1])

		for i := num11; i <= num12; i++ {
			array1 = append(array1, i)
		}

		for i := num21; i <= num22; i++ {
			array2 = append(array2, i)
		}

		found := true
		for _, entry := range array1 {
			if !contains(array2, entry) {
				found = false
			}
		}

		if found == true {
			totalScore++
			continue // Dont check the other way around
		}

		found = true
		for _, entry := range array2 {
			if !contains(array1, entry) {
				found = false
			}
		}

		if found == true {
			totalScore++
		}

	}

	println(totalScore)

}
