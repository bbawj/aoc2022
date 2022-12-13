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
	utils.GetInput(13)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	index := 1
	valid := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		var a, b any
		if line == "" {
			index++
			continue
		} else {
			a, _ = parseExpression(line)
			scanner.Scan()
			b, _ = parseExpression(scanner.Text())
			if less(a, b) >= 0 {
				valid += index
			}
		}
	}
	fmt.Println(valid)
}

func less(a, b any) int {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		if bInt < aInt {
			return -1
		} else if bInt > aInt {
			return 1
		}
		return 0
	}
	aArr, aIsArr := a.([]any)
	bArr, bIsArr := b.([]any)
	if !aIsArr {
		aArr = []any{aInt}
	}
	if !bIsArr {
		bArr = []any{bInt}
	}
	if len(aArr) == 0 && len(bArr) == 0 {
		return 0
	} else if len(aArr) == 0 && len(bArr) > 0 {
		return 1
	} else if len(aArr) > 0 && len(bArr) == 0 {
		return -1
	}
	for i := 0; i < len(aArr); i++ {
		if i >= len(bArr) {
			return -1
		}
		if sub := less(aArr[i], bArr[i]); sub != 0 {
			return sub
		}
	}

	if len(aArr) < len(bArr) {
		return 1
	} else if len(aArr) == len(bArr) {
		return less(aArr[len(aArr)-1], bArr[len(bArr)-1])
	}
	return 0
}

func parseExpression(expr string) (any, int) {
	chars := []rune(expr)
	idx := 0
	out := make([]any, 0)
	nChars := make([]rune, 0)
	for idx < len(chars) {
		char := chars[idx]
		switch char {
		case '[':
			x, i := parseExpression(string(chars[idx+1:]))
			out = append(out, x)
			idx += i + 1
		case ']':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
			return out, idx
		case ',':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
		default:
			nChars = append(nChars, char)
			idx++
		}
	}
	return out, idx
}
