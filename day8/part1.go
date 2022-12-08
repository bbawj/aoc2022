package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
	left := make([][]int64, m)
	right := make([][]int64, m)
	top := make([][]int64, m)
	bottom := make([][]int64, m)

	for i := 0; i < m; i++ {
		max := int64(-1)
		temp := make([]int64, n)
		for j := 0; j < n; j++ {
			temp[j] = max
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		left[i] = temp
	}
	for i := 0; i < m; i++ {
		max := int64(-1)
		temp := make([]int64, n)
		for j := n - 1; j >= 0; j-- {
			temp[j] = max
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		right[i] = temp
	}

	for i := 0; i < m; i++ {
		temp := make([]int64, n)
		top[i] = temp
		temp = make([]int64, n)
		bottom[i] = temp
	}

	for i := 0; i < n; i++ {
		max := int64(-1)
		for j := 0; j < m; j++ {
			top[j][i] = max
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
	}
	for i := 0; i < n; i++ {
		max := int64(-1)
		for j := m - 1; j >= 0; j-- {
			bottom[j][i] = max
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := grid[i][j]
			l := left[i][j]
			r := right[i][j]
			t := top[i][j]
			b := bottom[i][j]

			if cur > l || cur > r || cur > t || cur > b {
				count++
			}
		}
	}
	fmt.Println(count)
}
