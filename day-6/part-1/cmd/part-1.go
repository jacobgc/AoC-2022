package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

const UniqueByteCount = 4

func main() {
	start := time.Now()
	// Read File
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := []rune(dataString)

	// Go through every character
	for i := 0; i < len(dataStringArray); i++ {
		theMap := make(map[rune]bool)
		// Add required characters to a map
		for j := i; j < i+UniqueByteCount; j++ {
			theMap[dataStringArray[j]] = true
		}

		// Check if the map length is large enough for them all to be unique
		if len(theMap) == UniqueByteCount {
			println("Position Found at index: " + strconv.Itoa(i+UniqueByteCount))
			break
		}
	}
	elapsed := time.Since(start)
	log.Printf("Processing took: %s", elapsed)
}
