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

func dfs1(root *Node) int64 {
	size := root.Size
	for _, v := range root.Children {
		size += dfs1(v)
	}
	root.Size = size
	return size
}

func dfs2(root *Node, needed int64) []int64 {
	sizes := make([]int64, 0)
	for _, v := range root.Children {
		sizes = append(sizes, dfs2(v, needed)...)
	}
	if root.Size >= needed {
		sizes = append(sizes, root.Size)
	}
	return sizes
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
	dfs1(root)
	remaining := 70000000 - root.Size
	needed := 30000000 - remaining
	fmt.Println(remaining)
	fmt.Println(needed)
	sizes := dfs2(root, needed)

	min := root.Size
	for _, v := range sizes {
		fmt.Println("yeet ", v)
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}
