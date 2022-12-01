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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	elves := make([]int64, 0)
	cur := int64(0)
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			calorie, _ := strconv.ParseInt(line, 10, 64)
			cur += calorie
		} else {
			elves = append(elves, cur)
			cur = 0
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	sum := elves[0] + elves[1] + elves[2]
	fmt.Println(sum)
}
