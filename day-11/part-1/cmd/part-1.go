package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Items                []int
	operationAdd         bool
	operationTimes       bool
	operationOldTimesOld bool
	operationRight       int // -1 for old
	testDivisibleBy      int
	testTrue             int
	testFalse            int
}

func main() {
	// Read File
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := strings.Split(dataString, "\r\n")

	var monkeys []Monkey
	var tempMonkey Monkey
	for line, s := range dataStringArray {
		if strings.HasPrefix(s, "Monkey ") { // New Monkey
			tempMonkey = Monkey{}
		}
		if strings.HasPrefix(s, "  Starting items: ") { // Starting items
			tmp1 := strings.Split(s, "  Starting items: ")
			tmp2 := strings.Split(tmp1[1], ", ")
			for _, s2 := range tmp2 {
				itemNum, _ := strconv.Atoi(s2)
				tempMonkey.Items = append(tempMonkey.Items, itemNum)
			}
		}
		if strings.HasPrefix(s, "  Operation: new = old ") { // Operation
			tmp1 := strings.Split(s, "  Operation: new = old ")
			if strings.HasPrefix(tmp1[1], "+") {
				tempMonkey.operationAdd = true
				tmp2 := strings.Split(tmp1[1], "+ ")
				if tmp2[1] == "old" {
					tempMonkey.operationRight = -1
				} else {
					operationRight, _ := strconv.Atoi(tmp2[1])
					tempMonkey.operationRight = operationRight
				}
			}
			if strings.HasPrefix(tmp1[1], "*") {
				tempMonkey.operationAdd = true
				tmp2 := strings.Split(tmp1[1], "* ")
				if tmp2[1] == "old" {
					tempMonkey.operationRight = -1
				} else {
					operationRight, _ := strconv.Atoi(tmp2[1])
					tempMonkey.operationRight = operationRight
				}
			}
		}

		if strings.HasPrefix(s, "  Test: divisible by ") {
			tmp1 := strings.Split(s, "  Test: divisible by ")
			divisibleBy, _ := strconv.Atoi(tmp1[1])
			tempMonkey.testDivisibleBy = divisibleBy

			tmp2 := strings.Split(dataStringArray[line+1], "    If true: throw to monkey ")
			monkeyTrue, _ := strconv.Atoi(tmp2[1])
			tempMonkey.testTrue = monkeyTrue

			tmp3 := strings.Split(dataStringArray[line+2], "    If false: throw to monkey ")
			monkeyFalse, _ := strconv.Atoi(tmp3[1])
			tempMonkey.testFalse = monkeyFalse
		}
		if len(s) == 0 || line == len(dataStringArray)-1 {
			monkeys = append(monkeys, tempMonkey)
		}
	}

	println("Done")

}
