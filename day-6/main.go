package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const MARKERLEN = 14 // 4 for part one

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		start := 0
		end := 14 // 4 for part one
		marker := line[start:end]

		for {

			for n := 0; n < MARKERLEN; n++ {

				if strings.Count(marker, string(marker[n])) > 1 {
					start++
					end++
					marker = line[start:end]
					break
				}

				if n == MARKERLEN-1 {
					fmt.Println(end)
					os.Exit(0)
				}
			}
		}
	}
}
