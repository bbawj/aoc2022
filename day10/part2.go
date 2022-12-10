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

func main() {
	utils.GetInput(10)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cycleCount := 0
	register := 1
	screen := make([][]string, 6)
	for i := range screen {
		screen[i] = make([]string, 40)
	}
	x := 0
	y := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		val := strings.Split(line, " ")
		var arg int64
		op := val[0]
		if op == "addx" {
			arg, _ = strconv.ParseInt(val[1], 10, 64)
			for i := 0; i < 2; i++ {
				cycleCount++
				x = (cycleCount - 1) % 40
				y = int(math.Floor(float64(cycleCount-1) / 40))
				screen[y][x] = "  "
				if x >= register-1 && x <= register+1 {
					screen[y][x] = "##"
				}
			}
			register += int(arg)
		} else {
			cycleCount++
			x = (cycleCount - 1) % 40
			y = int(math.Floor(float64(cycleCount-1) / 40))
			screen[y][x] = "  "
			if x >= register-1 && x <= register+1 {
				screen[y][x] = "##"
			}
		}
	}
	for i := 0; i < len(screen); i++ {
		for j := 0; j < len(screen[i]); j++ {
			fmt.Print(screen[i][j])
		}
		fmt.Println()
	}
}
