package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	shape := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	rules := make(map[string]map[string]int)
	rules["A"] = map[string]int{
		"X": 3,
		"Y": 6,
		"Z": 0,
	}
	rules["B"] = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	rules["C"] = map[string]int{
		"X": 6,
		"Y": 0,
		"Z": 3,
	}
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, " ")
		opp := val[0]
		me := val[1]
		score += rules[opp][me]
		score += shape[me]
	}

	fmt.Println(score)
}
