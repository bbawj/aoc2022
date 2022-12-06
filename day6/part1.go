package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	utils.GetInput(6)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		window := make(map[byte]int)
		left := 0
		right := 0
		for right-left < 14 && right < len(line) {
			if pos, found := window[line[right]]; !found {
				window[line[right]] = right
			} else {
				for left <= pos {
					delete(window, line[left])
					left++
				}
				window[line[right]] = right
			}
			right++
		}
		fmt.Println(right)
	}
}
