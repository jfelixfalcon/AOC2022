package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ROCK = 0
const PAPER = 1
const SCISSORS = 2
const OFFSET = 1
const WON = 6
const TIE = 3

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	elf := [3]string{"A", "B", "C"}
	me := [3]string{"X", "Y", "Z"}

	var elfSelect, meSelect int
	var score, scoreTwo int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for i, n := range elf {
			if n == strings.Split(line, " ")[0] {
				elfSelect = i
				break
			}
		}

		for i, n := range me {
			if n == strings.Split(line, " ")[1] {
				meSelect = i
				break
			}
		}

		score += play(meSelect, elfSelect)
		scoreTwo += playTwo(meSelect, elfSelect)

	}

	fmt.Println("Part One: " + strconv.Itoa(score))
	fmt.Println("Part Two: " + strconv.Itoa(scoreTwo))
}

func play(meSelect, elfSelect int) int {

	var temp int

	if (meSelect == ROCK && elfSelect == SCISSORS) ||
		(meSelect == PAPER && elfSelect == ROCK) ||
		(meSelect == SCISSORS && elfSelect == PAPER) {
		temp = (WON + meSelect + OFFSET)
	} else if (meSelect == ROCK && elfSelect == PAPER) ||
		(meSelect == PAPER && elfSelect == SCISSORS) ||
		(meSelect == SCISSORS && elfSelect == ROCK) {
		temp = meSelect + OFFSET
	} else {
		temp = meSelect + OFFSET + TIE
	}

	return temp

}

func playTwo(meSelect, elfSelect int) int {

	if meSelect == ROCK { // Lose
		if elfSelect == PAPER {
			meSelect = ROCK
		} else if elfSelect == ROCK {
			meSelect = SCISSORS
		} else if elfSelect == SCISSORS {
			meSelect = PAPER
		}

	} else if meSelect == PAPER { // Tie

		if elfSelect == ROCK {
			meSelect = ROCK
		} else if elfSelect == PAPER {
			meSelect = PAPER
		} else if elfSelect == SCISSORS {
			meSelect = SCISSORS
		}

	} else {

		if elfSelect == ROCK {
			meSelect = PAPER
		} else if elfSelect == PAPER {
			meSelect = SCISSORS
		} else if elfSelect == SCISSORS {
			meSelect = ROCK
		}

	}

	return play(meSelect, elfSelect)

}
