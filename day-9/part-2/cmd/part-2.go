package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	amount    int
}

type vector struct {
	X int
	Y int
}

func main() {
	var commands []command
	uniqueMap := make(map[vector]bool)
	// Read File
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := strings.Split(dataString, "\n")
	for _, s := range dataStringArray {
		splitCommand := strings.Split(s, " ")
		amount, _ := strconv.Atoi(splitCommand[1])
		commands = append(commands, command{
			direction: splitCommand[0],
			amount:    amount,
		})
	}

	var tails []vector
	for i := 0; i < 10; i++ {
		tails = append(tails, vector{X: 0, Y: 0})
	}

	for _, command := range commands {
		for i := 0; i < command.amount; i++ {
			switch command.direction {
			case "U":
				tails[0].Y++
			case "D":
				tails[0].Y--
			case "L":
				tails[0].X--
			case "R":
				tails[0].X++
			default:
				println("ERROR")
			}

			for j := 1; j < 10; j++ {
				fixIfTailIsDiagonallyTooFarApart(tails[j-1], &tails[j])
				fixIfTailIsDirectlyTooFarApart(tails[j-1], &tails[j])
				uniqueMap[tails[9]] = true
			}
		}
	}

	println(strconv.Itoa(len(uniqueMap)))
}

func fixIfTailIsDirectlyTooFarApart(head vector, tail *vector) {
	if head.Y == tail.Y+2 && head.X == tail.X { // Up
		tail.Y++
	}
	if head.Y == tail.Y-2 && head.X == tail.X { // Down
		tail.Y--
	}

	if head.X == tail.X-2 && head.Y == tail.Y { // Left
		tail.X--
	}
	if head.X == tail.X+2 && head.Y == tail.Y { // Right
		tail.X++
	}
}

func fixIfTailIsDiagonallyTooFarApart(head vector, tail *vector) {
	if (head.Y == tail.Y+2 && head.X == tail.X-1) || (head.Y == tail.Y+1 && head.X == tail.X-2) || (head.Y == tail.Y+2 && head.X == tail.X-2) { // UpLeft
		tail.Y++
		tail.X--
	}
	if (head.Y == tail.Y+2 && head.X == tail.X+1) || (head.Y == tail.Y+1 && head.X == tail.X+2) || (head.Y == tail.Y+2 && head.X == tail.X+2) { // UpRight
		tail.Y++
		tail.X++
	}

	if (head.Y == tail.Y-2 && head.X == tail.X-1) || (head.Y == tail.Y-1 && head.X == tail.X-2) || (head.Y == tail.Y-2 && head.X == tail.X-2) { // DownLeft
		tail.Y--
		tail.X--
	}
	if (head.Y == tail.Y-2 && head.X == tail.X+1) || (head.Y == tail.Y-1 && head.X == tail.X+2) || (head.Y == tail.Y-2 && head.X == tail.X+2) { // DownRight
		tail.Y--
		tail.X++
	}
}
