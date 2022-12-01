package main

import (
	"github.com/jacobgc/AoC-2022/day-1/part-2/internal"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var elves []*internal.Elf

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}

	dataString := string(dat)
	dataArray := regexp.MustCompile(`\n\s*\n`).Split(dataString, -1)

	for _, elfArray := range dataArray {
		newElf := internal.NewElf()
		arrayOfCalories := strings.Split(elfArray, "\n")
		for _, calorie := range arrayOfCalories {
			if len(calorie) == 0 {
				continue
			}
			calorieInt, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatalln("Failed to convert string to int:" + err.Error())
			}
			newElf.AddCalories(calorieInt)
		}

		elves = append(elves, newElf)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].CarriedCalories > elves[j].CarriedCalories
	})

	println(elves[0].CarriedCalories)
	println(elves[1].CarriedCalories)
	println(elves[2].CarriedCalories)

	println(elves[0].CarriedCalories + elves[1].CarriedCalories + elves[2].CarriedCalories)
}
