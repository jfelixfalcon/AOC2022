package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const UPPEROFFSET = 38
const LOWEROFFSET = 96

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total, totalTwo int
	count := 1
	var temp []string

	for scanner.Scan() {
		rucksack := scanner.Text()
		temp = append(temp, rucksack)

		if count%3 == 0 && count >= 3 {
			totalTwo += partTwo(temp)
			temp = temp[:0]
		}

		total += partOne(rucksack)
		count++
	}

	fmt.Println("Part One: " + strconv.Itoa(total))
	fmt.Println("Part One: " + strconv.Itoa(totalTwo))

}

func partTwo(rucksacks []string) int {
	for _, letter := range rucksacks[0] {
		if strings.Contains(rucksacks[1], string(letter)) &&
			strings.Contains(rucksacks[2], string(letter)) {

			if unicode.IsUpper(letter) {
				return int(letter) - UPPEROFFSET

			} else if unicode.IsLower(letter) {
				return int(letter) - LOWEROFFSET
			}

		}
	}
	return 0
}

func partOne(rucksack string) int {
	strlen := len(rucksack) / 2
	compartmentOne := rucksack[:strlen]
	compartmentTwo := rucksack[strlen:]

	for _, letter := range compartmentOne {

		if strings.Contains(compartmentTwo, string(letter)) {
			if unicode.IsUpper(letter) {
				return int(letter) - UPPEROFFSET

			} else if unicode.IsLower(letter) {
				return int(letter) - LOWEROFFSET
			}
		}
	}

	return 0

}
