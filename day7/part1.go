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

type Node struct {
	Name     string
	Size     int64
	Children map[string]*Node
	Parent   *Node
}

func dfs(root *Node, sum *int64) int64 {
	size := root.Size
	for _, v := range root.Children {
		size += dfs(v, sum)
	}
	if size <= 100000 {
		*sum += size
	}
	return size
}

func main() {
	utils.GetInput(7)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	root := &Node{"/", 0, make(map[string]*Node), nil}
	var cd *Node
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "$" {
			if lineSplit[1] == "cd" {
				val := strings.Split(line, " ")
				// find the child and navigate in
				if val[2] == "/" {
					cd = root
				} else if val[2] == ".." {
					cd = cd.Parent
				} else {
					cd = cd.Children[val[2]]
				}
			} else {
				// ls command
				continue
			}
		} else {
			val := strings.Split(line, " ")
			if val[0] == "dir" {
				temp := &Node{val[1], 0, make(map[string]*Node), cd}
				cd.Children[val[1]] = temp
				continue
			}
			fileSize, _ := strconv.ParseInt(val[0], 10, 64)
			cd.Size += fileSize
		}

	}
	sum := int64(0)
	dfs(root, &sum)
	fmt.Println(sum)
}
