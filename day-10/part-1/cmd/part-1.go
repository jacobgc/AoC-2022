package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	command string
	value   int
}

func main() {
	// Read File
	var commands []command
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := strings.Split(dataString, "\n")
	for _, s := range dataStringArray {
		c := strings.Split(s, " ")
		value := 0
		if len(c) > 1 {
			value, _ = strconv.Atoi(c[1])
		}

		commands = append(commands, command{
			command: c[0],
			value:   value,
		})
	}

	x := 1
	blockedFor := 0
	currentCycle := 1
	cyclesSinceLastPrint := 0
	valueToAddWhenUnblocked := 0
	firstCycle := true
	totalScore := 0

	lenCommandsStart := len(commands)

	for i := 0; i <= lenCommandsStart; {
		if blockedFor == 0 {
			if !firstCycle {
				i++
			}
			firstCycle = false
			// Add any values we need when unblocked
			if valueToAddWhenUnblocked != 0 {
				x += valueToAddWhenUnblocked
				valueToAddWhenUnblocked = 0
			}

			if i >= lenCommandsStart {
				continue // Don't try to get any new commands
			}

			// Grab and pop the next command
			command := commands[0]
			commands = commands[1:]

			if command.command == "noop" {
				// Do nothing
			} else if command.command == "addx" {
				valueToAddWhenUnblocked = command.value
				blockedFor = 1
			}
		} else {
			blockedFor--
		}

		if currentCycle == 20 || cyclesSinceLastPrint == 40 {
			cyclesSinceLastPrint = 1
			totalScore += currentCycle * x
		} else {
			cyclesSinceLastPrint++
		}

		currentCycle++
	}
	println(totalScore)
	println("Done")
}
