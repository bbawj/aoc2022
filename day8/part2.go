package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction struct {
	x int
	y int
}

func main() {
	utils.GetInput(8)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		row := make([]int64, len(line))
		for i := range line {
			row[i], _ = strconv.ParseInt(string(line[i]), 10, 64)
		}
		grid = append(grid, row)
	}
	m := len(grid)
	n := len(grid[0])

	left := Direction{-1, 0}
	right := Direction{1, 0}
	up := Direction{0, 1}
	down := Direction{0, -1}
	dirs := []Direction{left, right, up, down}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			score := 1
			cur := 0
			for k := 0; k < len(dirs); k++ {
				x := j
				y := i
				for x > 0 && x < n-1 && y > 0 && y < m-1 {
					x += dirs[k].x
					y += dirs[k].y
					if grid[y][x] >= grid[i][j] {
						cur++
						break
					}
					cur++
				}
				score *= cur
				cur = 0
			}
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)
}
