package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stacks [][]string
	stacks = append(stacks, []string{"S", "C", "V", "N"})
	stacks = append(stacks, []string{"Z", "M", "J", "H", "N", "S"})
	stacks = append(stacks, []string{"M", "C", "T", "G", "J", "N", "D"})
	stacks = append(stacks, []string{"T", "D", "F", "J", "W", "R", "M"})
	stacks = append(stacks, []string{"P", "F", "H"})
	stacks = append(stacks, []string{"C", "T", "Z", "H", "J"})
	stacks = append(stacks, []string{"D", "P", "R", "Q", "F", "S", "L", "Z"})
	stacks = append(stacks, []string{"C", "S", "L", "H", "D", "F", "P", "W"})
	stacks = append(stacks, []string{"D", "S", "M", "P", "F", "N", "G", "Z"})

	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		count, _ := strconv.Atoi(temp[1])
		from, _ := strconv.Atoi(temp[3])
		to, _ := strconv.Atoi(temp[5])

		//partOne(count, from, to, stacks)
		partTwo(count, from, to, stacks)

	}

	temp := ""
	for i := range stacks {
		temp += stacks[i][len(stacks[i])-1]
	}

	fmt.Println("Results: " + temp)

}

func partTwo(count, from, to int, stacks [][]string) {

	from--
	to--
	fromLen := len(stacks[from])

	temp := stacks[from][fromLen-count:]
	stacks[to] = append(stacks[to], temp...)
	stacks[from] = stacks[from][:fromLen-count]

}

func partOne(count, from, to int, stacks [][]string) {

	from--
	to--
	fromLen := len(stacks[from])

	for n := 0; n < count; n++ {
		stacks[to] = append(stacks[to], stacks[from][fromLen-1])
		stacks[from] = stacks[from][:fromLen-1]
	}
}
