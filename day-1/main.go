package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var arr []int
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {

		x := scanner.Text()

		if x == "" {

			arr = append(arr, sum)
			sum = 0

		}

		n, _ := strconv.Atoi(x)
		sum += n

	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr[0])
	fmt.Print(arr[0] + arr[1] + arr[2])

}
