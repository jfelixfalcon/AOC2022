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

	scanner := bufio.NewScanner(file)
	var count, countTwo int

	for scanner.Scan() {

		pairs := scanner.Text()
		temp := strings.Split(pairs, ",")
		elfOne := strings.Split(temp[0], "-")
		elfTwo := strings.Split(temp[1], "-")

		a, _ := strconv.Atoi(elfOne[0])
		b, _ := strconv.Atoi(elfOne[1])
		c, _ := strconv.Atoi(elfTwo[0])
		d, _ := strconv.Atoi(elfTwo[1])

		if (a <= c && b >= d) || (a >= c && b <= d) {
			count++
		}

		if b >= c && d >= a {
			countTwo++
		}

	}

	fmt.Println("Total: " + strconv.Itoa(count))
	fmt.Println("Total Two: " + strconv.Itoa(countTwo))
}
