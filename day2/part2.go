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
		"A": 1,
		"B": 2,
		"C": 3,
	}
	outcome := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	rules := make(map[string]map[string]string)
	rules["A"] = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	rules["B"] = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	rules["C"] = map[string]string{
		"X": "B",
		"Y": "C",
		"Z": "A",
	}
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, " ")
		opp := val[0]
		me := val[1]
		score += shape[rules[opp][me]]
		score += outcome[me]
	}

	fmt.Println(score)
}
