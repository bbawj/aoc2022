package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	x int
	y int
}

type Pos struct {
	x int
	y int
}

func main() {
	utils.GetInput(9)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	head := Pos{0, 0}
	tail := Pos{0, 0}
	dirs := map[string]Dir{
		"L": {-1, 0},
		"R": {1, 0},
		"U": {0, 1},
		"D": {0, -1},
	}
	opposites := map[string]string{
		"L": "R",
		"R": "L",
		"U": "D",
		"D": "U",
	}

	visited := make(map[Pos]int)
	visited[tail] = 1
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		val := strings.Split(line, " ")
		dir := dirs[val[0]]
		count, _ := strconv.ParseInt(val[1], 10, 64)
		for i := int64(0); i < count; i++ {
			head.x += dir.x
			head.y += dir.y
			dist := math.Sqrt(math.Pow(float64(head.x-tail.x), 2) + math.Pow(float64(head.y-tail.y), 2))
			if dist >= 2 {
				opp := dirs[opposites[val[0]]]
				tail.x = head.x + opp.x
				tail.y = head.y + opp.y
				if _, found := visited[tail]; !found {
					visited[tail] = 1
				} else {
					visited[tail]++
				}
			}
		}
	}
	fmt.Println(len(visited))
}
