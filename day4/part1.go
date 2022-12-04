package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	utils.GetInput(4)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, ",")
		first := val[0]
		second := val[1]
		firstVal := strings.Split(first, "-")
		secondVal := strings.Split(second, "-")
		firstStart, _ := strconv.ParseInt(firstVal[0], 10, 64)
		firstEnd, _ := strconv.ParseInt(firstVal[1], 10, 64)
		secondStart, _ := strconv.ParseInt(secondVal[0], 10, 64)
		secondEnd, _ := strconv.ParseInt(secondVal[1], 10, 64)
		if (firstStart <= secondEnd && firstEnd >= secondStart) ||
			(secondStart <= firstEnd && secondEnd >= firstStart) {
			count++
		}
	}

	fmt.Println(count)
}
