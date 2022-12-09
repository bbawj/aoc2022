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
	// tail := Pos{0, 0}
	tails := make([]*Pos, 9)
	for i := range tails {
		tails[i] = &Pos{0, 0}
	}
	dirs := map[string]Dir{
		"L":  {-1, 0},
		"R":  {1, 0},
		"U":  {0, 1},
		"D":  {0, -1},
		"LU": {-1, 1},
		"LD": {-1, -1},
		"RU": {1, 1},
		"RD": {1, -1},
	}
	// opposites := map[string]string{
	// 	"L": "R",
	// 	"R": "L",
	// 	"U": "D",
	// 	"D": "U",
	// }

	visited := make(map[Pos]int)
	visited[*tails[0]] = 1
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		val := strings.Split(line, " ")
		dir := dirs[val[0]]
		count, _ := strconv.ParseInt(val[1], 10, 64)
		var opp Dir
		for i := int64(0); i < count; i++ {
			head.x += dir.x
			head.y += dir.y
			temp := &head
			for j := 0; j < 9; j++ {
				tail := tails[j]
				if temp.x-tail.x == 2 {
					opp = dirs["L"]
					if temp.y-tail.y == 2 {
						opp = dirs["LD"]
					} else if temp.y-tail.y == -2 {
						opp = dirs["LU"]
					}
				} else if temp.x-tail.x == -2 {
					opp = dirs["R"]
					if temp.y-tail.y == 2 {
						opp = dirs["RD"]
					} else if temp.y-tail.y == -2 {
						opp = dirs["RU"]
					}
				} else if temp.y-tail.y == 2 {
					opp = dirs["D"]
				} else if temp.y-tail.y == -2 {
					opp = dirs["U"]
				} else {
					opp = Dir{0, 0}
					break
				}
				// dist := math.Sqrt(math.Pow(float64(head.x-tail.x), 2) + math.Pow(float64(head.y-tail.y), 2))
				tail.x = temp.x + opp.x
				tail.y = temp.y + opp.y
				if _, found := visited[*tail]; j == 8 && !found {
					visited[*tail] = 1
				}
				temp = tail
			}
		}
	}
	fmt.Println(len(visited))
}
